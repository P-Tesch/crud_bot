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

func retrieve(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Song
	for results.Next() {
		var id int64
		var name string
		var url string
		var genre *entities.Genre
		var interpreters *[]entities.Interpreter
		var genreByte []byte
		var interpretersByte []byte

		results.Scan(&id, &name, &url, &genre, &interpreters)
		json.Unmarshal(genreByte, genre)
		json.Unmarshal(interpretersByte, interpreters)
		resultSet = append(resultSet, entities.NewSong(&id, &name, &url, interpreters, genre))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}

func RetrieveAllSongs() []byte {
	return retrieve(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name")
}

func RetrieveSongById(id int64) []byte {
	return retrieve(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE s.song_id = " + strconv.FormatInt(id, 10) + " " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongsByGenre(genre string) []byte {
	return retrieve(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE g.name iLike '" + genre + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongByInterpreter(interpreter string) []byte {
	return retrieve(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE i.name iLike '" + interpreter + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongByName(name string) []byte {
	return retrieve(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE s.name iLike '" + name + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}
