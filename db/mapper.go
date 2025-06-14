package db

import "github.com/JosephNinodG/poke-deck/domain"

func (p *PokemonCard) MapToDomain() domain.PokemonCard {
	var pokemonCard = domain.PokemonCard{
		ID:                     p.ID,
		Name:                   p.Name,
		Supertype:              p.Supertype,
		Subtypes:               p.Subtypes,
		Level:                  p.Level,
		Hp:                     p.Hp,
		Types:                  p.Types,
		EvolvesFrom:            p.EvolvesFrom,
		EvolvesTo:              p.EvolvesTo,
		Rules:                  p.Rules,
		RetreatCost:            p.RetreatCost,
		ConvertedRetreatCost:   p.ConvertedRetreatCost,
		Number:                 p.Number,
		Artist:                 p.Artist,
		Rarity:                 p.Rarity,
		FlavorText:             p.FlavorText,
		NationalPokedexNumbers: p.NationalPokedexNumbers,
	}

	for _, apiAbility := range p.Abilities {
		var traits = domain.Traits{
			Name:        apiAbility.Name,
			Description: apiAbility.Description,
			Type:        apiAbility.Type,
		}
		pokemonCard.Abilities = append(pokemonCard.Abilities, &traits)
	}

	for _, apiAttack := range p.Attacks {
		var attack = domain.Attack{
			Name:                apiAttack.Name,
			Description:         apiAttack.Description,
			Cost:                apiAttack.Cost,
			ConvertedEnergyCost: apiAttack.ConvertedEnergyCost,
			Damage:              apiAttack.Damage,
		}
		pokemonCard.Attacks = append(pokemonCard.Attacks, attack)
	}

	for _, apiWeakness := range p.Weaknesses {
		var traits = domain.Traits{
			Type:  apiWeakness.Type,
			Value: apiWeakness.Value,
		}
		pokemonCard.Weaknesses = append(pokemonCard.Weaknesses, traits)
	}

	for _, apiResistances := range p.Resistances {
		var traits = domain.Traits{
			Type:  apiResistances.Type,
			Value: apiResistances.Value,
		}
		pokemonCard.Resistances = append(pokemonCard.Resistances, traits)
	}

	if p.AncientTrait != nil {
		pokemonCard.AncientTrait = &domain.Traits{
			Name:        p.AncientTrait.Name,
			Description: p.AncientTrait.Description,
		}
	}

	pokemonCard.Images = domain.Images{
		Small: p.Images.Small,
		Large: p.Images.Large,
	}

	pokemonCard.Set = domain.Set{
		ID:          p.Set.ID,
		Name:        p.Set.Name,
		Series:      p.Set.Series,
		PtcgoCode:   p.Set.PtcgoCode,
		ReleaseDate: p.Set.ReleaseDate,
		UpdatedAt:   p.Set.UpdatedAt,
		Images: domain.Images{
			Symbol: p.Set.Images.Symbol,
			Logo:   p.Set.Images.Logo,
		},
	}

	return pokemonCard
}
