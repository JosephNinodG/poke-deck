package lookup

import (
	"context"
	"log/slog"
	"time"

	"github.com/JosephNinodG/poke-deck/domain"
)

type RecentlyViewedCard struct {
	Card       domain.PokemonCard
	TimeViewed time.Time
}

var RecentlyViewedCards map[string]RecentlyViewedCard

func UpdateRecentlyViewedCards(card domain.PokemonCard) {
	if isInRecentlyViewedCards(card.CardID) {
		updateViewedTime(card.CardID)
	} else {
		addToRecentlyViewedCards(card)
	}
}

func isInRecentlyViewedCards(cardId string) bool {
	_, ok := RecentlyViewedCards[cardId]
	return ok
}

func addToRecentlyViewedCards(card domain.PokemonCard) {

	var recentlyViewedCard = RecentlyViewedCard{
		Card:       card,
		TimeViewed: time.Now(),
	}

	RecentlyViewedCards[card.CardID] = recentlyViewedCard
}

func updateViewedTime(cardId string) {
	viewedCard := RecentlyViewedCards[cardId]
	viewedCard.TimeViewed = time.Now()
	RecentlyViewedCards[cardId] = viewedCard
}

func SetupLookup(ctx context.Context) error {
	RecentlyViewedCards = make(map[string]RecentlyViewedCard)

	cards, err := databaseHandler.GetAllCards(ctx)
	if err != nil {
		return err
	}

	setupTime := time.Now()

	for _, card := range cards {
		var recentlyViewedCard = RecentlyViewedCard{
			Card:       card,
			TimeViewed: setupTime,
		}
		RecentlyViewedCards[card.CardID] = recentlyViewedCard
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
