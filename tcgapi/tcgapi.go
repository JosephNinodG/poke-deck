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

	"github.com/JosephNinodG/poke-deck/model"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

//TODO: Add mapper from pokemon-tcg-sdk-go-v2 to model.PokemonCard to allow Legalities in struct

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

func GetCards(req model.GetCardsRequest, apikey string) ([]model.PokemonCard, error) {
	if req.Paramters.OrderBy != "" && req.Paramters.Desc {
		req.Paramters.OrderBy = fmt.Sprintf("-%v", req.Paramters.OrderBy)
	}

	var pokemonCards []model.PokemonCard
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
			pokemonCards = append(pokemonCards, model.PokemonCard(*card))
		}
	}

	return pokemonCards, nil
}

func getCardsWithQueryParams(req model.GetCardsRequest, apikey string) ([]model.PokemonCard, error) {
	queryList := buildQuery(req.Card)
	queryString := buildQueryString(queryList)
	//queryParams := addQueryParameters(queryString, req.Paramters)

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

	// requestURL := fmt.Sprintf("https://api.pokemontcg.io/v2/cards%v", queryParams)
	// res, err := http.Get(requestURL)
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
		Data []model.PokemonCard `json:"data"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&cardList); err != nil {
		return nil, err
	}

	// resBody, _ := getResponseBodyAsString(res.Body)

	// var data Data
	// err = json.Unmarshal([]byte(resBody), &data)
	// if err != nil {
	// 	return nil, err
	// }

	// var pokemonCards []model.PokemonCard
	// for _, card := range data.cards {
	// 	pokemonCards = append(pokemonCards, model.PokemonCard(*card))
	// }

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

func addQueryParameters(queryString string, parameters model.Parameters) string {
	if parameters.OrderBy != "" {
		queryString = fmt.Sprintf("%v&orderBy=%v", queryString, parameters.OrderBy)
	}

	if parameters.MaxCards != 0 {
		queryString = fmt.Sprintf("%v&pageSize=%v", queryString, parameters.MaxCards)
	}

	return queryString
}
