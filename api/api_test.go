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

	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/tcgapi"
)

var tcgapikey string

func TestMain(m *testing.M) {
	ctx := context.Background()
	Configure(tcgapi.StubTcgApiHandler{}, db.StubDatabaseHandler{})

	tcgapi.SetUpStubRepository(ctx, tcgapikey)

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
		expectedCardId        string
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
		{
			name:                  "Success - Retreived card for given Id",
			id:                    "test-ID-1",
			queryParam:            "id",
			httpMethod:            http.MethodGet,
			expectedStatusCode:    http.StatusOK,
			expectedErrorResponse: "",
			expectedCardId:        "test-ID-1",
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
				t.Fatalf("expected\n%v\nactual\n%v", test.expectedStatusCode, res.StatusCode)
			}

			if test.expectedErrorResponse != "" {
				response, err := getResponseBodyAsString(res.Body)
				if err != nil {
					t.Fatalf("error converting response body to string")
				}

				if test.expectedErrorResponse != response {
					t.Fatalf("\nexpected\n%v\nactual\n%v", test.expectedErrorResponse, response)
				}
			}

			if test.expectedStatusCode == http.StatusOK {
				var response domain.PokemonCard

				err := json.NewDecoder(res.Body).Decode(&response)
				if err != nil {
					t.Fatalf("error decoding json body: %v", err)
				}

				if !reflect.DeepEqual(test.expectedCardId, response.ID) {
					t.Fatalf("\nexpected\n%v\nactual\n%v", test.expectedCardId, response.ID)
				}
			}

		})
	}
}

func Test_GetCards(t *testing.T) {
	tests := []struct {
		name                  string
		request               []byte
		httpMethod            string
		expectedStatusCode    int
		expectedErrorResponse string
		expectedCardId        []string
	}{
		{
			name:                  "Error - Incorrect http method",
			request:               []byte(`{"card":{"name":"test-name","type":"test-type","supertype":"test-supertype","subtype":"test-subtype","set":"test-set","attack":"test-attack","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}}`),
			httpMethod:            http.MethodDelete,
			expectedStatusCode:    http.StatusMethodNotAllowed,
			expectedErrorResponse: "HTTP method not allowed on route. Expected GET",
		},
		{
			name:               "Successful - Retrieved card",
			request:            []byte(`{"card":{"name":"test-name-1","type":"test-type-1","supertype":"test-supertype","subtype":"test-subtype-1","set":"test-set","attack":"test-attack-1","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{"test-ID-1"},
		},
		{
			name:               "Successful - Retrieved mulitple cards",
			request:            []byte(`{"card":{"type":"test-type-1","supertype":"test-supertype","subtype":"test-subtype-1","set":"test-set","attack":"test-attack-1","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{"test-ID-1", "test-ID-2"},
		},
		{
			name:               "Successful - No cards for given params",
			request:            []byte(`{"card":{"type":"test-type-false","supertype":"test-supertype","subtype":"test-subtype-false","set":"test-set","attack":"test-attack-false","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{},
		},
		{
			name:               "Successful - Valid request for multiple cards but filtered by MaxCards",
			request:            []byte(`{"card":{"type":"test-type-1","supertype":"test-supertype","subtype":"test-subtype-1","set":"test-set","attack":"test-attack-1","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}, "parameters":{"maxCards":1}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{"test-ID-1"},
		},
		{
			name:               "Successful - Retrieved mulitple cards and ordered using OrderBy = number",
			request:            []byte(`{"card":{"type":"test-type-1","supertype":"test-supertype","subtype":"test-subtype-1","set":"test-set","attack":"test-attack-1","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}, "parameters":{"orderby":"number"}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{"test-ID-2", "test-ID-1"}, //50, 100
		},
		{
			name:               "Successful - Retrieved mulitple cards and ordered using OrderBy = number - desc",
			request:            []byte(`{"card":{"type":"test-type-1","supertype":"test-supertype","subtype":"test-subtype-1","set":"test-set","attack":"test-attack-1","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}, "parameters":{"orderby":"number", "desc": true}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{"test-ID-1", "test-ID-2"}, //100, 50
		},
		{
			name:               "Successful - Retrieved mulitple cards and ordered using OrderBy = name - desc",
			request:            []byte(`{"card":{"type":"test-type-1","supertype":"test-supertype","subtype":"test-subtype-1","set":"test-set","attack":"test-attack-1","legalities":{"standard":"legal","expanded":"legal","unlimited":"legal"}}, "parameters":{"orderby":"name", "desc": true}}`),
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedCardId:     []string{"test-ID-2", "test-ID-1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.httpMethod, "/getcards", bytes.NewBuffer(test.request))
			w := httptest.NewRecorder()
			GetCards(w, req)
			res := w.Result()
			defer func() {
				err := res.Body.Close()
				if err != nil {
					t.Fatalf("error closing io.ReadCloser for res.Body")
				}
			}()

			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("expected\n%v\nactual\n%v", test.expectedStatusCode, res.StatusCode)
			}

			if test.expectedErrorResponse != "" {
				response, err := getResponseBodyAsString(res.Body)
				if err != nil {
					t.Fatalf("error converting response body to string")
				}

				if test.expectedErrorResponse != response {
					t.Fatalf("\nexpected\n%v\nactual\n%v", test.expectedErrorResponse, response)
				}
			}

			if test.expectedStatusCode == http.StatusOK {
				var response []domain.PokemonCard

				err := json.NewDecoder(res.Body).Decode(&response)
				if err != nil {
					t.Fatalf("error decoding json body: %v", err)
				}

				for i, card := range response {
					if !reflect.DeepEqual(test.expectedCardId[i], card.ID) {
						t.Fatalf("\nexpected\n%v\nactual\n%v", test.expectedCardId[i], card.ID)
					}
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
