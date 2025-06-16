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
	collection map[int]collection
}

type collection struct {
	collectionName  string
	collectionCards []domain.PokemonCard
}

type stubRepository struct {
	collections []userCollection
}

type StubDatabaseHandler struct{}

func SetUpStubRepository(ctx context.Context, apikey string) {
	stubRepo = stubRepository{collections: stubCollections}
	slog.InfoContext(ctx, "New stub data repository created")
}

func (d StubDatabaseHandler) GetUserCollection(ctx context.Context, req domain.GetUserCollectionRequest) ([]domain.PokemonCard, error) {
	for _, collection := range stubRepo.collections {
		if collection.userID == req.UserID {
			userCollection, ok := collection.collection[req.CollectionID]
			if !ok {
				return nil, fmt.Errorf("no collections for userID %v", req.UserID)
			}

			return userCollection.collectionCards, nil
		}
	}

	return nil, nil
}

func (d StubDatabaseHandler) CreateUserCollection(ctx context.Context, req domain.CreateUserCollectionRequest) error {
	var currentHighestKey int

	for k := range stubRepo.collections {
		if k > currentHighestKey {
			currentHighestKey = k
		}
	}

	currentHighestKey++

	newCollection := userCollection{userID: req.UserID, collection: map[int]collection{currentHighestKey: {collectionName: req.CollectionName, collectionCards: []domain.PokemonCard{}}}}

	stubRepo.collections = append(stubRepo.collections, newCollection)

	return nil
}

// TODO: Add stubbing for funcs
func (d StubDatabaseHandler) GetAllCards(ctx context.Context) (map[int]domain.PokemonCard, error) {
	return nil, nil
}

func (d StubDatabaseHandler) AddUserCollectionCard(ctx context.Context, cardID, collectionID int) error {
	return nil
}

func (d StubDatabaseHandler) GetCardById(ctx context.Context, cardID string) (domain.DbCard, error) {
	return domain.DbCard{}, nil
}

func (d StubDatabaseHandler) AddCard(ctx context.Context, setLegalities, cardLegalities int, card domain.PokemonCard) (int, error) {
	return 0, nil
}

var stubCollections = []userCollection{
	{
		userID: 1,
		collection: map[int]collection{
			1: {
				collectionName: "test-user1-collection1",
				collectionCards: []domain.PokemonCard{
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
	},
}
