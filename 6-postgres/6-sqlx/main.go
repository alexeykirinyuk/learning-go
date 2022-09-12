package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
)

type product struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	db, err := sqlx.Open("pgx", "postgres://akirinyuk@localhost:5432/migrations_3")
	if err != nil {
		log.Fatal().Err(err).Msg("sqlx.Open() error")
	}

	ctx := context.Background()
	fmt.Println()
	fmt.Println("GetContext")

	{
		sql := `SELECT p.id, p.name, p.created_at
				FROM products p
				WHERE p.id = :product_id`
		query, args, err := sqlx.Named(sql, map[string]interface{}{
			"product_id": 4,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("sqlx.Named() error")
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query)

		p := product{}
		err = db.GetContext(ctx, &p, query, args...)
		if err != nil {
			log.Fatal().Err(err).Msg("db.GetContext() error")
		}

		fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
	}

	fmt.Println()
	fmt.Println("GetContext + sqlx.in")

	{
		sql := `SELECT p.id, p.name, p.created_at
				FROM products p
				WHERE p.id in (:product_ids)`

		productIDs := []int64{1, 3, 5}
		query, args, err := sqlx.Named(sql, map[string]interface{}{
			"product_ids": productIDs,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("sqlx.Named() error")
		}

		query, args, err = sqlx.In(query, args...)
		if err != nil {
			log.Fatal().Err(err).Msg("sqlx.In() error")
		}

		query = db.Rebind(query)

		p := product{}
		err = db.GetContext(ctx, &p, query, args...)
		if err != nil {
			log.Fatal().Err(err).Msg("db.GetContext() error")
		}

		fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
	}

	fmt.Println()
	fmt.Println("GetContext + NamedQueryContext")

	{
		sql := `SELECT p.id, p.name, p.created_at
				FROM products p
				WHERE p.category_id = :category_id`

		rows, err := db.NamedQueryContext(ctx, sql, map[string]interface{}{
			"category_id": 2,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("db.NamedQueryContext() error")
		}

		var products []product
		for rows.Next() {
			var p product
			err := rows.StructScan(&p)
			if err != nil {
				log.Fatal().Err(err).Msg("rows.StructScan() error")
			}

			products = append(products, p)
		}

		for _, p := range products {
			fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
		}
	}

	fmt.Println()
	fmt.Println("SelectContext + slice scan")

	{
		sql := `SELECT p.id, p.name
				FROM products p
				ORDER BY p.created_at`

		var products []product
		err = db.SelectContext(ctx, &products, sql)
		if err != nil {
			log.Fatal().Err(err).Msg("db.SelectContext()")
		}

		for _, p := range products {
			fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
		}
	}
}
