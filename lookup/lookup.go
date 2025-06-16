package lookup

import (
	"strings"

	"github.com/JosephNinodG/poke-deck/domain"
	"github.com/JosephNinodG/poke-deck/handler"
)

const (
	AllLegal                = 1
	StandardBanned          = 2
	ExpandedBanned          = 3
	UnlimitedBanned         = 4
	StandardExpandedBanned  = 5
	ExpandedUnlimitedBanned = 6
	StandardUnlimitedBanned = 7
	AllBanned               = 8
)

var (
	databaseHandler handler.Database
)

func Configure(databaseHandlerOverride handler.Database) {
	if databaseHandler != nil {
		panic("databaseHandler instance already set")
	}

	databaseHandler = databaseHandlerOverride
}

func MapLegality(legalities domain.Legalities) int {
	standard := strings.ToLower(legalities.Standard)
	expanded := strings.ToLower(legalities.Expanded)
	unlimited := strings.ToLower(legalities.Unlimited)

	if standard == "banned" && expanded == "legal" && unlimited == "legal" {
		return StandardBanned
	}

	if standard == "legal" && expanded == "banned" && unlimited == "legal" {
		return ExpandedBanned
	}

	if standard == "legal" && expanded == "legal" && unlimited == "banned" {
		return UnlimitedBanned
	}

	if standard == "banned" && expanded == "banned" && unlimited == "legal" {
		return StandardExpandedBanned
	}

	if standard == "legal" && expanded == "banned" && unlimited == "banned" {
		return ExpandedUnlimitedBanned
	}

	if standard == "banned" && expanded == "legal" && unlimited == "banned" {
		return StandardUnlimitedBanned
	}

	if standard == "banned" && expanded == "banned" && unlimited == "banned" {
		return AllBanned
	}

	return AllLegal

}
