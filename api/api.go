package api

import (
	"net/http"

	"github.com/JosephNinodG/poke-deck/handler"
)

var (
	cardHandler     handler.Card
	databaseHandler handler.Database
)

func Configure(cardHandlerOverride handler.Card, databaseHandlerOverride handler.Database) {
	if cardHandler != nil {
		panic("cardHandler instance already set")
	}

	if databaseHandler != nil {
		panic("databaseHandler instance already set")
	}

	cardHandler = cardHandlerOverride
	databaseHandler = databaseHandlerOverride
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
