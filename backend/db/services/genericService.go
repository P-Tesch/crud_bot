package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createGeneric(query string) int64 {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return 0
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to begin transaction: %v\n", err)
		return 0
	}

	results, err := tx.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
		return 0
	}

	results.Next()
	var id int64
	results.Scan(&id)
	results.Close()

	err = tx.Commit(context.Background())
	return id
}
