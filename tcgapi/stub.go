package tcgapi

import (
	"context"
	"log/slog"
	"sort"

	"github.com/JosephNinodG/poke-deck/domain"
)

var stubRepo stubRepository
var cardList []domain.PokemonCard

type stubRepository struct {
	cards []domain.PokemonCard
}

type StubTcgApiHandler struct{}

func SetUpStubRepository(ctx context.Context, apikey string) {
	stubRepo = stubRepository{cards: stubPokemonCards}
	slog.InfoContext(ctx, "New stub data repository created")
}

func (t StubTcgApiHandler) GetCardById(id string) (domain.PokemonCard, error) {
	for _, card := range stubRepo.cards {
		if card.ID == id {
			return card, nil
		}
	}

	return domain.PokemonCard{}, nil
}

func (t StubTcgApiHandler) GetCards(req domain.GetCardsRequest) ([]domain.PokemonCard, error) {
	cardList = []domain.PokemonCard{}

	for _, card := range stubRepo.cards {
		if IsValidCard(req, card) {
			cardList = append(cardList, card)
		}
	}

	switch req.Paramters.OrderBy {
	case "name":
		sort.SliceStable(cardList, func(i, j int) bool {
			return cardList[i].Name < cardList[j].Name
		})
	case "number":
		sort.SliceStable(cardList, func(i, j int) bool {
			return cardList[i].Number > cardList[j].Number
		})
	case "set":
		sort.SliceStable(cardList, func(i, j int) bool {
			return cardList[i].Set.Name < cardList[j].Set.Name
		})
	default:
	}

	if req.Paramters.Desc {
		sort.SliceStable(cardList, func(i, j int) bool {
			return i > j
		})
	}

	if req.Paramters.MaxCards == 0 {
		req.Paramters.MaxCards = 250
	}

	if len(cardList) > req.Paramters.MaxCards {
		trimmedList := []domain.PokemonCard{}
		for i := 0; i < req.Paramters.MaxCards; i++ {
			trimmedList = append(trimmedList, cardList[i])
		}

		cardList = trimmedList
	}

	return cardList, nil
}

func IsValidCard(req domain.GetCardsRequest, card domain.PokemonCard) bool {
	if req.Card.Name != "" && req.Card.Name != card.Name {
		return false
	}

	//TODO: Add check for legality

	if req.Card.Supertype != "" && req.Card.Supertype != card.Supertype {
		return false
	}

	if req.Card.Set != "" && req.Card.Set != card.Set.Name {
		return false
	}

	if req.Card.Type != "" {
		matchFound := false

		for _, primaryType := range card.Types {
			if primaryType == req.Card.Type {
				matchFound = true
			}
		}

		if !matchFound {
			return false
		}
	}

	if req.Card.Subtype != "" {
		matchFound := false

		for _, subType := range card.Subtypes {
			if subType == req.Card.Subtype {
				matchFound = true
			}
		}

		if !matchFound {
			return false
		}
	}

	if req.Card.Attack != "" {
		matchFound := false

		for _, attack := range card.Attacks {
			if attack.Name == req.Card.Attack {
				matchFound = true
			}
		}

		if !matchFound {
			return false
		}
	}

	return true
}

var stubPokemonCards = []domain.PokemonCard{
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
}
