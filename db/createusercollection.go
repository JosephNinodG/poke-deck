package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/domain"
)

func CreateUserCollection(ctx context.Context, req domain.CreateUserCollectionRequest) error {
	result, err := conn.Exec(ctx, createUserCollectionQuery, req.CollectionName, req.UserID)
	if err != nil {
		return fmt.Errorf("unable to connect to execute CreateUserCollection query. %v", err.Error())
	}

	rows := result.RowsAffected()

	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	slog.DebugContext(ctx, "request to database successful", "UserID", req.UserID)

	return nil
}

var createUserCollectionQuery = "INSERT INTO collection (name, user_id) VALUES ($1, $2);"
