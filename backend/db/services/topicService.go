package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTopic(topic string) (int64, error) {
	return createGeneric(
		"INSERT INTO topics (topic) " +
			"VALUES ('" + topic + "') " +
			"RETURNING topic_id")
}

func DeleteTopic(id string) error {
	return deleteGeneric("DELETE FROM topics WHERE topic_id = " + id)
}

func retrieveTopíc(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Topic
	for results.Next() {
		var id int64
		var topic string
		results.Scan(&id, &topic)
		resultSet = append(resultSet, entities.NewTopic(&id, &topic))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}

func RetrieveAllTopics() []byte {
	return retrieveTopíc("SELECT * FROM topics")
}

func RetrieveTopicById(id string) []byte {
	return retrieveGenre("SELECT * FROM topics t WHERE t.topic_id = " + id)
}

func RetrieveTopicByTopic(topic string) []byte {
	return retrieveGenre("SELECT * FROM topics t WHERE t.topic iLike '" + topic + "'")
}
