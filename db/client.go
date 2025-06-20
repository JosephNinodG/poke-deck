package db

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Host       string
	Port       int
	DbUser     string
	DbPassword string
	DbName     string
}

var conn *pgx.Conn

func (c *Connection) NewConnection(ctx context.Context) {

	var err error
	databaseConnection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.DbUser, c.DbPassword, c.Host, strconv.Itoa(c.Port), c.DbName)

	conn, err = pgx.Connect(ctx, databaseConnection)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}

	slog.InfoContext(ctx, "db connection initialised")
}

func CloseConnection(ctx context.Context) {
	err := conn.Close(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		panic(err)
	}
}
