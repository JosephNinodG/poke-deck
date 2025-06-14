package handler

import (
	"context"

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
)

type Database interface {
	GetUserCollection(ctx context.Context, req domain.GetUserCollectionRequest) ([]domain.PokemonCard, error)
	CreateUserCollection(ctx context.Context, req domain.CreateUserCollectionRequest) error
}

type DatabaseHandler struct {
}

func (d DatabaseHandler) GetUserCollection(ctx context.Context, req domain.GetUserCollectionRequest) ([]domain.PokemonCard, error) {
	return db.GetUserCollection(ctx, req)
}

func (d DatabaseHandler) CreateUserCollection(ctx context.Context, req domain.CreateUserCollectionRequest) error {
	return db.CreateUserCollection(ctx, req)
}
