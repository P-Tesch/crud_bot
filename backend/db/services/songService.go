package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RetrieveAllSongs() []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "select s.song_id, s.name, s.url, to_json(g), to_json(array_agg(i)) from songs s join genres g on g.genre_id = s.genre_id join songs_interpreters si on s.song_id = si.song_id join interpreters i on i.interpreter_id = si.interpreter_id group by s.song_id, s.name, s.url, g.genre_id, g.name")
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
