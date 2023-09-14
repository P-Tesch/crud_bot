package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetAllGenres() {
	connection, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/postgres-bot")
	defer connection.Close(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "SELECT * FROM music_genres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	for results.Next() {
		var name string
		var id int64
		results.Scan(&name, &id)
		fmt.Println(name)
	}
}
