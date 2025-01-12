package domain

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

// type PokemonCard struct {
// 	ID           string   `db:"" json:"id"`
// 	Name         string   `db:"" json:"name"`
// 	Supertype    string   `db:"" json:"supertype"`
// 	Subtypes     []string `db:"" json:"subtypes"`
// 	Level        string   `db:"" json:"level"`
// 	Hp           string   `db:"" json:"hp"`
// 	Types        []string `db:"" json:"types"`
// 	EvolvesFrom  string   `db:"" json:"evolvesFrom"`
// 	EvolvesTo    []string `db:"" json:"evolvesTo"`
// 	Rules        []string `db:"" json:"rules"`
// 	AncientTrait *struct {
// 		Name string `db:"" json:"name"`
// 		description string `db:"" json:"description"`
// 	} `db:"" json:"ancientTrait"`
// 	Abilities []struct {
// 		Name string `db:"" json:"name"`
// 		description string `db:"" json:"description"`
// 		Type string `db:"" json:"type"`
// 	} `db:"" json:"abilities"`
// 	Attacks []struct {
// 		Name                string   `db:"" json:"name"`
// 		Cost                []string `db:"" json:"cost"`
// 		ConvertedEnergyCost int      `db:"" json:"convertedEnergyCost"`
// 		Damage              string   `db:"" json:"damage"`
// 		description                string   `db:"" json:"description"`
// 	} `db:"" json:"attacks"`
// 	Weaknesses []struct {
// 		Type  string `db:"" json:"type"`
// 		Value string `db:"" json:"value"`
// 	} `db:"" json:"weaknesses"`
// 	Resistances []struct {
// 		Type  string `db:"" json:"type"`
// 		Value string `db:"" json:"value"`
// 	} `db:"" json:"resistances"`
// 	RetreatCost          []string `db:"" json:"retreatCost"`
// 	ConvertedRetreatCost int      `db:"" json:"convertedRetreatCost"`
// 	Set                  struct {
// 		ID           string `db:"" json:"id"`
// 		Name         string `db:"" json:"name"`
// 		Series       string `db:"" json:"series"`
// 		PrintedTotal int    `db:"" json:"printedTotal"`
// 		Total        int    `db:"" json:"total"`
// 		Legalities   struct {
// 			Unlimited string `db:"" json:"unlimited"`
// 		} `db:"" json:"legalities"`
// 		PtcgoCode   string `db:"" json:"ptcgoCode"`
// 		ReleaseDate string `db:"" json:"releaseDate"`
// 		UpdatedAt   string `db:"" json:"updatedAt"`
// 		Images      struct {
// 			Symbol string `db:"" json:"symbol"`
// 			Logo   string `db:"" json:"logo"`
// 		} `db:"" json:"images"`
// 	} `db:"" json:"set"`
// 	Number                 string `db:"" json:"number"`
// 	Artist                 string `db:"" json:"artist"`
// 	Rarity                 string `db:"" json:"rarity"`
// 	Flavordescription             string `db:"" json:"flavordescription"`
// 	NationalPokedexNumbers []int  `db:"" json:"nationalPokedexNumbers"`
// 	Legalities             struct {
// 		Unlimited string `db:"" json:"unlimited"`
// 	} `db:"" json:"legalities"`
// 	Images struct {
// 		Small string `db:"" json:"small"`
// 		Large string `db:"" json:"large"`
// 	} `db:"" json:"images"`
// 	TCGPlayer struct {
// 		URL       string `db:"" json:"url"`
// 		UpdatedAt string `db:"" json:"updatedAt"`
// 		Prices    struct {
// 			Holofoil *struct {
// 				Low    float64 `db:"" json:"low"`
// 				Mid    float64 `db:"" json:"mid"`
// 				High   float64 `db:"" json:"high"`
// 				Market float64 `db:"" json:"market"`
// 			} `db:"" json:"holofoil,omitempty"`
// 			ReverseHolofoil *struct {
// 				Low    float64 `db:"" json:"low"`
// 				Mid    float64 `db:"" json:"mid"`
// 				High   float64 `db:"" json:"high"`
// 				Market float64 `db:"" json:"market"`
// 			} `db:"" json:"reverseHolofoil,omitempty"`
// 			Normal *struct {
// 				Low    float64 `db:"" json:"low"`
// 				Mid    float64 `db:"" json:"mid"`
// 				High   float64 `db:"" json:"high"`
// 				Market float64 `db:"" json:"market"`
// 			} `db:"" json:"normal,omitempty"`
// 		} `db:"" json:"prices"`
// 	} `db:"" json:"tcgplayer"`
// 	CardMarket struct {
// 		URL       string `db:"" json:"url"`
// 		UpdatedAt string `db:"" json:"updatedAt"`
// 		Prices    struct {
// 			AverageSellPrice *float64 `db:"" json:"averageSellPrice"`
// 			LowPrice         *float64 `db:"" json:"lowPrice"`
// 			TrendPrice       *float64 `db:"" json:"trendPrice"`
// 			GermanProLow     *float64 `db:"" json:"germanProLow"`
// 			SuggestedPrice   *float64 `db:"" json:"suggestedPrice"`
// 			ReverseHoloSell  *float64 `db:"" json:"reverseHoloSell"`
// 			ReverseHoloLow   *float64 `db:"" json:"reverseHoloLow"`
// 			ReverseHoloTrend *float64 `db:"" json:"reverseHoloTrend"`
// 			LowPriceExPlus   *float64 `db:"" json:"lowPriceExPlus"`
// 			Avg1             *float64 `db:"" json:"avg1"`
// 			Avg7             *float64 `db:"" json:"avg7"`
// 			Avg30            *float64 `db:"" json:"avg30"`
// 			ReverseHoloAvg1  *float64 `db:"" json:"reverseHoloAvg1"`
// 			ReverseHoloAvg7  *float64 `db:"" json:"reverseHoloAvg7"`
// 			ReverseHoloAvg30 *float64 `db:"" json:"reverseHoloAvg30"`
// 		} `db:"" json:"prices"`
// 	} `db:"" json:"cardmarket"`
// }
