package tcgapi

import (
	"fmt"

	"github.com/JosephNinodG/poke-deck/model"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

//TODO: May have to build entire new api wrapper to handle requests instead of relying on sdk
// struct provided by sdk is missing legalities.standard and legalities.expanded
// func GetCardByLegality(queryParams string) {
// 	ctx := context.Background()
// 	requestURL := fmt.Sprintf("https://api.pokemontcg.io/v2/cards%v", queryParams)
// 	res, err := http.Get(requestURL)
// 	if err != nil {
// 		slog.ErrorContext(ctx, err.Error())
// 	}
// }

func GetCardById(id string) (model.PokemonCard, error) {
	card, err := client.GetCardByID(id)
	if err != nil {
		return model.PokemonCard{}, err
	}

	if card == nil {
		return model.PokemonCard{}, nil
	}

	pokemonCard := model.PokemonCard(*card)

	return pokemonCard, nil
}

func GetCards(req model.GetCardsRequest) ([]model.PokemonCard, error) {
	if req.Paramters.OrderBy != "" && req.Paramters.Desc {
		req.Paramters.OrderBy = fmt.Sprintf("-%v", req.Paramters.OrderBy)
	}

	cards, err := client.GetCards(
		request.Query(buildQuery(req.Card)...),
		request.OrderBy(req.Paramters.OrderBy),
		request.PageSize(req.Paramters.MaxCards),
	)
	if err != nil {
		return nil, err
	}

	var pokemonCards []model.PokemonCard
	for _, card := range cards {
		pokemonCards = append(pokemonCards, model.PokemonCard(*card))
	}

	return pokemonCards, nil
}

func buildQuery(card model.Card) []string {
	var query []string

	if card.Name != "" {
		query = append(query, fmt.Sprintf("name:%v", card.Name))
	}

	if card.Type != "" {
		query = append(query, fmt.Sprintf("types:%v", card.Type))
	}

	if card.Supertype != "" {
		query = append(query, fmt.Sprintf("supertypes:%v", card.Supertype))
	}

	if card.Subtype != "" {
		query = append(query, fmt.Sprintf("subtypes:%v", card.Subtype))
	}

	if card.Set != "" {
		query = append(query, fmt.Sprintf("set.name:%v", card.Set))
	}

	if card.Attack != "" {
		query = append(query, fmt.Sprintf("attacks.name:%v", card.Attack))
	}

	if card.Legalities.Standard != "" {
		query = append(query, fmt.Sprintf("legalities.standard:%v", card.Legalities.Standard))
	}

	if card.Legalities.Expanded != "" {
		query = append(query, fmt.Sprintf("legalities.expanded:%v", card.Legalities.Expanded))
	}

	if card.Legalities.Unlimited != "" {
		query = append(query, fmt.Sprintf("legalities.unlimited:%v", card.Legalities.Unlimited))
	}

	return query
}
