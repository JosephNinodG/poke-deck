package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

type Connection struct {
	Host       string
	Port       int
	DbUser     string
	DbPassword string
	DbName     string
}

var client *sql.DB

func (c *Connection) NewClient(ctx context.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.DbUser, c.DbPassword, c.DbName)

	var err error

	client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = client.Ping()
	if err != nil {
		panic(err)
	}

	slog.InfoContext(ctx, "db connection initialised")
}

func CloseClient(ctx context.Context) {
	err := client.Close()
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		panic(err)
	}
}
