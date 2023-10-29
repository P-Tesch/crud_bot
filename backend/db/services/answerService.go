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

func CreateAnswer(answer string, correct bool, question_id int64, username string, password string) error {
	return executeQuery(
		"INSERT INTO answers (answer, correct, question_id) "+
			"VALUES ('"+answer+"', '"+strconv.FormatBool(correct)+"', '"+strconv.FormatInt(question_id, 10)+"') ", username, password)
}

func DeleteAnswer(id string, username string, password string) error {
	return executeQuery("DELETE FROM answers WHERE answer_id = "+id, username, password)
}

func UpdateAnswer(answer entities.Answer, username string, password string) error {
	return executeQuery(
		"INSERT INTO answers (answer_id, answer, correct, question_id) "+
			"VALUES ('"+strconv.FormatInt(*answer.Answer_id, 10)+"', '"+*answer.Answer+"', '"+strconv.FormatBool(*answer.Correct)+"', '"+strconv.FormatInt(*answer.Question_id, 10)+"') "+
			"ON CONFLICT (answer_id) DO UPDATE SET answer = '"+*answer.Answer+"', correct = '"+strconv.FormatBool(*answer.Correct)+"', question_id = '"+strconv.FormatInt(*answer.Question_id, 10)+"'", username, password)
}

func retrieveAnswer(query string, username string, password string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"&user="+username+"&password="+password)
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

func RetrieveAllAnswers(username string, password string) []byte {
	return retrieveAnswer("SELECT * FROM answers", username, password)
}

func RetrieveAnswerById(id string, username string, password string) []byte {
	return retrieveAnswer(
		"SELECT * FROM answers a "+
			"WHERE a.answer_id = "+id, username, password)
}

func RetrieveAnswerByQuestionId(questionId string, username string, password string) []byte {
	return retrieveAnswer(
		"SELECT * FROM answers a "+
			"WHERE a.question_id = "+questionId, username, password)
}
