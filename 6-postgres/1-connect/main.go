package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println()

	db, err := sql.Open("pgx", "postgres://akirinyuk:pwd@localhost:5432/db")
	if err != nil {
		log.Fatal().Err(err).Msg("main -- failed -- sql.Open(\"pgx\")")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().
				Err(err).
				Msg("main -- failed -- error when db.Close()")
		}
	}(db)

	ctx := context.Background()

	// только сейчас создается подключение
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("main -- failed -- db.PingContext(ctx)")
	}

	fmt.Println("done")
}
