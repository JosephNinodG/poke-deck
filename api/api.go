package api

import (
	"net/http"

	"github.com/JosephNinodG/poke-deck/handler"
)

var cardHandler handler.CardHandler

func Configure(cardHandlerOverride handler.CardHandler) {
	if cardHandler != nil {
		panic("cardHandler instance already set")
	}
	cardHandler = cardHandlerOverride
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
