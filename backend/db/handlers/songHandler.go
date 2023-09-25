package handlers

import (
	"crud_bot/db/entities"
	"crud_bot/db/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func RegisterSongHandler() {
	http.HandleFunc("/songs/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")
			genreName := urlQuery.Get("genre_name")
			interpreterName := urlQuery.Get("interpreter_name")
			genreId := urlQuery.Get("genre_id")
			interpreterId := urlQuery.Get("interpreter_id")

			if id != "" {
				result = services.RetrieveSongById(id)

			} else if name != "" {
				result = services.RetrieveSongByName(name)

			} else if genreName != "" {
				result = services.RetrieveSongsByGenreName(genreName)

			} else if interpreterName != "" {
				result = services.RetrieveSongByInterpreterName(interpreterName)

			} else if genreId != "" {
				result = services.RetrieveSongsByGenreId(genreId)

			} else if interpreterId != "" {
				result = services.RetrieveSongByInterpreterId(interpreterId)

			} else {
				result = services.RetrieveAllSongs()
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			song := new(entities.Song)
			json.Unmarshal(body, song)
			result, err := services.CreateSong(*song.Name, *song.Url, *song.Interpreters, *song.Genre)

			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"song_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "songs/")[1]
			err := services.DeleteSong(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
