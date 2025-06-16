package db

import (
	"context"
	"fmt"
	"log/slog"
)

func AddUserCollectionCard(ctx context.Context, cardID, collectionID int) error {
	result, err := conn.Exec(ctx, addUserCollectionCardQuery, cardID, collectionID)
	if err != nil {
		return fmt.Errorf("unable to connect to execute AddUserCollectionCard query. %v", err.Error())
	}

	rows := result.RowsAffected()

	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	slog.DebugContext(ctx, "request to database successful", "CardID", cardID, "CollectionID", collectionID)

	return nil
}

var addUserCollectionCardQuery = "INSERT INTO collection_card (card_id, collection_id) VALUES ($1, $2);"
