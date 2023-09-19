package main

import (
	"crud_bot/db/services"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/songs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")
			genre := urlQuery.Get("genre")
			interpreter := urlQuery.Get("interpreter")

			if id != "" {
				idInt, err := strconv.ParseInt(id, 10, 64)

				if err != nil {
					fmt.Fprintf(os.Stderr, "Unable to parse int: %v\n", err)
				}

				result = services.RetrieveSongById(idInt)

			} else if name != "" {
				result = services.RetrieveSongByName(name)

			} else if genre != "" {
				result = services.RetrieveSongsByGenre(genre)

			} else if interpreter != "" {
				result = services.RetrieveSongByInterpreter(interpreter)

			} else {
				result = services.RetrieveAllSongs()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllGenres()
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
