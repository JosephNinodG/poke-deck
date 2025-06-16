package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/domain"
)

func GetAllCards(ctx context.Context) (map[int]domain.PokemonCard, error) {
	rows, err := conn.Query(ctx, selectUserCollectionQuery)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to execute GetAllCards query %v", err.Error())
	}
	defer rows.Close()

	var cards map[int]domain.PokemonCard
	for rows.Next() {
		var cardByte []byte
		if err := rows.Scan(&cardByte); err != nil {
			return nil, fmt.Errorf("unable to connect to read rows %v", err.Error())
		}

		var pokemonCard PokemonCard
		if err := json.Unmarshal(cardByte, &pokemonCard); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON %v", err.Error())
		}

		cards[pokemonCard.ID] = pokemonCard.MapToDomain()
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error. %v", err.Error())
	}

	slog.DebugContext(ctx, "request to database successful")

	return cards, nil
}

var selectAllCardsInDbQuery = "SELECT cardID FROM card GROUP BY cardID;"
