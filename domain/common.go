package domain

type PokemonCard struct {
	ID                     string     `json:"id"`
	Name                   string     `json:"name"`
	Supertype              string     `json:"supertype"`
	Subtypes               []string   `json:"subtypes"`
	Level                  string     `json:"level"`
	Hp                     string     `json:"hp"`
	Types                  []string   `json:"types"`
	EvolvesFrom            string     `json:"evolvesFrom"`
	EvolvesTo              []string   `json:"evolvesTo"`
	Rules                  []string   `json:"rules"`
	AncientTrait           *Traits    `json:"ancientTrait"`
	Abilities              []Traits   `json:"abilities"`
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

// type PokemonCard struct {
// 	ID           string   `json:"id"`
// 	Name         string   `json:"name"`
// 	Supertype    string   `json:"supertype"`
// 	Subtypes     []string `json:"subtypes"`
// 	Level        string   `json:"level"`
// 	Hp           string   `json:"hp"`
// 	Types        []string `json:"types"`
// 	EvolvesFrom  string   `json:"evolvesFrom"`
// 	EvolvesTo    []string `json:"evolvesTo"`
// 	Rules        []string `json:"rules"`
// 	AncientTrait *struct {
// 		Name string `json:"name"`
// 		description string `json:"description"`
// 	} `json:"ancientTrait"`
// 	Abilities []struct {
// 		Name string `json:"name"`
// 		description string `json:"description"`
// 		Type string `json:"type"`
// 	} `json:"abilities"`
// 	Attacks []struct {
// 		Name                string   `json:"name"`
// 		Cost                []string `json:"cost"`
// 		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
// 		Damage              string   `json:"damage"`
// 		description                string   `json:"description"`
// 	} `json:"attacks"`
// 	Weaknesses []struct {
// 		Type  string `json:"type"`
// 		Value string `json:"value"`
// 	} `json:"weaknesses"`
// 	Resistances []struct {
// 		Type  string `json:"type"`
// 		Value string `json:"value"`
// 	} `json:"resistances"`
// 	RetreatCost          []string `json:"retreatCost"`
// 	ConvertedRetreatCost int      `json:"convertedRetreatCost"`
// 	Set                  struct {
// 		ID           string `json:"id"`
// 		Name         string `json:"name"`
// 		Series       string `json:"series"`
// 		PrintedTotal int    `json:"printedTotal"`
// 		Total        int    `json:"total"`
// 		Legalities   struct {
// 			Unlimited string `json:"unlimited"`
// 		} `json:"legalities"`
// 		PtcgoCode   string `json:"ptcgoCode"`
// 		ReleaseDate string `json:"releaseDate"`
// 		UpdatedAt   string `json:"updatedAt"`
// 		Images      struct {
// 			Symbol string `json:"symbol"`
// 			Logo   string `json:"logo"`
// 		} `json:"images"`
// 	} `json:"set"`
// 	Number                 string `json:"number"`
// 	Artist                 string `json:"artist"`
// 	Rarity                 string `json:"rarity"`
// 	Flavordescription             string `json:"flavordescription"`
// 	NationalPokedexNumbers []int  `json:"nationalPokedexNumbers"`
// 	Legalities             struct {
// 		Unlimited string `json:"unlimited"`
// 	} `json:"legalities"`
// 	Images struct {
// 		Small string `json:"small"`
// 		Large string `json:"large"`
// 	} `json:"images"`
// 	TCGPlayer struct {
// 		URL       string `json:"url"`
// 		UpdatedAt string `json:"updatedAt"`
// 		Prices    struct {
// 			Holofoil *struct {
// 				Low    float64 `json:"low"`
// 				Mid    float64 `json:"mid"`
// 				High   float64 `json:"high"`
// 				Market float64 `json:"market"`
// 			} `json:"holofoil,omitempty"`
// 			ReverseHolofoil *struct {
// 				Low    float64 `json:"low"`
// 				Mid    float64 `json:"mid"`
// 				High   float64 `json:"high"`
// 				Market float64 `json:"market"`
// 			} `json:"reverseHolofoil,omitempty"`
// 			Normal *struct {
// 				Low    float64 `json:"low"`
// 				Mid    float64 `json:"mid"`
// 				High   float64 `json:"high"`
// 				Market float64 `json:"market"`
// 			} `json:"normal,omitempty"`
// 		} `json:"prices"`
// 	} `json:"tcgplayer"`
// 	CardMarket struct {
// 		URL       string `json:"url"`
// 		UpdatedAt string `json:"updatedAt"`
// 		Prices    struct {
// 			AverageSellPrice *float64 `json:"averageSellPrice"`
// 			LowPrice         *float64 `json:"lowPrice"`
// 			TrendPrice       *float64 `json:"trendPrice"`
// 			GermanProLow     *float64 `json:"germanProLow"`
// 			SuggestedPrice   *float64 `json:"suggestedPrice"`
// 			ReverseHoloSell  *float64 `json:"reverseHoloSell"`
// 			ReverseHoloLow   *float64 `json:"reverseHoloLow"`
// 			ReverseHoloTrend *float64 `json:"reverseHoloTrend"`
// 			LowPriceExPlus   *float64 `json:"lowPriceExPlus"`
// 			Avg1             *float64 `json:"avg1"`
// 			Avg7             *float64 `json:"avg7"`
// 			Avg30            *float64 `json:"avg30"`
// 			ReverseHoloAvg1  *float64 `json:"reverseHoloAvg1"`
// 			ReverseHoloAvg7  *float64 `json:"reverseHoloAvg7"`
// 			ReverseHoloAvg30 *float64 `json:"reverseHoloAvg30"`
// 		} `json:"prices"`
// 	} `json:"cardmarket"`
// }
