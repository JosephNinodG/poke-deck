package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/JosephNinodG/poke-deck/model"
	"github.com/JosephNinodG/poke-deck/tcgapi"
)

var tcgapikey string

func TestMain(m *testing.M) {
	ctx := context.Background()
	tcgapi.SetUpClient(ctx, tcgapikey)

	m.Run()
}

func Test_GetCardById(t *testing.T) {
	tests := []struct {
		name                  string
		id                    string
		queryParam            string
		httpMethod            string
		expectedStatusCode    int
		expectedErrorResponse string
		expectedResponse      model.PokemonCard
	}{
		{
			name:                  "Error - Incorrect http method",
			id:                    "test-id",
			queryParam:            "id",
			httpMethod:            http.MethodDelete,
			expectedStatusCode:    http.StatusMethodNotAllowed,
			expectedErrorResponse: "HTTP method not allowed on route. Expected GET",
		},
		{
			name:                  "Error - Incorrect query parameter",
			id:                    "test-id",
			queryParam:            "incorrectparam",
			httpMethod:            http.MethodGet,
			expectedStatusCode:    http.StatusBadRequest,
			expectedErrorResponse: "missing id in request",
		},
		{
			name:                  "Error - Empty Id",
			id:                    "",
			queryParam:            "id",
			httpMethod:            http.MethodGet,
			expectedStatusCode:    http.StatusBadRequest,
			expectedErrorResponse: "missing id in request",
		},
		{
			name:                  "Error - No card for given Id",
			id:                    "abcdefgh",
			queryParam:            "id",
			httpMethod:            http.MethodGet,
			expectedStatusCode:    http.StatusNotFound,
			expectedErrorResponse: "no card matching that Id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.httpMethod, fmt.Sprintf("/getcardbyid?%v=%v", test.queryParam, test.id), nil)
			w := httptest.NewRecorder()
			GetCardById(w, req)
			res := w.Result()
			defer func() {
				err := res.Body.Close()
				if err != nil {
					t.Fatalf("error closing io.ReadCloser for res.Body")
				}
			}()

			if res.StatusCode != test.expectedStatusCode {
				t.Fatalf("expected\n%v\nactual\n%v", res.StatusCode, test.expectedStatusCode)
			}

			if test.expectedErrorResponse != "" {
				response, err := getResponseBodyAsString(res.Body)
				if err != nil {
					t.Fatalf("error converting response body to string")
				}

				if test.expectedErrorResponse != response {
					t.Fatalf("\nexpected\n%v\nactual\n%v", test.expectedResponse, response)
				}
			}

			if test.expectedStatusCode == http.StatusOK {
				var response model.PokemonCard

				err := json.NewDecoder(res.Body).Decode(&response)
				if err != nil {
					t.Fatalf("error decoding json body: %v", err)
				}

				if !reflect.DeepEqual(test.expectedResponse, response) {
					t.Fatalf("\nexpected\n%v\nactual\n%v", test.expectedResponse, response)
				}
			}

		})
	}
}

func getResponseBodyAsString(responseBody io.ReadCloser) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(responseBody)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

var pokemonCard = model.PokemonCard{
	ID:          "",
	Name:        "",
	Supertype:   "",
	Subtypes:    []string{},
	Level:       "",
	Hp:          "",
	Types:       []string{},
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
	}{},
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
	}{},
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
}
