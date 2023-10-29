package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateGenre(name string, username string, password string) error {
	return executeQuery(
		"INSERT INTO genres (name) "+
			"VALUES ('"+name+"') ", username, password)
}

func DeleteGenre(id string, username string, password string) error {
	return executeQuery("DELETE FROM genres WHERE genre_id = "+id, username, password)
}

func UpdateGenre(genre entities.Genre, username string, password string) error {
	return executeQuery(
		"INSERT INTO genres (genre_id, name) "+
			"VALUES ('"+strconv.FormatInt(*genre.Genre_id, 10)+"', '"+*genre.Name+"') "+
			"ON CONFLICT (genre_id) DO UPDATE SET name = '"+*genre.Name+"'", username, password)
}

func retrieveGenre(query string, username string, password string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"&user="+username+"&password="+password)
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

func RetrieveAllGenres(username string, password string) []byte {
	return retrieveGenre("SELECT * FROM genres", username, password)
}

func RetrieveGenreById(id string, username string, password string) []byte {
	return retrieveGenre("SELECT * FROM genres g WHERE g.genre_id = "+id, username, password)
}

func RetrieveGenreByName(name string, username string, password string) []byte {
	return retrieveGenre("SELECT * FROM genres g WHERE g.name iLike '"+name+"'", username, password)
}
