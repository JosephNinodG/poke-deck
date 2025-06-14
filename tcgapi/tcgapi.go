package tcgapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

//TODO: Add mapper from pokemon-tcg-sdk-go-v2 to model.PokemonCard to allow Legalities in struct

func GetCardById(id string) (domain.PokemonCard, error) {
	card, err := client.GetCardByID(id)
	if err != nil {
		return domain.PokemonCard{}, err
	}

	if card == nil {
		return domain.PokemonCard{}, nil
	}

	pokemonCard := CardMapper(*card)

	return pokemonCard, nil
}

func GetCards(req domain.GetCardsRequest, apikey string) ([]domain.PokemonCard, error) {
	if req.Paramters.OrderBy != "" && req.Paramters.Desc {
		req.Paramters.OrderBy = fmt.Sprintf("-%v", req.Paramters.OrderBy)
	}

	var pokemonCards []domain.PokemonCard
	var err error

	if req.Card.Legalities.Standard != "" || req.Card.Legalities.Expanded != "" {
		pokemonCards, err = getCardsWithQueryParams(req, apikey)
		if err != nil {
			return nil, err
		}

	} else {
		cards, err := client.GetCards(
			request.Query(buildQuery(req.Card)...),
			request.OrderBy(req.Paramters.OrderBy),
			request.PageSize(req.Paramters.MaxCards),
		)
		if err != nil {
			return nil, err
		}

		for _, card := range cards {
			pokemonCards = append(pokemonCards, CardMapper(*card))
		}
	}

	return pokemonCards, nil
}

func getCardsWithQueryParams(req domain.GetCardsRequest, apikey string) ([]domain.PokemonCard, error) {
	queryList := buildQuery(req.Card)
	queryString := buildQueryString(queryList)

	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, "GET", "https://api.pokemontcg.io/v2/cards", bytes.NewBuffer([]byte(``)))
	if err != nil {
		return nil, err
	}

	request.Header.Add("x-api-key", apikey)
	values := request.URL.Query()
	values.Add("q", queryString)
	values.Add("orderBy", req.Paramters.OrderBy)
	values.Add("pageSize", strconv.Itoa(req.Paramters.MaxCards))

	request.URL.RawQuery = values.Encode()
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			slog.ErrorContext(ctx, err.Error())
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TCGAPI returned with status code: %v", res.StatusCode)
	}

	cardList := struct {
		Data []domain.PokemonCard `json:"data"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&cardList); err != nil {
		return nil, err
	}

	return cardList.Data, nil
}

func getResponseBodyAsString(responseBody io.ReadCloser) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(responseBody)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func buildQuery(card domain.CardDetails) []string {
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

func buildQueryString(queryList []string) string {
	//queryString := "?q="
	var queryString string

	for i, query := range queryList {
		if i == 0 {
			queryString = fmt.Sprintf("%v%v", queryString, query)
		} else {
			queryString = fmt.Sprintf("%v %v", queryString, query)
		}

	}

	return queryString
}

func addQueryParameters(queryString string, parameters domain.Parameters) string {
	if parameters.OrderBy != "" {
		queryString = fmt.Sprintf("%v&orderBy=%v", queryString, parameters.OrderBy)
	}

	if parameters.MaxCards != 0 {
		queryString = fmt.Sprintf("%v&pageSize=%v", queryString, parameters.MaxCards)
	}

	return queryString
}
