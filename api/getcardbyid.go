package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/lookup"
)

func GetCardById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	endpointName := "GetCardById"
	enableCors(&w)

	if strings.ToUpper(r.Method) != http.MethodGet {
		slog.ErrorContext(ctx, "HTTP method not allowed on route", "path", r.URL.Path, "expected", http.MethodGet, "actual", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(fmt.Sprintf("HTTP method not allowed on route. Expected %v", http.MethodGet)))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	if id == "" {
		slog.ErrorContext(ctx, "no Id provided in query param", "endpoint", endpointName)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("missing id in request"))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	slog.InfoContext(ctx, "request received", "endpoint", endpointName, "cardId", id)

	var card domain.PokemonCard
	var err error
	_, ok := lookup.RecentlyViewedCards[id]
	if !ok {
		card, err = cardHandler.GetCardById(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.ErrorContext(ctx, "error getting specified card", "endpoint", endpointName, "cardId", id, "error", err)
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
	}

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error encoding response body", "endpoint", endpointName, "error", err)
		return
	}

	lookup.UpdateRecentlyViewedCards(nil, card)

	slog.InfoContext(ctx, "response returned successfully", "endpoint", endpointName, "cardId", id)
}
