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

func CreateAnswer(answer string, correct bool, question_id int64) error {
	return executeQuery(
		"INSERT INTO answers (answer, correct, question_id) " +
			"VALUES ('" + answer + "', '" + strconv.FormatBool(correct) + "', '" + strconv.FormatInt(question_id, 10) + "') ")
}

func DeleteAnswer(id string) error {
	return executeQuery("DELETE FROM answers WHERE answer_id = " + id)
}

func UpdateAnswer(answer entities.Answer) error {
	return executeQuery(
		"INSERT INTO answers (answer_id, answer, correct, question_id) " +
			"VALUES ('" + strconv.FormatInt(*answer.Answer_id, 10) + "', '" + *answer.Answer + "', '" + strconv.FormatBool(*answer.Correct) + "', '" + strconv.FormatInt(*answer.Question_id, 10) + "') " +
			"ON CONFLICT (answer_id) DO UPDATE SET answer = '" + *answer.Answer + "', correct = '" + strconv.FormatBool(*answer.Correct) + "', question_id = '" + strconv.FormatInt(*answer.Question_id, 10) + "'")
}

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
