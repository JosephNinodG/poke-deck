package tcgapi

import (
	"context"
	"log/slog"

	"github.com/JosephNinodG/poke-deck/model"
)

var stubRepo stubRepository
var cardList []model.PokemonCard

type stubRepository struct {
	cards []model.PokemonCard
}

type StubTcgApiHandler struct{}

func SetUpStubRepository(ctx context.Context, apikey string) {
	stubRepo = stubRepository{cards: stubPokemonCards}
	slog.InfoContext(ctx, "New stub data repository created")
}

func (t StubTcgApiHandler) GetCardById(id string) (model.PokemonCard, error) {
	for _, card := range stubRepo.cards {
		if card.ID == id {
			return card, nil
		}
	}

	return model.PokemonCard{}, nil
}

func (t StubTcgApiHandler) GetCards(req model.GetCardsRequest) ([]model.PokemonCard, error) {
	cardList = []model.PokemonCard{}

	for _, card := range stubRepo.cards {
		if IsValidCard(req, card) {
			cardList = append(cardList, card)
		}
	}

	//TODO: Implement MaxCards
	//TODO: Implement OrderBy

	return cardList, nil
}

func IsValidCard(req model.GetCardsRequest, card model.PokemonCard) bool {
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

var stubPokemonCards = []model.PokemonCard{
	{
		ID:          "test-ID-1",
		Name:        "test-name-1",
		Supertype:   "test-supertype",
		Subtypes:    []string{"test-subtype1", "test-subtype2"},
		Level:       "",
		Hp:          "",
		Types:       []string{"test-type1", "test-type2"},
		EvolvesFrom: "",
		EvolvesTo:   []string{},
		Rules:       []string{},
		AncientTrait: &struct {
			Name string "json:\"name\""
			Text string "json:\"text\""
		}{},
		Abilities: []struct {
			Name string "json:\"name\""
			Text string "json:\"text\""
			Type string "json:\"type\""
		}{},
		Attacks: []struct {
			Name                string   "json:\"name\""
			Cost                []string "json:\"cost\""
			ConvertedEnergyCost int      "json:\"convertedEnergyCost\""
			Damage              string   "json:\"damage\""
			Text                string   "json:\"text\""
		}{{Name: "test-attack1"}, {Name: "test-attack2"}},
		Weaknesses: []struct {
			Type  string "json:\"type\""
			Value string "json:\"value\""
		}{},
		Resistances: []struct {
			Type  string "json:\"type\""
			Value string "json:\"value\""
		}{},
		RetreatCost:          []string{},
		ConvertedRetreatCost: 0,
		Set: struct {
			ID           string "json:\"id\""
			Name         string "json:\"name\""
			Series       string "json:\"series\""
			PrintedTotal int    "json:\"printedTotal\""
			Total        int    "json:\"total\""
			Legalities   struct {
				Unlimited string "json:\"unlimited\""
			} "json:\"legalities\""
			PtcgoCode   string "json:\"ptcgoCode\""
			ReleaseDate string "json:\"releaseDate\""
			UpdatedAt   string "json:\"updatedAt\""
			Images      struct {
				Symbol string "json:\"symbol\""
				Logo   string "json:\"logo\""
			} "json:\"images\""
		}{Name: "test-set"},
		Number:                 "",
		Artist:                 "",
		Rarity:                 "",
		FlavorText:             "",
		NationalPokedexNumbers: []int{},
		Legalities: struct {
			Unlimited string "json:\"unlimited\""
		}{},
		Images: struct {
			Small string "json:\"small\""
			Large string "json:\"large\""
		}{},
		TCGPlayer: struct {
			URL       string "json:\"url\""
			UpdatedAt string "json:\"updatedAt\""
			Prices    struct {
				Holofoil *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"holofoil,omitempty\""
				ReverseHolofoil *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"reverseHolofoil,omitempty\""
				Normal *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"normal,omitempty\""
			} "json:\"prices\""
		}{},
		CardMarket: struct {
			URL       string "json:\"url\""
			UpdatedAt string "json:\"updatedAt\""
			Prices    struct {
				AverageSellPrice *float64 "json:\"averageSellPrice\""
				LowPrice         *float64 "json:\"lowPrice\""
				TrendPrice       *float64 "json:\"trendPrice\""
				GermanProLow     *float64 "json:\"germanProLow\""
				SuggestedPrice   *float64 "json:\"suggestedPrice\""
				ReverseHoloSell  *float64 "json:\"reverseHoloSell\""
				ReverseHoloLow   *float64 "json:\"reverseHoloLow\""
				ReverseHoloTrend *float64 "json:\"reverseHoloTrend\""
				LowPriceExPlus   *float64 "json:\"lowPriceExPlus\""
				Avg1             *float64 "json:\"avg1\""
				Avg7             *float64 "json:\"avg7\""
				Avg30            *float64 "json:\"avg30\""
				ReverseHoloAvg1  *float64 "json:\"reverseHoloAvg1\""
				ReverseHoloAvg7  *float64 "json:\"reverseHoloAvg7\""
				ReverseHoloAvg30 *float64 "json:\"reverseHoloAvg30\""
			} "json:\"prices\""
		}{},
	},
	{
		ID:          "test-ID-2",
		Name:        "test-name-2",
		Supertype:   "test-supertype",
		Subtypes:    []string{"test-subtype1", "test-subtype2"},
		Level:       "",
		Hp:          "",
		Types:       []string{"test-type1", "test-type2"},
		EvolvesFrom: "",
		EvolvesTo:   []string{},
		Rules:       []string{},
		AncientTrait: &struct {
			Name string "json:\"name\""
			Text string "json:\"text\""
		}{},
		Abilities: []struct {
			Name string "json:\"name\""
			Text string "json:\"text\""
			Type string "json:\"type\""
		}{},
		Attacks: []struct {
			Name                string   "json:\"name\""
			Cost                []string "json:\"cost\""
			ConvertedEnergyCost int      "json:\"convertedEnergyCost\""
			Damage              string   "json:\"damage\""
			Text                string   "json:\"text\""
		}{{Name: "test-attack1"}, {Name: "test-attack2"}},
		Weaknesses: []struct {
			Type  string "json:\"type\""
			Value string "json:\"value\""
		}{},
		Resistances: []struct {
			Type  string "json:\"type\""
			Value string "json:\"value\""
		}{},
		RetreatCost:          []string{},
		ConvertedRetreatCost: 0,
		Set: struct {
			ID           string "json:\"id\""
			Name         string "json:\"name\""
			Series       string "json:\"series\""
			PrintedTotal int    "json:\"printedTotal\""
			Total        int    "json:\"total\""
			Legalities   struct {
				Unlimited string "json:\"unlimited\""
			} "json:\"legalities\""
			PtcgoCode   string "json:\"ptcgoCode\""
			ReleaseDate string "json:\"releaseDate\""
			UpdatedAt   string "json:\"updatedAt\""
			Images      struct {
				Symbol string "json:\"symbol\""
				Logo   string "json:\"logo\""
			} "json:\"images\""
		}{Name: "test-set"},
		Number:                 "",
		Artist:                 "",
		Rarity:                 "",
		FlavorText:             "",
		NationalPokedexNumbers: []int{},
		Legalities: struct {
			Unlimited string "json:\"unlimited\""
		}{},
		Images: struct {
			Small string "json:\"small\""
			Large string "json:\"large\""
		}{},
		TCGPlayer: struct {
			URL       string "json:\"url\""
			UpdatedAt string "json:\"updatedAt\""
			Prices    struct {
				Holofoil *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"holofoil,omitempty\""
				ReverseHolofoil *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"reverseHolofoil,omitempty\""
				Normal *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"normal,omitempty\""
			} "json:\"prices\""
		}{},
		CardMarket: struct {
			URL       string "json:\"url\""
			UpdatedAt string "json:\"updatedAt\""
			Prices    struct {
				AverageSellPrice *float64 "json:\"averageSellPrice\""
				LowPrice         *float64 "json:\"lowPrice\""
				TrendPrice       *float64 "json:\"trendPrice\""
				GermanProLow     *float64 "json:\"germanProLow\""
				SuggestedPrice   *float64 "json:\"suggestedPrice\""
				ReverseHoloSell  *float64 "json:\"reverseHoloSell\""
				ReverseHoloLow   *float64 "json:\"reverseHoloLow\""
				ReverseHoloTrend *float64 "json:\"reverseHoloTrend\""
				LowPriceExPlus   *float64 "json:\"lowPriceExPlus\""
				Avg1             *float64 "json:\"avg1\""
				Avg7             *float64 "json:\"avg7\""
				Avg30            *float64 "json:\"avg30\""
				ReverseHoloAvg1  *float64 "json:\"reverseHoloAvg1\""
				ReverseHoloAvg7  *float64 "json:\"reverseHoloAvg7\""
				ReverseHoloAvg30 *float64 "json:\"reverseHoloAvg30\""
			} "json:\"prices\""
		}{},
	},
	{
		ID:          "test-ID-3",
		Name:        "test-name-3",
		Supertype:   "test-supertype",
		Subtypes:    []string{"test-subtype-2", "test-subtype-3"},
		Level:       "",
		Hp:          "",
		Types:       []string{"test-type-3", "test-type-4"},
		EvolvesFrom: "",
		EvolvesTo:   []string{},
		Rules:       []string{},
		AncientTrait: &struct {
			Name string "json:\"name\""
			Text string "json:\"text\""
		}{},
		Abilities: []struct {
			Name string "json:\"name\""
			Text string "json:\"text\""
			Type string "json:\"type\""
		}{},
		Attacks: []struct {
			Name                string   "json:\"name\""
			Cost                []string "json:\"cost\""
			ConvertedEnergyCost int      "json:\"convertedEnergyCost\""
			Damage              string   "json:\"damage\""
			Text                string   "json:\"text\""
		}{{Name: "test-attack1"}, {Name: "test-attack2"}},
		Weaknesses: []struct {
			Type  string "json:\"type\""
			Value string "json:\"value\""
		}{},
		Resistances: []struct {
			Type  string "json:\"type\""
			Value string "json:\"value\""
		}{},
		RetreatCost:          []string{},
		ConvertedRetreatCost: 0,
		Set: struct {
			ID           string "json:\"id\""
			Name         string "json:\"name\""
			Series       string "json:\"series\""
			PrintedTotal int    "json:\"printedTotal\""
			Total        int    "json:\"total\""
			Legalities   struct {
				Unlimited string "json:\"unlimited\""
			} "json:\"legalities\""
			PtcgoCode   string "json:\"ptcgoCode\""
			ReleaseDate string "json:\"releaseDate\""
			UpdatedAt   string "json:\"updatedAt\""
			Images      struct {
				Symbol string "json:\"symbol\""
				Logo   string "json:\"logo\""
			} "json:\"images\""
		}{Name: "test-set"},
		Number:                 "",
		Artist:                 "",
		Rarity:                 "",
		FlavorText:             "",
		NationalPokedexNumbers: []int{},
		Legalities: struct {
			Unlimited string "json:\"unlimited\""
		}{},
		Images: struct {
			Small string "json:\"small\""
			Large string "json:\"large\""
		}{},
		TCGPlayer: struct {
			URL       string "json:\"url\""
			UpdatedAt string "json:\"updatedAt\""
			Prices    struct {
				Holofoil *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"holofoil,omitempty\""
				ReverseHolofoil *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"reverseHolofoil,omitempty\""
				Normal *struct {
					Low    float64 "json:\"low\""
					Mid    float64 "json:\"mid\""
					High   float64 "json:\"high\""
					Market float64 "json:\"market\""
				} "json:\"normal,omitempty\""
			} "json:\"prices\""
		}{},
		CardMarket: struct {
			URL       string "json:\"url\""
			UpdatedAt string "json:\"updatedAt\""
			Prices    struct {
				AverageSellPrice *float64 "json:\"averageSellPrice\""
				LowPrice         *float64 "json:\"lowPrice\""
				TrendPrice       *float64 "json:\"trendPrice\""
				GermanProLow     *float64 "json:\"germanProLow\""
				SuggestedPrice   *float64 "json:\"suggestedPrice\""
				ReverseHoloSell  *float64 "json:\"reverseHoloSell\""
				ReverseHoloLow   *float64 "json:\"reverseHoloLow\""
				ReverseHoloTrend *float64 "json:\"reverseHoloTrend\""
				LowPriceExPlus   *float64 "json:\"lowPriceExPlus\""
				Avg1             *float64 "json:\"avg1\""
				Avg7             *float64 "json:\"avg7\""
				Avg30            *float64 "json:\"avg30\""
				ReverseHoloAvg1  *float64 "json:\"reverseHoloAvg1\""
				ReverseHoloAvg7  *float64 "json:\"reverseHoloAvg7\""
				ReverseHoloAvg30 *float64 "json:\"reverseHoloAvg30\""
			} "json:\"prices\""
		}{},
	},
}
