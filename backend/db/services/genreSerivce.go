package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateGenre(name string) (int64, error) {
	return createGeneric(
		"INSERT INTO genres (name) " +
			"VALUES ('" + name + "') " +
			"RETURNING genre_id")
}

func DeleteGenre(id string) error {
	return deleteGeneric("DELETE FROM genres WHERE genre_id = " + id)
}

func retrieveGenre(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
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

func RetrieveAllGenres() []byte {
	return retrieveGenre("SELECT * FROM genres")
}

func RetrieveGenreById(id string) []byte {
	return retrieveGenre("SELECT * FROM genres g WHERE g.genre_id = " + id)
}

func RetrieveGenreByName(name string) []byte {
	return retrieveGenre("SELECT * FROM genres g WHERE g.name iLike '" + name + "'")
}
