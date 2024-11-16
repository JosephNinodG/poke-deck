package api

import "github.com/JosephNinodG/poke-deck/handler"

var cardHandler handler.CardHandler

func Configure(cardHandlerOverride handler.CardHandler) {
	if cardHandler != nil {
		panic("cardHandler instance already set")
	}
	cardHandler = cardHandlerOverride
}
