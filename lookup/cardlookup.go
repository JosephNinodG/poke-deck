package lookup

import (
	"context"
	"log/slog"
	"time"

	"github.com/JosephNinodG/poke-deck/domain"
)

type RecentlyViewedCard struct {
	DatabaseId *int
	Card       domain.PokemonCard
	TimeViewed time.Time
}

var RecentlyViewedCards map[string]RecentlyViewedCard

func UpdateRecentlyViewedCards(databaseId *int, card domain.PokemonCard) {
	if isInRecentlyViewedCards(card.ID) {
		updateViewedTime(card.ID)
	} else {
		addToRecentlyViewedCards(databaseId, card)
	}
}

func isInRecentlyViewedCards(cardId string) bool {
	_, ok := RecentlyViewedCards[cardId]
	return ok
}

func addToRecentlyViewedCards(databaseId *int, card domain.PokemonCard) {

	var recentlyViewedCard = RecentlyViewedCard{
		DatabaseId: databaseId,
		Card:       card,
		TimeViewed: time.Now(),
	}

	RecentlyViewedCards[card.ID] = recentlyViewedCard
}

func updateViewedTime(cardId string) {
	viewedCard := RecentlyViewedCards[cardId]
	viewedCard.TimeViewed = time.Now()
	RecentlyViewedCards[cardId] = viewedCard
}

func SetupLookup(ctx context.Context) error {
	cards, err := databaseHandler.GetAllCards(ctx)
	if err != nil {
		return err
	}

	setupTime := time.Now()

	for databaseId, card := range cards {
		var recentlyViewedCard = RecentlyViewedCard{
			DatabaseId: &databaseId,
			Card:       card,
			TimeViewed: setupTime,
		}
		RecentlyViewedCards[card.ID] = recentlyViewedCard
	}

	return nil
}

func CheckLookupTimes(ctx context.Context, done <-chan bool) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			select {
			case <-done:
				slog.InfoContext(ctx, "Stopping ticker")
				ticker.Stop()
				return
			case <-ticker.C:

				for key, recentlyViewCard := range RecentlyViewedCards {
					if time.Now().Sub(recentlyViewCard.TimeViewed).Minutes() > 5 {
						delete(RecentlyViewedCards, key)
					}
				}
			}
		}
	}()
}
