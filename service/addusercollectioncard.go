package service

import (
	"context"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/lookup"
)

func AddUserCollectionCard(ctx context.Context, cardID string, collectionID int) error {
	var dbCardID int
	var err error

	var card domain.PokemonCard
	recentlyViewedCard, ok := lookup.RecentlyViewedCards[cardID]
	if !ok || recentlyViewedCard.DatabaseId != nil {

		dbCard, err := db.GetCardById(ctx, cardID)
		if err != nil {
			return err
		}

		if reflect.ValueOf(dbCard).IsZero() {
			card, err = cardHandler.GetCardById(cardID)
			if err != nil {
				return err
			}

			setLegalities := lookup.MapLegality(card.Set.Legalities)
			cardLegalities := lookup.MapLegality(card.Legalities)

			dbCardID, err = databaseHandler.AddCard(ctx, setLegalities, cardLegalities, card)
			if err != nil {
				return err
			}
		} else {
			dbCardID = dbCard.ID
		}

	} else {
		dbCardID = *recentlyViewedCard.DatabaseId
	}

	if reflect.ValueOf(dbCard).IsZero() {

		setLegalities := lookup.MapLegality(card.Set.Legalities)
		cardLegalities := lookup.MapLegality(card.Legalities)

		dbCardID, err = databaseHandler.AddCard(ctx, setLegalities, cardLegalities, card)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("error adding card to card table"))
			if err != nil {
				slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
			}
			return
		}

		if reflect.ValueOf(dbCardID).IsZero() {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("error getting db ID of newly added card"))
			if err != nil {
				slog.ErrorContext(ctx, "error writing to HTTP response body", "endpoint", endpointName, "error", err)
			}
			return
		}

	} else {
		dbCardID = dbCard.ID
	}

	return databaseHandler.AddUserCollectionCard(ctx, dbCardID, req.CollectionID)
}
