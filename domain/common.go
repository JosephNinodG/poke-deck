package domain

type PokemonCard struct {
	ID                     string
	Name                   string
	Supertype              string
	Subtypes               []string
	Level                  *string
	Hp                     string
	Types                  []string
	EvolvesFrom            *string
	EvolvesTo              []string
	Rules                  []string
	AncientTrait           *Traits
	Abilities              []*Traits
	Attacks                []Attack
	Weaknesses             []Traits
	Resistances            []Traits
	RetreatCost            []string
	ConvertedRetreatCost   int
	Set                    Set
	Number                 string
	Artist                 string
	Rarity                 string
	FlavorText             string
	NationalPokedexNumbers []int
	Legalities             Legalities
	Images                 Images
}

type Traits struct {
	Name        *string
	Description *string
	Type        *string
	Value       *string
}

type Images struct {
	Small  *string
	Large  *string
	Symbol *string
	Logo   *string
}

type Legalities struct {
	Standard  string
	Expanded  string
	Unlimited string
}

type Attack struct {
	Name                string
	Cost                []string
	ConvertedEnergyCost int
	Damage              string
	Description         string
}

type Set struct {
	ID           string
	Name         string
	Series       string
	PrintedTotal int
	Total        int
	Legalities   Legalities
	PtcgoCode    string
	ReleaseDate  string
	UpdatedAt    string
	Images       Images
}
