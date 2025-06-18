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
		return dbCardID, fmt.Errorf("unable to execute QueryRow for AddCard. %v", err.Error())
	}

	for _, ability := range card.Abilities {
		err = addCardAbility(ctx, dbCardID, ability)
		if err != nil {
			return dbCardID, err
		}
	}

	for _, attack := range card.Attacks {
		err = addCardAttack(ctx, dbCardID, attack)
		if err != nil {
			return dbCardID, err
		}
	}

	for _, resistance := range card.Resistances {
		err = addCardResisitance(ctx, dbCardID, resistance)
		if err != nil {
			return dbCardID, err
		}
	}

	for _, weakness := range card.Weaknesses {
		err = addCardWeakness(ctx, dbCardID, weakness)
		if err != nil {
			return dbCardID, err
		}
	}

	slog.DebugContext(ctx, "card added to database successfully", "CardID", card.ID)

	return dbCardID, nil
}

func addCardAbility(ctx context.Context, dbCardID int, ability *domain.Traits) error {
	var addCardAbilityQuery = `
		ins_ability AS (
		INSERT INTO ability("name", "description", "type")
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING
		RETURNING id
		), get_ability AS (
		SELECT id FROM ins_ability
		UNION
		SELECT id FROM ability WHERE "name" = $1
		),

		card_ability_link AS (
		INSERT INTO card_ability(card_id, ability_id)
		SELECT $4, get_ability.id FROM $4, get_ability
		ON CONFLICT DO NOTHING
		),

		SELECT 'Card and related data inserted successfully.';
	`

	result, err := conn.Exec(ctx, addCardAbilityQuery, &ability.Name, &ability.Description, &ability.Type, dbCardID)
	if err != nil {
		return fmt.Errorf("unable to connect to execute addCardAbility query. %v", err.Error())
	}

	rows := result.RowsAffected()

	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
}

func addCardAttack(ctx context.Context, dbCardID int, attack domain.Attack) error {
	var addCardAttackQuery = `
	ins_attack AS (
	INSERT INTO attack("name", "cost", "convertedEnergyCost", "damage", "description")
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT DO NOTHING
	RETURNING id
	), get_attack AS (
	SELECT id FROM ins_attack
	UNION
	SELECT id FROM attack WHERE "name" = $1
	),

	card_attack_link AS (
	INSERT INTO card_attack(card_id, attack_id)
	SELECT $6, get_attack.id FROM $6, get_attack
	ON CONFLICT DO NOTHING
	)

	SELECT 'Card and related data inserted successfully.';
	`

	result, err := conn.Exec(ctx, addCardAttackQuery, attack.Name, attack.Cost, attack.ConvertedEnergyCost, attack.Damage, attack.Description, dbCardID)
	if err != nil {
		return fmt.Errorf("unable to connect to execute addCardAttack query. %v", err.Error())
	}

	rows := result.RowsAffected()

	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
}

func addCardResisitance(ctx context.Context, dbCardID int, resistance domain.Traits) error {
	var addCardResisitanceQuery = `
	ins_resistance AS (
	INSERT INTO resistance("type", "value")
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING
	RETURNING id
	), get_resistance AS (
	SELECT id FROM ins_resistance
	UNION
	SELECT id FROM resistance WHERE "type" = $1 AND "value" = $2
	),

	card_resistance_link AS (
	INSERT INTO card_resistance(card_id, resistance_id)
	SELECT $3, get_resistance.id FROM $3, get_resistance
	ON CONFLICT DO NOTHING
	)

	SELECT 'Card and related data inserted successfully.';
	`

	result, err := conn.Exec(ctx, addCardResisitanceQuery, resistance.Type, resistance.Value, dbCardID)
	if err != nil {
		return fmt.Errorf("unable to connect to execute addCardResisitance query. %v", err.Error())
	}

	rows := result.RowsAffected()

	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
}

func addCardWeakness(ctx context.Context, dbCardID int, weakness domain.Traits) error {
	var addCardWeaknessQuery = `
	ins_weakness AS (
	INSERT INTO weakness("type", "value")
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING
	RETURNING id
	), get_weakness AS (
	SELECT id FROM ins_weakness
	UNION
	SELECT id FROM weakness WHERE "type" = $1 AND "value" = $2
	),

	card_weakness_link AS (
	INSERT INTO card_weakness(card_id, weakness_id)
	SELECT $3, get_weakness.id FROM $3, get_weakness
	ON CONFLICT DO NOTHING
	),

	SELECT 'Card and related data inserted successfully.';
	`

	result, err := conn.Exec(ctx, addCardWeaknessQuery, weakness.Type, weakness.Value, dbCardID)
	if err != nil {
		return fmt.Errorf("unable to connect to execute addCardWeakness query. %v", err.Error())
	}

	rows := result.RowsAffected()

	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
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
