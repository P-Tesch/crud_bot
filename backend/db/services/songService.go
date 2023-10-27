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

func CreateSong(name string, url string, interpreters []entities.Interpreter, genre entities.Genre, username string, password string) error {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"?user="+username+"&password="+password)
	defer connection.Close()
	if err != nil {
		return err
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		return err
	}

	resultsSongs, err := tx.Query(context.Background(),
		"INSERT INTO songs (name, url, genre_id) "+
			"VALUES ('"+name+"', '"+url+"', '"+strconv.FormatInt(*genre.Genre_id, 10)+"') "+
			"RETURNING song_id")
	if err != nil {
		return err
	}

	var id int64
	if resultsSongs.Next() {
		resultsSongs.Scan(&id)
		resultsSongs.Close()
	}
	err = resultsSongs.Err()
	if err != nil {
		return err
	}

	for i := range interpreters {
		resultsJoin, err := tx.Query(context.Background(),
			"INSERT INTO songs_interpreters (song_id, interpreter_id) "+
				"VALUES ('"+strconv.FormatInt(id, 10)+"', '"+strconv.FormatInt(*interpreters[i].Interpreter_id, 10)+"')")
		if err != nil {
			return err
		}
		resultsJoin.Close()
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func UpdateSong(song entities.Song, username string, password string) error {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"?user="+username+"&password="+password)
	defer connection.Close()
	if err != nil {
		return err
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		return err
	}

	songId := strconv.FormatInt(*song.Song_id, 10)
	name := *song.Name
	url := *song.Url
	genreId := strconv.FormatInt(*song.Genre.Genre_id, 10)
	interpreters := *song.Interpreters

	resultsSongs, err := tx.Query(context.Background(),
		"INSERT INTO songs (song_id, name, url, genre_id) "+
			"VALUES ('"+songId+"', '"+name+"', '"+url+"', '"+genreId+"') "+
			"ON CONFLICT (song_id) DO UPDATE SET name = '"+name+"', url = '"+url+"', genre_id = '"+genreId+"'")
	if err != nil {
		return err
	}

	resultsSongs.Close()

	resultsDelete, err := tx.Query(context.Background(), "DELETE FROM songs_interpreters WHERE song_id = "+songId)
	if err != nil {
		return err
	}
	resultsDelete.Close()

	for i := range interpreters {
		resultsJoin, err := tx.Query(context.Background(),
			"INSERT INTO songs_interpreters (song_id, interpreter_id) "+
				"VALUES ('"+songId+"', '"+strconv.FormatInt(*interpreters[i].Interpreter_id, 10)+"')")
		if err != nil {
			return err
		}
		resultsJoin.Close()
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func DeleteSong(id string, username string, password string) error {
	return executeQuery("DELETE FROM songs WHERE song_id = "+id, username, password)
}

func retrieveSong(query string, username string, password string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"?user="+username+"&password="+password)
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

func RetrieveAllSongs(username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name", username, password)
}

func RetrieveSongById(songId string, username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"WHERE s.song_id = "+songId+" "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ", username, password)
}

func RetrieveSongsByGenreName(genreName string, username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"WHERE g.name iLike '"+genreName+"' "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ", username, password)
}

func RetrieveSongByInterpreterName(interpreterName string, username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"WHERE i.name iLike '"+interpreterName+"' "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ", username, password)
}

func RetrieveSongsByGenreId(genreId string, username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"WHERE g.genre_id = '"+genreId+"' "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ", username, password)
}

func RetrieveSongByInterpreterId(interpreterId string, username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"WHERE i.interpreter_id = '"+interpreterId+"' "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ", username, password)
}

func RetrieveSongByName(name string, username string, password string) []byte {
	return retrieveSong(
		"SELECT s.song_id, s.name, s.url, TO_JSON(g), TO_JSON(ARRAY_AGG(i)) FROM songs s "+
			"JOIN genres g ON g.genre_id = s.genre_id "+
			"JOIN songs_interpreters si ON s.song_id = si.song_id "+
			"JOIN interpreters i ON i.interpreter_id = si.interpreter_id "+
			"WHERE s.name iLike '"+name+"' "+
			"GROUP BY s.song_id, s.name, s.url, g.genre_id, g.name ", username, password)
}
