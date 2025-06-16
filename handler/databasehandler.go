package handler

import (
	"context"

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
)

type Database interface {
	GetUserCollection(ctx context.Context, req domain.GetUserCollectionRequest) ([]domain.PokemonCard, error)
	CreateUserCollection(ctx context.Context, req domain.CreateUserCollectionRequest) error
	GetAllCards(ctx context.Context) (map[int]domain.PokemonCard, error)
	AddUserCollectionCard(ctx context.Context, cardID, collectionID int) error
	GetCardById(ctx context.Context, cardID string) (domain.DbCard, error)
	AddCard(ctx context.Context, setLegalities, cardLegalities int, card domain.PokemonCard) error
}

type DatabaseHandler struct {
}

func (d DatabaseHandler) GetUserCollection(ctx context.Context, req domain.GetUserCollectionRequest) ([]domain.PokemonCard, error) {
	return db.GetUserCollection(ctx, req)
}

func (d DatabaseHandler) CreateUserCollection(ctx context.Context, req domain.CreateUserCollectionRequest) error {
	return db.CreateUserCollection(ctx, req)
}

func (d DatabaseHandler) GetAllCards(ctx context.Context) (map[int]domain.PokemonCard, error) {
	return db.GetAllCards(ctx)
}

func (d DatabaseHandler) AddUserCollectionCard(ctx context.Context, cardID, collectionID int) error {
	return db.AddUserCollectionCard(ctx, cardID, collectionID)
}

func (d DatabaseHandler) GetCardById(ctx context.Context, cardID string) (domain.DbCard, error) {
	return db.GetCardById(ctx, cardID)
}

func (d DatabaseHandler) AddCard(ctx context.Context, setLegalities, cardLegalities int, card domain.PokemonCard) error {
	return db.AddCard(ctx, setLegalities, cardLegalities, card)
}
