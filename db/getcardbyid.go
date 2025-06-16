package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/domain"
)

func GetCardById(ctx context.Context, cardID string) (domain.DbCard, error) {
	rows, err := conn.Query(ctx, selectCardByIdQuery, cardID)
	if err != nil {
		return domain.DbCard{}, fmt.Errorf("unable to connect to execute GetCardById query %v", err.Error())
	}
	defer rows.Close()

	var cards []PokemonCard
	for rows.Next() {
		var cardByte []byte
		if err := rows.Scan(&cardByte); err != nil {
			return domain.DbCard{}, fmt.Errorf("unable to connect to read rows %v", err.Error())
		}

		var pokemonCard PokemonCard
		if err := json.Unmarshal(cardByte, &pokemonCard); err != nil {
			return domain.DbCard{}, fmt.Errorf("failed to unmarshal JSON %v", err.Error())
		}

		cards = append(cards, pokemonCard)
	}

	if err := rows.Err(); err != nil {
		return domain.DbCard{}, fmt.Errorf("row iteration error. %v", err.Error())
	}

	if len(cards) > 1 {
		return domain.DbCard{}, fmt.Errorf("expected single card, got %d cards", len(cards))
	}

	dbCard := domain.DbCard{
		ID:   cards[0].ID,
		Card: cards[0].MapToDomain(),
	}

	slog.DebugContext(ctx, "request to database successful")

	return dbCard, nil
}

var selectCardByIdQuery = "SELECT cardID FROM card WHERE cardID = $1;"
