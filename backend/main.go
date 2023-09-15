package main

import (
	"crud_bot/db/services"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.GetAllSongs()
			fmt.Fprintf(w, string(result))
		}
	})
	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.GetAllGenres()
			fmt.Fprintf(w, string(result))
		}
	})
	http.HandleFunc("/interpreters", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.GetAllInterpreters()
			fmt.Fprintf(w, string(result))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
