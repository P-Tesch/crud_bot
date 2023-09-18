package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RetrieveAllScores() []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "SELECT * FROM scores")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Score
	for results.Next() {
		var id int64
		var musicle_total int
		var musicle_win int
		var quiz_total int
		var quiz_win int
		var tictactoe_total int
		var tictactoe_win int
		var chess_total int
		var chess_win int
		results.Scan(&id, &musicle_total, &musicle_win, &quiz_total, &quiz_win, &tictactoe_total, &tictactoe_win, &chess_total, &chess_win)
		resultSet = append(resultSet, entities.NewScore(&id, &musicle_total, &musicle_win, &quiz_total, &quiz_win, &tictactoe_total, &tictactoe_win, &chess_total, &chess_win))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}
