package main

import (
	"crud_bot/db/handlers"
	"crud_bot/db/migration"
	"fmt"
	"net/http"
)

func main() {
	err := migration.Migrate()
	if err != nil {
		fmt.Println(err)
	}
	handlers.RegisterAllHandlers()

	http.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("./assets/frontend"))))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
