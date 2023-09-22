package services

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createGeneric(query string) (int64, error) {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		return 0, err
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		return 0, err
	}

	results, err := tx.Query(context.Background(), query)
	if err != nil {
		return 0, err
	}

	var id int64
	if results.Next() {
		results.Scan(&id)
		results.Close()
	}
	err = results.Err()
	if err != nil {
		return 0, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, err
	}
	return id, nil
}
