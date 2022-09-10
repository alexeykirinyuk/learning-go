package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://akirinyuk:pwd@localhost:5432/db")
	if err != nil {
		log.Fatal().Err(err).Msg("main -- failed -- pgx.Connect")
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Error().
				Err(err).
				Msg("main -- failed -- conn.Close(ctx)")
		}
	}(conn, ctx)

	// только сейчас создается подключение
	var pgVersion string
	err = conn.QueryRow(ctx, "SELECT version();").Scan(&pgVersion)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("main -- failed -- db.PingContext(ctx)")
	}

	fmt.Printf("Postgres version: %v\n", pgVersion)
}
