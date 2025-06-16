package service

import (
	"context"
	"reflect"

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/lookup"
)

func AddUserCollectionCard(ctx context.Context, cardID string, collectionID int) error {
	var dbCardID int

	recentlyViewedCard, ok := lookup.RecentlyViewedCards[cardID]
	if !ok {

		dbCard, err := db.GetCardById(ctx, cardID)
		if err != nil {
			return err
		}

		if reflect.ValueOf(dbCard).IsZero() {
			var card domain.PokemonCard
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

			lookup.UpdateRecentlyViewedCards(&dbCardID, card)

		} else {
			dbCardID = dbCard.ID

			lookup.UpdateRecentlyViewedCards(&dbCardID, dbCard.Card)
		}

	} else if recentlyViewedCard.DatabaseId == nil {
		dbCard, err := db.GetCardById(ctx, cardID)
		if err != nil {
			return err
		}

		if reflect.ValueOf(dbCard).IsZero() {

			setLegalities := lookup.MapLegality(recentlyViewedCard.Card.Set.Legalities)
			cardLegalities := lookup.MapLegality(recentlyViewedCard.Card.Legalities)

			dbCardID, err = databaseHandler.AddCard(ctx, setLegalities, cardLegalities, recentlyViewedCard.Card)
			if err != nil {
				return err
			}

			lookup.UpdateRecentlyViewedCards(&dbCardID, recentlyViewedCard.Card)

		} else {
			dbCardID = dbCard.ID

			lookup.UpdateRecentlyViewedCards(&dbCardID, dbCard.Card)
		}

	} else {
		dbCardID = *recentlyViewedCard.DatabaseId

		lookup.UpdateRecentlyViewedCards(&dbCardID, recentlyViewedCard.Card)
	}

	return databaseHandler.AddUserCollectionCard(ctx, dbCardID, collectionID)
}
