package handler

import (
	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/tcgapi"
)

type CardHandler interface {
	GetCardById(id string) (domain.PokemonCard, error)
	GetCards(req domain.GetCardsRequest) ([]domain.PokemonCard, error)
}

type TcgApiHandler struct {
	Apikey string
}

func (t TcgApiHandler) GetCardById(id string) (domain.PokemonCard, error) {
	return tcgapi.GetCardById(id)
}

func (t TcgApiHandler) GetCards(req domain.GetCardsRequest) ([]domain.PokemonCard, error) {
	return tcgapi.GetCards(req, t.Apikey)
}
