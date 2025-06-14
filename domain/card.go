package domain

type GetCardsRequest struct {
	Card      CardDetails
	Paramters Parameters
}

type CardDetails struct {
	Name       string
	Type       string
	Supertype  string
	Subtype    string
	Set        string
	Attack     string
	Legalities Legalities
}

type Parameters struct {
	MaxCards int
	OrderBy  string
	Desc     bool
}
