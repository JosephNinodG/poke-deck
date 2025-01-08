package tcgapi

import (
	"github.com/JosephNinodG/poke-deck/domain"
	pokemontcgv2 "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
)

func CardMapper(apiCard pokemontcgv2.PokemonCard) domain.PokemonCard {
	var pokemonCard = domain.PokemonCard{
		ID:                     apiCard.ID,
		Name:                   apiCard.Name,
		Supertype:              apiCard.Supertype,
		Subtypes:               apiCard.Subtypes,
		Level:                  apiCard.Level,
		Hp:                     apiCard.Hp,
		Types:                  apiCard.Types,
		EvolvesFrom:            apiCard.EvolvesFrom,
		EvolvesTo:              apiCard.EvolvesTo,
		Rules:                  apiCard.Rules,
		RetreatCost:            apiCard.RetreatCost,
		ConvertedRetreatCost:   apiCard.ConvertedRetreatCost,
		Number:                 apiCard.Number,
		Artist:                 apiCard.Artist,
		Rarity:                 apiCard.Rarity,
		FlavorText:             apiCard.FlavorText,
		NationalPokedexNumbers: apiCard.NationalPokedexNumbers,
	}

	for _, apiAbility := range apiCard.Abilities {
		var traits = domain.Traits{
			Name:        &apiAbility.Name,
			Description: &apiAbility.Text,
			Type:        &apiAbility.Type,
		}
		pokemonCard.Abilities = append(pokemonCard.Abilities, traits)
	}

	for _, apiAttack := range apiCard.Attacks {
		var attack = domain.Attack{
			Name:                apiAttack.Name,
			Description:         apiAttack.Text,
			Cost:                apiAttack.Cost,
			ConvertedEnergyCost: apiAttack.ConvertedEnergyCost,
			Damage:              apiAttack.Damage,
		}
		pokemonCard.Attacks = append(pokemonCard.Attacks, attack)
	}

	for _, apiWeakness := range apiCard.Weaknesses {
		var traits = domain.Traits{
			Type:  &apiWeakness.Type,
			Value: &apiWeakness.Value,
		}
		pokemonCard.Weaknesses = append(pokemonCard.Weaknesses, traits)
	}

	for _, apiResistances := range apiCard.Resistances {
		var traits = domain.Traits{
			Type:  &apiResistances.Type,
			Value: &apiResistances.Value,
		}
		pokemonCard.Resistances = append(pokemonCard.Resistances, traits)
	}

	if apiCard.AncientTrait != nil {
		pokemonCard.AncientTrait = &domain.Traits{
			Name:        &apiCard.AncientTrait.Name,
			Description: &apiCard.AncientTrait.Text,
		}
	}

	pokemonCard.Images = domain.Images{
		Small: &apiCard.Images.Small,
		Large: &apiCard.Images.Large,
	}

	pokemonCard.Set = domain.Set{
		ID:          apiCard.Set.ID,
		Name:        apiCard.Set.Name,
		Series:      apiCard.Set.Series,
		PtcgoCode:   apiCard.Set.PtcgoCode,
		ReleaseDate: apiCard.Set.ReleaseDate,
		UpdatedAt:   apiCard.Set.UpdatedAt,
		Images: domain.Images{
			Symbol: &apiCard.Set.Images.Symbol,
			Logo:   &apiCard.Set.Images.Logo,
		},
	}

	return pokemonCard
}
