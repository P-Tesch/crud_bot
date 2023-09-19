package main

import (
	"crud_bot/db/services"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/songs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
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
		}
	})

	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
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
		}
	})

	http.HandleFunc("/interpreters", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllInterpreters()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllTopics()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/subtopics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllSubtopics()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllQuestions()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllAnswers()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/scores", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllScores()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllItems()
			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/botusers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllBotusers()
			fmt.Fprintf(w, string(result))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
