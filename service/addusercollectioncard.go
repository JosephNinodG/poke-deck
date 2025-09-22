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
			var apiCard domain.PokemonCard
			apiCard, err = cardHandler.GetCardById(cardID)
			if err != nil {
				return err
			}

			setLegalities := lookup.MapLegality(apiCard.Set.Legalities)
			cardLegalities := lookup.MapLegality(apiCard.Legalities)

			dbCardID, err = databaseHandler.AddCard(ctx, setLegalities, cardLegalities, apiCard)
			if err != nil {
				return err
			}

			apiCard.ID = &dbCardID

			lookup.UpdateRecentlyViewedCards(apiCard)

		} else {

			lookup.UpdateRecentlyViewedCards(dbCard)
		}

	} else if recentlyViewedCard.Card.ID == nil {
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

			recentlyViewedCard.Card.ID = &dbCardID

			lookup.UpdateRecentlyViewedCards(recentlyViewedCard.Card)

		} else {

			lookup.UpdateRecentlyViewedCards(dbCard)
		}

	} else {

		dbCardID = *recentlyViewedCard.Card.ID

		lookup.UpdateRecentlyViewedCards(recentlyViewedCard.Card)
	}

	return databaseHandler.AddUserCollectionCard(ctx, dbCardID, collectionID)
}
