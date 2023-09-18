package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RetrieveAllQuestions() []byte {
	connection, err := pgxpool.New(context.Background(), "postgres://username:password@localhost:5432/postgres-bot")
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "SELECT q.question_id, q.question, to_json(s), to_json(array_agg(a)) FROM questions q join subtopics s on s.subtopic_id = q.subtopic_id join answers a on a.question_id = q.question_id group by q.question_id, q.question, s.subtopic_id, s.subtopic ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Question
	for results.Next() {
		var id int64
		var question string
		var subtopic *entities.Subtopic
		var answers *[]entities.Answer
		var subtopicByte []byte
		var answersByte []byte

		results.Scan(&id, &question, &subtopic, &answers)
		json.Unmarshal(subtopicByte, subtopic)
		json.Unmarshal(answersByte, answers)
		resultSet = append(resultSet, entities.NewQuestion(&id, &question, subtopic, answers))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}
