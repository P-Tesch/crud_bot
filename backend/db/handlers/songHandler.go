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
		username, password, _ := r.BasicAuth()
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
				result = services.RetrieveSongById(id, username, password)

			} else if name != "" {
				result = services.RetrieveSongByName(name, username, password)

			} else if genreName != "" {
				result = services.RetrieveSongsByGenreName(genreName, username, password)

			} else if interpreterName != "" {
				result = services.RetrieveSongByInterpreterName(interpreterName, username, password)

			} else if genreId != "" {
				result = services.RetrieveSongsByGenreId(genreId, username, password)

			} else if interpreterId != "" {
				result = services.RetrieveSongByInterpreterId(interpreterId, username, password)

			} else {
				result = services.RetrieveAllSongs(username, password)
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			song := new(entities.Song)
			json.Unmarshal(body, song)
			err = services.CreateSong(*song.Name, *song.Url, *song.Interpreters, *song.Genre, username, password)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "songs/")[1]
			err := services.DeleteSong(id, username, password)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "PUT":
			id := strings.Split(r.URL.Path, "songs/")[1]
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			song := new(entities.Song)
			err = json.Unmarshal(body, song)
			if err != nil {
				fmt.Fprintf(w, "Error unmarshalling body: %v\n", err)
			}

			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error converting id: Error parsing from string to int64")
			}

			if *song.Song_id != idInt {
				fmt.Fprintf(w, "Error in path: Path id does not match object id")
			}

			err = services.UpdateSong(*song, username, password)
			if err == nil {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
