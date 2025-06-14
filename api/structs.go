package api

import "reflect"

type GetUserCollectionRequest struct {
	UserID       int
	CollectionID int
}

func (r *GetUserCollectionRequest) IsValid() bool {
	return r.UserID > 0 && r.CollectionID > 0
}

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

type Legalities struct {
	Standard  string `db:"standard" json:"standard"`
	Expanded  string `db:"expanded" json:"expanded"`
	Unlimited string `db:"unlimited" json:"unlimited"`
}

type Parameters struct {
	MaxCards int    `json:"maxCards"`
	OrderBy  string `json:"orderBy"`
	Desc     bool   `json:"desc"`
}

func (r GetCardsRequest) IsValid() (bool, string) {
	if reflect.ValueOf(r.Card).IsZero() {
		return false, "invalid request. At least one value in the card object must be provided"
	}

	legalities := r.Card.Legalities

	if legalities.Standard != "" {
		if legalities.Standard != "legal" && legalities.Standard != "banned" {
			return false, "invalid value provided for standard legality"
		}
	}

	if legalities.Expanded != "" {
		if legalities.Expanded != "legal" && legalities.Expanded != "banned" {
			return false, "invalid value provided for expanded legality"
		}
	}

	if legalities.Unlimited != "" {
		if legalities.Unlimited != "legal" && legalities.Unlimited != "banned" {
			return false, "invalid value provided for unlimited legality"
		}
	}

	if r.Paramters.MaxCards < 0 || r.Paramters.MaxCards > 250 { //This can be 0 as the tcgapi sets a default of 0 to 250
		return false, "invalid value provided for maximum number of cards to return. Must be between 0 - 250"
	}

	if r.Paramters.OrderBy != "" && r.Paramters.OrderBy != "name" && r.Paramters.OrderBy != "number" && r.Paramters.OrderBy != "set" {
		return false, "invalid value provided to order cards by. Must be 'name', 'number', or 'set'"
	}

	return true, ""
}
