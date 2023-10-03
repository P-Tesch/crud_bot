package services

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func executeQuery(query string) error {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		return err
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		return err
	}

	results, err := tx.Query(context.Background(), query)
	if err != nil {
		return err
	}

	results.Close()
	err = results.Err()
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}
