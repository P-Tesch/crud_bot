package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RetrieveAllSubtopics() []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "SELECT s.subtopic_id, s.subtopic, to_json(t) FROM subtopics s JOIN topics t ON s.topic_id = t.topic_id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Subtopic
	for results.Next() {
		var id int64
		var subtopic string
		var topicByte []byte
		var topic *entities.Topic
		results.Scan(&id, &subtopic, &topic)
		json.Unmarshal(topicByte, topic)
		resultSet = append(resultSet, entities.NewSubtopic(&id, &subtopic, topic))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}
