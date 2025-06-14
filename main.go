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
	"github.com/JosephNinodG/poke-deck/db"
	"github.com/JosephNinodG/poke-deck/handler"
	"github.com/JosephNinodG/poke-deck/tcgapi"
)

var (
	httpPort  int
	tcgapikey string
	localDev  bool

	dbhost     string
	dbport     int
	dbuser     string
	dbpassword string
	dbname     string
)

func main() {
	flag.StringVar(&tcgapikey, "tcg_api_key", "", "Pokemon TCG API key")
	flag.IntVar(&httpPort, "http_port", 8080, "Http Port")
	flag.BoolVar(&localDev, "local_dev", true, "Local Dev")
	flag.StringVar(&dbhost, "db_host", "localhost", "Database Connection Host")
	flag.IntVar(&dbport, "db_port", 5432, "Database Connection Port")
	flag.StringVar(&dbuser, "db_user", "postgres", "Database Connection User")
	flag.StringVar(&dbpassword, "db_password", "postgres", "Database Connection Password")
	flag.StringVar(&dbname, "db_name", "pokedeck", "Database Name")
	flag.Parse()

	var dbconnection = db.Connection{
		Host:       dbhost,
		Port:       dbport,
		DbUser:     dbuser,
		DbPassword: dbpassword,
		DbName:     dbname,
	}

	appname := "poke-deck"
	ctx, cancelFunc := context.WithCancel(context.Background())
	slog.InfoContext(ctx, fmt.Sprintf("App starting: %v", appname))

	dbconnection.NewClient(ctx)
	api.Configure(handler.TcgApiHandler{Apikey: tcgapikey}, handler.DatabaseHandler{})
	tcgapi.SetUpClient(ctx, tcgapikey)

	go startHTTPServer(ctx, cancelFunc, appname)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	//Wait until a termination signal or a context cancellation
	select {
	case <-termChan:
		db.CloseClient(ctx)
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

	http.HandleFunc("/"+appname+"/getcards", api.GetCards)
	http.HandleFunc("/"+appname+"/getcardbyid", api.GetCardById)
	http.HandleFunc("/"+appname+"/getusercollection", api.GetUserCollection)
	http.HandleFunc("/"+appname+"/createusercollection", api.CreateUserCollection)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "Error starting HTTP server", "error", err)
		}
	}
}
