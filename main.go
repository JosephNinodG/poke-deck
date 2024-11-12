package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephNinodG/poke-deck/api"
	"github.com/JosephNinodG/poke-deck/tcgapi"
)

var (
	httpPort  int
	tcgapikey string
)

func main() {
	flag.StringVar(&tcgapikey, "tcg_api_key", "", "Pokemon TCG API key")
	flag.IntVar(&httpPort, "http_port", 8080, "Http Port")
	flag.Parse()

	appname := "poke-deck"
	ctx, cancelFunc := context.WithCancel(context.Background())
	slog.InfoContext(ctx, fmt.Sprintf("App starting: %v", appname))

	tcgapi.SetUpClient(ctx, tcgapikey)

	go startHTTPServer(ctx, cancelFunc, appname)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	//Wait until a termination signal or a context cancellation
	select {
	case <-termChan:
		cancelFunc()
	case <-ctx.Done():
	}
}

func startHTTPServer(ctx context.Context, cancelFunc context.CancelFunc, appname string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("/health endpoint failed", "error", err)
			cancelFunc()
			os.Exit(1)
		}
	}()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/"+appname+"/getcard", api.GetCard)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "Error starting HTTP server", "error", err)
		}
	}
}
