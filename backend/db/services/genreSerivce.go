package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5"
)

func RetrieveAllGenres() []byte {
	connection, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/postgres-bot")
	defer connection.Close(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "SELECT * FROM genres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Genre
	for results.Next() {
		var id int64
		var name string
		results.Scan(&id, &name)
		resultSet = append(resultSet, entities.NewGenre(&id, &name))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}
