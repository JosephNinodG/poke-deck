package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/JosephNinodG/poke-deck/domain"
)

func CreateUserCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	endpointName := "CreateUserCollection"

	var req CreateUserCollectionRequest

	if strings.ToUpper(r.Method) != http.MethodPost {
		slog.ErrorContext(ctx, "HTTP method not allowed on route", "path", r.URL.Path, "expected", http.MethodPost, "actual", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(fmt.Sprintf("HTTP method not allowed on route. Expected %v", http.MethodPost)))
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
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("payload must have non-zero values"))
		if err != nil {
			slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
		}
		return
	}

	slog.InfoContext(ctx, "request received", "endpoint", endpointName, "request", req)

	createUserCollectionRequest := domain.CreateUserCollectionRequest{
		UserID:         req.UserID,
		CollectionName: req.CollectionName,
	}

	err = databaseHandler.CreateUserCollection(ctx, createUserCollectionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(ctx, "error getting specified card", "endpoint", endpointName, "request", req, "error", err)
		return
	}

	slog.DebugContext(ctx, "response returned successfully", "endpoint", endpointName, "request", req)
}
