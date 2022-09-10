package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
)

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

	ctx := context.Background()

	{ // Выполнение запросов
		const query = `INSERT INTO products(name, category_id, created_at)
					   VALUES ($1, $2, now())`

		result, err := db.ExecContext(ctx, query, "Baget", 4)
		if err != nil {
			log.Fatal().Err(err).Msg("main -- failed -- db.ExecContext(...)")
		}

		// result.LastInsertId() - не поддерживается

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal().Err(err).Msg("main -- failed -- result.RowsAffected()")
		}

		fmt.Println("Rows affected", rowsAffected)
	}

	fmt.Println()

	{ // Получение одной строки
		const query = `SELECT p.id, p.name
					   FROM products p
					   WHERE p.id = $1`

		row := db.QueryRowContext(ctx, query, 1)
		var (
			id   int64
			name string
		)
		err := row.Scan(&id, &name)
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("row is not found")
		} else if err != nil {
			log.Fatal().Err(err).Msg("main -- failed -- db.QueryRowContext(...)")
		}

		fmt.Println("id", id, "name", name)
	}

	fmt.Println()

	{ // Выполнение запросов с возвращением значений
		const query = `INSERT INTO products(name, category_id, created_at)
					   VALUES ($1, $2, now())
					   RETURNING id`

		row := db.QueryRowContext(ctx, query, "Baget", 2)
		if err != nil {
			log.Fatal().Err(err).Msg("main -- failed -- db.ExecContext(...)")
		}
		var id int64
		err := row.Scan(&id)
		if err != nil {
			log.Fatal().Err(err).Msg("main -- failed -- db.QueryRowContext(...)")
		}

		fmt.Println("new record id", id)
	}

	fmt.Println()

	{ // Получение нескольких строк
		const query = `SELECT p.id, p.name FROM products p WHERE category_id = $1`

		rows, err := db.QueryContext(ctx, query, 1)
		if err != nil {
			log.Fatal().Err(err).Msg("main -- failed -- db.QueryContext(...)")
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Error().Err(err).Msg("main -- failed -- rows.Close()")
			}
		}(rows)

		type product struct {
			id   int64
			name string
		}

		var products []product

		for rows.Next() {
			var p product
			err := rows.Scan(&p.id, &p.name)
			if err != nil {
				log.Fatal().Err(err).Msg("main -- failed -- rows.Scan(...)")
			}
			products = append(products, p)
		}

		for _, p := range products {
			fmt.Printf("%d - %s\n", p.id, p.name)
		}
	}
}
