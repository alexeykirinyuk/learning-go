package main

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func main() {
	db, err := sqlx.Open("pgx", "postgres://akirinyuk@localhost:5432/migrations_3")
	if err != nil {
		log.Fatal().Err(err).Msg("sqlx.Open() error")
	}

	ctx := context.Background()
	fmt.Println()

	{
		sb := psql.Select("id", "name").
			From("products")

		productIDs := []int64{1, 3, 4}
		if len(productIDs) != 0 {
			sb = sb.Where(sq.Eq{"id": productIDs})
		}

		and := sq.Or{
			sq.Eq{"name": "Audi"},       // name = "Audi"
			sq.Like{"name": "%cap"},     // name LIKE "%cap"
			sq.NotEq{"created_at": nil}, // created_at is not null
		}
		sb = sb.Where(and)

		query, args, err := sb.ToSql()
		if err != nil {
			log.Fatal().Err(err).Msg("sb.ToSql()")
		}

		fmt.Println(query, args)

		type product struct {
			ID   int64  `db:"id"`
			Name string `db:"name"`
		}

		var products []product
		err = db.SelectContext(ctx, &products, query, args...)
		if err != nil {
			log.Fatal().Err(err).Msg("db.SelectContext()")
		}

		for _, p := range products {
			fmt.Println(p)
		}
	}

	{

	}
}
