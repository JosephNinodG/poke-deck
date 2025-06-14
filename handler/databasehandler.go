package handler

import (
	"context"

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
)

type Database interface {
	GetUserCollection(ctx context.Context, req domain.GetUserCollection) ([]domain.PokemonCard, error)
}

type DatabaseHandler struct {
}

func (d DatabaseHandler) GetUserCollection(ctx context.Context, req domain.GetUserCollection) ([]domain.PokemonCard, error) {
	return db.GetUserCollection(ctx, req)
}
