package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/JosephNinodG/poke-deck/domain"
)

func GetUserCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	endpointName := "GetUserCollection"

	var req GetUserCollectionRequest

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

	if !req.IsValid() {
		slog.ErrorContext(ctx, "request is invalid", "endpoint", endpointName, "request", req)
		w.WriteHeader(http.StatusBadRequest) //TODO: Change to 204
		_, err := w.Write([]byte("payload must have non-zero values"))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	slog.InfoContext(ctx, "request received", "endpoint", endpointName, "request", req)

	getUserCollectionRequest := domain.GetUserCollection{
		UserID:       req.UserID,
		CollectionID: req.CollectionID,
	}

	response, err := databaseHandler.GetUserCollection(ctx, getUserCollectionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error getting specified card", "endpoint", endpointName, "request", req, "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error encoding response body", "endpoint", endpointName, "error", err)
		return
	}

	slog.InfoContext(ctx, "response returned successfully", "endpoint", endpointName, "request", req)
}
