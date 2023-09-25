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

func RegisterGenreHandler() {
	http.HandleFunc("/genres/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")

			if id != "" {
				result = services.RetrieveGenreById(id)

			} else if name != "" {
				result = services.RetrieveGenreByName(name)

			} else {
				result = services.RetrieveAllGenres()
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			genre := new(entities.Genre)
			json.Unmarshal(body, genre)
			result, err := services.CreateGenre(*genre.Name)

			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"genre_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "genres/")[1]
			err := services.DeleteGenre(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
