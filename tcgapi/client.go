package tcgapi

import (
	"context"
	"log/slog"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
)

var client tcg.Client

func SetUpClient(ctx context.Context, apikey string) {
	client = tcg.NewClient(apikey)
	slog.InfoContext(ctx, "New TCG API client created")
}
