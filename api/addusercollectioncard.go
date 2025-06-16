package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/lookup"
)

func AddUserCollectionCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	endpointName := "AddUserCollectionCard"

	var req AddUserCollectionCardRequest
	var card domain.PokemonCard
	var err error

	if strings.ToUpper(r.Method) != http.MethodPost {
		slog.ErrorContext(ctx, "HTTP method not allowed on route", "path", r.URL.Path, "expected", http.MethodPost, "actual", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(fmt.Sprintf("HTTP method not allowed on route. Expected %v", http.MethodPost)))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.ErrorContext(ctx, "error reading request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("error decoding JSON request body"))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "error", err)
		}
		return
	}

	if !req.IsValid() {
		slog.ErrorContext(ctx, "request is invalid", "endpoint", endpointName, "request", req)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("payload must have non-zero values"))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	slog.InfoContext(ctx, "request received", "endpoint", endpointName, "request", req)

	var dbCardID int
	dbCard, err := db.GetCardById(ctx, req.CardID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error getting specified card from db", "endpoint", endpointName, "cardId", req.CardID, "error", err)
		return
	}

	if reflect.ValueOf(dbCard).IsZero() {
		var card domain.PokemonCard
		recentlyViewedCard, ok := lookup.RecentlyViewedCards[req.CardID]
		if !ok {
			card, err = cardHandler.GetCardById(req.CardID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				slog.ErrorContext(ctx, "error getting specified card from tcgapi", "endpoint", endpointName, "cardId", req.CardID, "error", err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if reflect.ValueOf(card).IsZero() {
				w.WriteHeader(http.StatusNotFound)
				_, err := w.Write([]byte("no card matching that Id"))
				if err != nil {
					slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
				}
				return
			}

		} else {
			card = recentlyViewedCard.Card
		}

		setLegalities := lookup.MapLegality(card.Set.Legalities)
		cardLegalities := lookup.MapLegality(card.Legalities)

		err = databaseHandler.AddCard(ctx, setLegalities, cardLegalities, card)

		//TODO: Return newly added cardID from DB and assign to dbCardID

	} else {
		dbCardID = dbCard.ID
	}

	err = databaseHandler.AddUserCollectionCard(ctx, dbCardID, req.CollectionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error adding specified card to collection", "endpoint", endpointName, "cardId", req.CardID, "collectionId", req.CollectionID, "error", err)
		return
	}

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error encoding response body", "endpoint", endpointName, "error", err)
		return
	}

	lookup.UpdateRecentlyViewedCards(nil, card)

	slog.InfoContext(ctx, "response returned successfully", "endpoint", endpointName, "cardId", card.ID)
}
