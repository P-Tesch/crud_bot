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

func CreateSong(name string, url string, interpreters []entities.Interpreter, genre entities.Genre) int64 {
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

	resultsSongs, err := tx.Query(context.Background(), "INSERT INTO songs (name, url, genre_id) VALUES ('"+name+"', '"+url+"', '"+strconv.FormatInt(*genre.Genre_id, 10)+"') RETURNING song_id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
		return 0
	}

	resultsSongs.Next()
	var id int64
	resultsSongs.Scan(&id)
	resultsSongs.Close()

	for i := range interpreters {
		resultsJoin, err := tx.Query(context.Background(), "INSERT INTO songs_interpreters (song_id, interpreter_id) VALUES ('"+strconv.FormatInt(id, 10)+"', '"+strconv.FormatInt(*interpreters[i].Interpreter_id, 10)+"')")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
			return 0
		}
		resultsJoin.Close()
	}

	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
		return 0
	}
	return id
}

func retrieveSong(query string) []byte {
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
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name")
}

func RetrieveSongById(songId string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE s.song_id = " + songId + " " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongsByGenreName(genreName string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE g.name iLike '" + genreName + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongByInterpreterName(interpreterName string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE i.name iLike '" + interpreterName + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongsByGenreId(genreId string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE g.genre_id = '" + genreId + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongByInterpreterId(interpreterId string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE i.interpreter_id = '" + interpreterId + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}

func RetrieveSongByName(name string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s " +
			"JOIN genres g ON g.genre_id = s.genre_id " +
			"JOIN songs_interpreters si ON s.song_id = si.song_id " +
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id " +
			"WHERE s.name iLike '" + name + "' " +
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ")
}
