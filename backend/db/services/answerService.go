package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func retrieveAnswer(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Answer
	for results.Next() {
		var id int64
		var answer string
		var correct bool
		var question_id int64

		results.Scan(&id, &answer, &correct, &question_id)
		resultSet = append(resultSet, entities.NewAnswer(&id, &answer, &correct, &question_id))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}

func RetrieveAllAnswers() []byte {
	return retrieveAnswer("SELECT * FROM answers")
}

func RetrieveAnswerById(id string) []byte {
	return retrieveAnswer(
		"SELECT * FROM answers a " +
			"WHERE a.answer_id = " + id)
}

func RetrieveAnswerByQuestionId(questionId string) []byte {
	return retrieveAnswer(
		"SELECT * FROM answers a " +
			"WHERE a.question_id = " + questionId)
}
