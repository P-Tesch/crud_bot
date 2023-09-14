package main

import (
	"crud_bot/db/services"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, mundo!")
	})
	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Songs")
	})
	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		result := services.GetAllGenres()
		fmt.Fprintf(w, string(result))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
