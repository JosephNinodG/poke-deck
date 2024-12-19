package handler

import (
	"github.com/JosephNinodG/poke-deck/model"
	"github.com/JosephNinodG/poke-deck/tcgapi"
)

type CardHandler interface {
	GetCardById(id string) (model.PokemonCard, error)
	GetCards(req model.GetCardsRequest) ([]model.PokemonCard, error)
}

type TcgApiHandler struct{}

func (t TcgApiHandler) GetCardById(id string) (model.PokemonCard, error) {
	return tcgapi.GetCardById(id)
}

func (t TcgApiHandler) GetCards(req model.GetCardsRequest) ([]model.PokemonCard, error) {
	return tcgapi.GetCards(req)
}
