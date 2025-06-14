package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/domain"
)

var stubRepo stubRepository

type userCollection struct {
	userID     int
	collection map[int][]domain.PokemonCard
}

type stubRepository struct {
	collections []userCollection
}

type StubDatabaseHandler struct{}

func SetUpStubRepository(ctx context.Context, apikey string) {
	stubRepo = stubRepository{collections: stubCollections}
	slog.InfoContext(ctx, "New stub data repository created")
}

func (d StubDatabaseHandler) GetUserCollection(ctx context.Context, req domain.GetUserCollection) ([]domain.PokemonCard, error) {
	for _, collection := range stubRepo.collections {
		if collection.userID == req.UserID {
			userCollection, ok := collection.collection[req.CollectionID]
			if !ok {
				return nil, fmt.Errorf("no collections for userID %v", req.UserID)
			}

			return userCollection, nil
		}
	}

	return nil, nil
}

var stubCollections = []userCollection{
	{
		userID: 1,
		collection: map[int][]domain.PokemonCard{
			1: {
				{
					ID:        "test-ID-1",
					Name:      "test-name-1",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-1", "test-subtype-2"},
					Types:     []string{"test-type-1", "test-type-2"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "100",
				},
				{
					ID:        "test-ID-2",
					Name:      "test-name-2",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-1", "test-subtype-2"},
					Types:     []string{"test-type-1", "test-type-2"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "50",
				},
				{
					ID:        "test-ID-3",
					Name:      "test-name-3",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-2", "test-subtype-3"},
					Types:     []string{"test-type-3", "test-type-4"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "0",
				},
			},
		},
	},
	{
		userID: 2,
		collection: map[int][]domain.PokemonCard{
			1: {
				{
					ID:        "test-ID-1",
					Name:      "test-name-1",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-1", "test-subtype-2"},
					Types:     []string{"test-type-1", "test-type-2"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "100",
				},
			},
			2: {
				{
					ID:        "test-ID-1",
					Name:      "test-name-1",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-1", "test-subtype-2"},
					Types:     []string{"test-type-1", "test-type-2"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "100",
				},
				{
					ID:        "test-ID-2",
					Name:      "test-name-2",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-1", "test-subtype-2"},
					Types:     []string{"test-type-1", "test-type-2"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "50",
				},
				{
					ID:        "test-ID-3",
					Name:      "test-name-3",
					Supertype: "test-supertype",
					Subtypes:  []string{"test-subtype-2", "test-subtype-3"},
					Types:     []string{"test-type-3", "test-type-4"},
					Attacks:   []domain.Attack{{Name: "test-attack-1"}, {Name: "test-attack-2"}},
					Set:       domain.Set{Name: "test-set"},
					Number:    "0",
				},
			},
		},
	},
}
