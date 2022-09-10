package main

import (
	"database/sql"
	"embed"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

const ()

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	const conn = "user=akirinyuk password=pwd host=localhost database=migrations_3 port=5432 sslmode=disable"
	db, err := sql.Open("pgx", conn)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("main -- failed -- sql.Open(...)")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().
				Err(err).
				Msg("main -- failed -- sql.Close()")
		}
	}(db)

	goose.SetBaseFS(embedMigrations)
	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("main -- failed -- goose.Run(...)")
	}
}
