package main

import (
	"database/sql"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
)

func main() {
	const conn = "user=akirinyuk password=pwd host=localhost database=migrations_3 port=5432 sslmode=disable"
	db, err := sql.Open("pgx", conn)
	if err != nil {
		log.Fatal().Err(err).Msg("main -- failed -- sql.Open(...)")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Err(err).Msg("main -- failed -- sql.Close()")
		}
	}(db)

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Paths(
			"6-postgres/4-fixtures/fixtures",
		),
		testfixtures.DangerousSkipTestDatabaseCheck(), // allow non-test DB name
	)
	if err != nil {
		log.Fatal().Err(err).Msg("main -- failed -- testfixtures.New(...)")
	}

	err = fixtures.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("main -- failed -- fixtures.Load()")
	}
}
