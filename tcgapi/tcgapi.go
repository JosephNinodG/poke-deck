package tcgapi

import (
	"log"

	"github.com/JosephNinodG/poke-deck/model"
	pokemontcgv2 "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func GetCard(GetCardRequest model.GetCardRequest) (*pokemontcgv2.PokemonCard, error) {
	card, err := client.GetCardByID(GetCardRequest.Id)
	//TODO: Handle specific response codes

	if err != nil {
		return nil, err
	}

	return card, nil
}

func GetCardById(Id string) (model.PokemonCard, error) {
	card, err := client.GetCardByID(Id)
	if err != nil {
		return model.PokemonCard{}, err
	}

	if card == nil {
		return model.PokemonCard{}, nil
	}

	pokemonCard := model.PokemonCard(*card)

	return pokemonCard, nil
}

// Example GetCards func
func GetCards() {
	cards, err := client.GetCards(
		request.Query("name:jirachi", "types:psychic"),
		request.PageSize(5),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, card := range cards {
		log.Println(card.Name)
	}
}
