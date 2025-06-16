package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/domain"
)

func AddCard(ctx context.Context, setLegalities, cardLegalities int, card domain.PokemonCard) (int, error) {
	var dbCardID int
	err := conn.QueryRow(ctx, addCardQuery,
		card.Set.Images.Symbol, card.Set.Images.Logo,
		card.Set.Name, card.Set.Series, card.Set.PrintedTotal, card.Set.Total, card.Set.PtcgoCode, card.Set.ReleaseDate, card.Set.UpdatedAt, setLegalities,
		card.Images.Small, card.Images.Large,
		card.AncientTrait.Name, card.AncientTrait.Description,
		card.ID, card.Name, card.Supertype, card.Subtypes, card.Level, card.Hp, card.Types,
		card.EvolvesFrom, card.EvolvesTo, card.Rules, card.RetreatCost, card.ConvertedRetreatCost,
		card.Number, card.Artist, card.Rarity, card.FlavorText, card.NationalPokedexNumbers,
		cardLegalities,
	).Scan(&dbCardID)
	if err != nil {
		return dbCardID, fmt.Errorf("unable execute QueryRow for AddCard. %v", err.Error())
	}

	slog.DebugContext(ctx, "card added to database successfully", "CardID", card.ID)

	return dbCardID, nil
}

// TODO: Add other queries for attack, ability, resistance, and weakness arrays.
var addCardQuery = `
	WITH
	ins_set_images AS (
		INSERT INTO set_images(symbol, logo)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
		RETURNING id
	), get_set_images AS (
		SELECT id FROM ins_set_images
		UNION
		SELECT id FROM set_images WHERE symbol = $1 AND logo = $2
	),
	ins_set AS (
		INSERT INTO "set"(
			name, series, printedTotal, total, ptcgoCode,
			releaseDate, updatedAt, set_legalities_id, set_images_id
		)
		SELECT $3, $4, $5, $6, $7, $8, $9, $10, get_set_images.id
		FROM get_set_images
		ON CONFLICT DO NOTHING
		RETURNING id
	), get_set AS (
		SELECT id FROM ins_set
		UNION
		SELECT id FROM "set" WHERE name = $3
	),
	ins_card_images AS (
		INSERT INTO card_images(small, large)
		VALUES ($11, $12)
		ON CONFLICT DO NOTHING
		RETURNING id
	), get_card_images AS (
		SELECT id FROM ins_card_images
		UNION
		SELECT id FROM card_images WHERE small = $11 AND large = $12
	),
	ins_ancient_trait AS (
		INSERT INTO ancient_trait(name, description)
		VALUES ($13, $14)
		ON CONFLICT DO NOTHING
		RETURNING id
	), get_ancient_trait AS (
		SELECT id FROM ins_ancient_trait
		UNION
		SELECT id FROM ancient_trait WHERE name = $13
	),
	ins_card AS (
		INSERT INTO card(
			cardID, name, supertype, subtypes, level, hp, types,
			evolvesFrom, evolvesTo, rules, retreatCost, convertedRetreatCost,
			number, artist, rarity, flavorText, nationalPokedexNumbers,
			ancient_trait_id, set_id, card_legalities_id, card_images_id
		)
		SELECT
			$15, $16, $17, $18, $19, $20, $21,
			$22, $23, $24, $25, $26,
			$27, $28, $29, $30, $31,
			get_ancient_trait.id, get_set.id, $32, get_card_images.id
		FROM get_ancient_trait, get_set, get_card_images
		RETURNING id
	)
	SELECT id FROM ins_card;
	`
