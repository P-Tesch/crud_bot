package main

import (
	"crud_bot/db/handlers"
	"fmt"
	"net/http"
)

func main() {
	handlers.RegisterAllHandlers()

	http.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("./frontend"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
