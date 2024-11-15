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

func Test_GetCards(t *testing.T) {
	tests := []struct {
		name                  string
		request               model.GetCardsRequest
		queryParam            string
		httpMethod            string
		expectedStatusCode    int
		expectedErrorResponse string
		expectedResponse      model.PokemonCard
	}{
		{
			name: "Error - Incorrect http method",
			request: model.GetCardsRequest{
				Card: model.Card{
					Name:      "test-name",
					Type:      "test-type",
					Supertype: "test-supertype",
					Subtype:   "test-subtype",
					Set:       "test-set",
					Attack:    "test-attack",
					Legalities: model.Legalities{
						Standard:  "legal",
						Expanded:  "legal",
						Unlimited: "legal",
					},
				},
				Paramters: model.Parameters{
					MaxCards: 1,
					OrderBy:  "name",
				},
			},
			queryParam:            "id",
			httpMethod:            http.MethodDelete,
			expectedStatusCode:    http.StatusMethodNotAllowed,
			expectedErrorResponse: "HTTP method not allowed on route. Expected GET",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.httpMethod, "/getcards", nil)
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
