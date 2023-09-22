package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateScore(musicle_total int, musicle_win int, quiz_total int, quiz_win int, tictactoe_total int, tictactoe_win int, chess_total int, chess_win int) (int64, error) {
	return createGeneric(
		"INSERT INTO scores (musicle_total, musicle_win, quiz_total, quiz_win, tictactoe_total, tictactoe_win, chess_total, chess_win)" +
			"VALUES (" +
			"'" + strconv.Itoa(musicle_total) + "', " +
			"'" + strconv.Itoa(musicle_win) + "', " +
			"'" + strconv.Itoa(quiz_total) + "', " +
			"'" + strconv.Itoa(quiz_win) + "', " +
			"'" + strconv.Itoa(tictactoe_total) + "', " +
			"'" + strconv.Itoa(tictactoe_win) + "', " +
			"'" + strconv.Itoa(chess_total) + "', " +
			"'" + strconv.Itoa(chess_win) + "') " +
			"RETURNING score_id")
}

func retrieveScore(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
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

func RetrieveAllScores() []byte {
	return retrieveScore("SELECT * FROM scores")
}

func RetrieveScoreById(id string) []byte {
	return retrieveScore(
		"SELECT * FROM scores s " +
			"WHERE s.score_id = " + id)
}

func RetrieveScoreByBotuserId(botuserId string) []byte {
	return retrieveScore(
		"SELECT * FROM scores s " +
			"JOIN botusers b ON s.score_id = b.score_id " +
			"WHERE b.botuser_id = " + botuserId)
}
