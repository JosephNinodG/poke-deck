package db

type PokemonCard struct {
	ID                     int        `json:"id"`
	CardID                 string     `json:"cardID"`
	Name                   string     `json:"name"`
	Supertype              string     `json:"supertype"`
	Subtypes               []string   `json:"subtypes"`
	Level                  *string    `json:"level"`
	Hp                     string     `json:"hp"`
	Types                  []string   `json:"types"`
	EvolvesFrom            *string    `json:"evolvesFrom"`
	EvolvesTo              []string   `json:"evolvesTo"`
	Rules                  []string   `json:"rules"`
	AncientTrait           *Traits    `json:"ancientTrait"`
	Abilities              []*Traits  `json:"abilities"`
	Attacks                []Attack   `json:"attacks"`
	Weaknesses             []Traits   `json:"weaknesses"`
	Resistances            []Traits   `json:"resistances"`
	RetreatCost            []string   `json:"retreatCost"`
	ConvertedRetreatCost   int        `json:"convertedRetreatCost"`
	Set                    Set        `json:"set"`
	Number                 string     `json:"number"`
	Artist                 string     `json:"artist"`
	Rarity                 string     `json:"rarity"`
	FlavorText             string     `json:"flavorText"`
	NationalPokedexNumbers []int      `json:"nationalPokedexNumbers"`
	Legalities             Legalities `json:"legalities"`
	Images                 Images     `json:"images"`
}

type Traits struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Type        *string `json:"type"`
	Value       *string `json:"value"`
}

type Images struct {
	Small  *string `json:"small"`
	Large  *string `json:"large"`
	Symbol *string `json:"symbol"`
	Logo   *string `json:"logo"`
}

type Legalities struct {
	Standard  string `json:"standard"`
	Expanded  string `json:"expanded"`
	Unlimited string `json:"unlimited"`
}

type Attack struct {
	Name                string   `json:"name"`
	Cost                []string `json:"cost"`
	ConvertedEnergyCost int      `json:"convertedEnergyCost"`
	Damage              string   `json:"damage"`
	Description         string   `json:"description"`
}

type Set struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Series       string     `json:"series"`
	PrintedTotal int        `json:"printedTotal"`
	Total        int        `json:"total"`
	Legalities   Legalities `json:"legalities"`
	PtcgoCode    string     `json:"ptcgoCode"`
	ReleaseDate  string     `json:"releaseDate"`
	UpdatedAt    string     `json:"updatedAt"`
	Images       Images     `json:"images"`
}
