package domain

type GetCardsRequest struct {
	Card      CardDetails `json:"card"`
	Paramters Parameters  `json:"parameters"`
}

type CardDetails struct {
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Supertype  string     `json:"supertype"`
	Subtype    string     `json:"subtype"`
	Set        string     `json:"set"`
	Attack     string     `json:"attack"`
	Legalities Legalities `json:"legalities"`
}

type Parameters struct {
	MaxCards int    `json:"maxCards"`
	OrderBy  string `json:"orderBy"`
	Desc     bool   `json:"desc"`
}

type PokemonCard struct {
	ID                     string     `db:"id" json:"id"`
	Name                   string     `db:"name" json:"name"`
	Supertype              string     `db:"supertype" json:"supertype"`
	Subtypes               []string   `db:"subtypes" json:"subtypes"`
	Level                  *string    `db:"level" json:"level"`
	Hp                     string     `db:"hp" json:"hp"`
	Types                  []string   `db:"types" json:"types"`
	EvolvesFrom            *string    `db:"evolvesFrom" json:"evolvesFrom"`
	EvolvesTo              []string   `db:"evolvesTo" json:"evolvesTo"`
	Rules                  []string   `db:"rules" json:"rules"`
	AncientTrait           *Traits    `db:"ancientTrait" json:"ancientTrait"`
	Abilities              []*Traits  `db:"abilities" json:"abilities"`
	Attacks                []Attack   `db:"attacks" json:"attacks"`
	Weaknesses             []Traits   `db:"weaknesses" json:"weaknesses"`
	Resistances            []Traits   `db:"resistances" json:"resistances"`
	RetreatCost            []string   `db:"retreatCost" json:"retreatCost"`
	ConvertedRetreatCost   int        `db:"convertedRetreatCost" json:"convertedRetreatCost"`
	Set                    Set        `db:"set" json:"set"`
	Number                 string     `db:"number" json:"number"`
	Artist                 string     `db:"artist" json:"artist"`
	Rarity                 string     `db:"rarity" json:"rarity"`
	FlavorText             string     `db:"flavorText" json:"flavorText"`
	NationalPokedexNumbers []int      `db:"nationalPokedexNumbers" json:"nationalPokedexNumbers"`
	Legalities             Legalities `db:"legalities" json:"legalities"`
	Images                 Images     `db:"images" json:"images"`
}

type Traits struct {
	Name        *string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	Type        *string `db:"type" json:"type"`
	Value       *string `db:"value" json:"value"`
}

type Images struct {
	Small  *string `db:"small" json:"small"`
	Large  *string `db:"large" json:"large"`
	Symbol *string `db:"symbol" json:"symbol"`
	Logo   *string `db:"logo" json:"logo"`
}

type Legalities struct {
	Standard  string `db:"standard" json:"standard"`
	Expanded  string `db:"expanded" json:"expanded"`
	Unlimited string `db:"unlimited" json:"unlimited"`
}

type Attack struct {
	Name                string   `db:"name" json:"name"`
	Cost                []string `db:"cost" json:"cost"`
	ConvertedEnergyCost int      `db:"convertedEnergyCost" json:"convertedEnergyCost"`
	Damage              string   `db:"damage" json:"damage"`
	Description         string   `db:"description" json:"description"`
}

type Set struct {
	ID           string     `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Series       string     `db:"series" json:"series"`
	PrintedTotal int        `db:"printedTotal" json:"printedTotal"`
	Total        int        `db:"total" json:"total"`
	Legalities   Legalities `db:"legalities" json:"legalities"`
	PtcgoCode    string     `db:"ptcgoCode" json:"ptcgoCode"`
	ReleaseDate  string     `db:"releaseDate" json:"releaseDate"`
	UpdatedAt    string     `db:"updatedAt" json:"updatedAt"`
	Images       Images     `db:"images" json:"images"`
}
