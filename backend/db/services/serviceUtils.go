package services

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func executeQuery(query string, username string, password string) error {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"&user="+username+"&password="+password)
	defer connection.Close()
	if connection.Ping(context.Background()) != nil {
		return errors.New("ERROR: connection ping failed(Authentication may have failed)")
	}
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
