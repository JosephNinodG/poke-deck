package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/lookup"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	endpointName := "GetCard"

	var req GetCardsRequest

	if strings.ToUpper(r.Method) != http.MethodGet {
		slog.ErrorContext(ctx, "HTTP method not allowed on route", "path", r.URL.Path, "expected", http.MethodGet, "actual", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(fmt.Sprintf("HTTP method not allowed on route. Expected %v", http.MethodGet)))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.ErrorContext(ctx, "error reading request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("error decoding JSON request body"))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "error", err)
		}
		return
	}

	valid, message := req.IsValid()
	if !valid {
		slog.ErrorContext(ctx, "request is invalid", "endpoint", endpointName, "request", req)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(message))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	slog.InfoContext(ctx, "request received", "endpoint", endpointName, "request", req)

	getCardRequest := domain.GetCardsRequest{
		Card: domain.CardDetails{
			Name:       req.Card.Name,
			Type:       req.Card.Type,
			Supertype:  req.Card.Supertype,
			Subtype:    req.Card.Subtype,
			Set:        req.Card.Set,
			Attack:     req.Card.Attack,
			Legalities: domain.Legalities(req.Card.Legalities),
		},
		Paramters: domain.Parameters(req.Paramters),
	}

	cards, err := cardHandler.GetCards(getCardRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error getting specified card", "endpoint", endpointName, "request", req, "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error encoding response body", "endpoint", endpointName, "error", err)
		return
	}

	for _, card := range cards {
		lookup.UpdateRecentlyViewedCards(nil, card)
	}

	slog.InfoContext(ctx, "response returned successfully", "endpoint", endpointName, "request", req)
}
