package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/domain"
)

func GetAllCards(ctx context.Context) ([]domain.PokemonCard, error) {
	rows, err := conn.Query(ctx, selectAllCardsInDbQuery)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to execute GetAllCards query %v", err.Error())
	}
	defer rows.Close()

	var pokemonCards []PokemonCard
	for rows.Next() {
		var cardByte []byte
		if err := rows.Scan(&cardByte); err != nil {
			return nil, fmt.Errorf("unable to connect to read rows %v", err.Error())
		}

		if err := json.Unmarshal(cardByte, &pokemonCards); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON %v", err.Error())
		}
	}

	var cards = []domain.PokemonCard{}

	for _, pokemonCard := range pokemonCards {
		cards = append(cards, pokemonCard.MapToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error. %v", err.Error())
	}

	slog.DebugContext(ctx, "request to database successful")

	return cards, nil
}

var selectAllCardsInDbQuery = `
SELECT json_agg(card_data) AS cards
FROM (
  SELECT
    c.*,
    -- Subqueries to include referenced objects
    row_to_json(set_data) AS "set",
    row_to_json(card_legalities) AS "card_legalities",
    row_to_json(card_images) AS "card_images",
    row_to_json(ancient_trait) AS "ancient_trait",

    -- Arrays of joined entities
    (
      SELECT json_agg(row_to_json(a))
      FROM (
        SELECT ab.*
        FROM card_ability ca
        JOIN ability ab ON ca.ability_id = ab.id
        WHERE ca.card_id = c.id
      ) a
    ) AS abilities,

    (
      SELECT json_agg(row_to_json(atk))
      FROM (
        SELECT atk.*
        FROM card_attack ca
        JOIN attack atk ON ca.attack_id = atk.id
        WHERE ca.card_id = c.id
      ) atk
    ) AS attacks,

    (
      SELECT json_agg(row_to_json(w))
      FROM (
        SELECT w.*
        FROM card_weakness cw
        JOIN weakness w ON cw.weakness_id = w.id
        WHERE cw.card_id = c.id
      ) w
    ) AS weaknesses,

    (
      SELECT json_agg(row_to_json(r))
      FROM (
        SELECT r.*
        FROM card_resistance cr
        JOIN resistance r ON cr.resistance_id = r.id
        WHERE cr.card_id = c.id
      ) r
    ) AS resistances

  FROM card c

  LEFT JOIN ancient_trait ON c.ancient_trait_id = ancient_trait.id
  LEFT JOIN card_legalities ON c.card_legalities_id = card_legalities.id
  LEFT JOIN card_images ON c.card_images_id = card_images.id
  LEFT JOIN (
    SELECT s.*, row_to_json(si) AS set_images
    FROM "set" s
    LEFT JOIN set_images si ON s.set_images_id = si.id
  ) set_data ON c.set_id = set_data.id

) card_data;
`
