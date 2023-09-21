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

func CreateQuestion(question string, subtopic entities.Subtopic, answers []entities.Answer) int64 {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return 0
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to begin transaction: %v\n", err)
		return 0
	}

	resultsQuestions, err := tx.Query(context.Background(),
		"INSERT INTO questions (question, subtopic_id) "+
			"VALUES ('"+question+"', '"+strconv.FormatInt(*subtopic.Subtopic_id, 10)+"') "+
			"RETURNING question_id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
		return 0
	}

	resultsQuestions.Next()
	var id int64
	resultsQuestions.Scan(&id)
	resultsQuestions.Close()

	for i := range answers {
		resultsAnswers, err := tx.Query(context.Background(), "INSERT INTO answers (answer, correct, question_id) VALUES ('"+*answers[i].Answer+"', "+strconv.FormatBool(*answers[i].Correct)+", '"+strconv.FormatInt(id, 10)+"')")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
			return 0
		}
		resultsAnswers.Close()
	}

	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
		return 0
	}
	return id
}

func retrieveQuestion(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Question
	for results.Next() {
		var id int64
		var question string
		var subtopic *entities.Subtopic
		var answers *[]entities.Answer
		var topic *entities.Topic
		var subtopicByte []byte
		var answersByte []byte
		var topicByte []byte

		results.Scan(&id, &question, &subtopic, &answers, &topic)
		json.Unmarshal(subtopicByte, subtopic)
		json.Unmarshal(answersByte, answers)
		json.Unmarshal(topicByte, topic)
		subtopic.Topic = topic
		resultSet = append(resultSet, entities.NewQuestion(&id, &question, subtopic, answers))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}

func RetrieveAllQuestions() []byte {
	return retrieveQuestion(
		"SELECT q.question_id, q.question, TO_JSON(s), TO_JSON(ARRAY_AGG(a)), TO_JSON(t) FROM questions q " +
			"JOIN subtopics s ON s.subtopic_id = q.subtopic_id " +
			"JOIN topics t ON t.topic_id = s.topic_id " +
			"JOIN answers a ON a.question_id = q.question_id " +
			"GROUP BY q.question_id, q.question, s.subtopic_id, s.subtopic, t.topic_id, t.topic")
}

func RetrieveQuestionById(id string) []byte {
	return retrieveQuestion(
		"SELECT q.question_id, q.question, TO_JSON(s), TO_JSON(ARRAY_AGG(a)), TO_JSON(t) FROM questions q " +
			"JOIN subtopics s ON s.subtopic_id = q.subtopic_id " +
			"JOIN topics t ON t.topic_id = s.topic_id " +
			"JOIN answers a ON a.question_id = q.question_id " +
			"WHERE q.question_id = " + id + " " +
			"GROUP BY q.question_id, q.question, s.subtopic_id, s.subtopic, t.topic_id, t.topic")
}

func RetrieveQuestionBySubtopicId(subtopicId string) []byte {
	return retrieveQuestion(
		"SELECT q.question_id, q.question, TO_JSON(s), TO_JSON(ARRAY_AGG(a)), TO_JSON(t) FROM questions q " +
			"JOIN subtopics s ON s.subtopic_id = q.subtopic_id " +
			"JOIN topics t ON t.topic_id = s.topic_id " +
			"JOIN answers a ON a.question_id = q.question_id " +
			"WHERE s.subtopic_id = " + subtopicId + " " +
			"GROUP BY q.question_id, q.question, s.subtopic_id, s.subtopic, t.topic_id, t.topic")
}

func RetrieveQuestionBySubtopicSubtopic(subtopicSubtopic string) []byte {
	return retrieveQuestion(
		"SELECT q.question_id, q.question, TO_JSON(s), TO_JSON(ARRAY_AGG(a)), TO_JSON(t) FROM questions q " +
			"JOIN subtopics s ON s.subtopic_id = q.subtopic_id " +
			"JOIN topics t ON t.topic_id = s.topic_id " +
			"JOIN answers a ON a.question_id = q.question_id " +
			"WHERE s.subtopic iLike '" + subtopicSubtopic + "' " +
			"GROUP BY q.question_id, q.question, s.subtopic_id, s.subtopic, t.topic_id, t.topic")
}

func RetrieveQuestionByTopicId(topicId string) []byte {
	return retrieveQuestion(
		"SELECT q.question_id, q.question, TO_JSON(s), TO_JSON(ARRAY_AGG(a)), TO_JSON(t) FROM questions q " +
			"JOIN subtopics s ON s.subtopic_id = q.subtopic_id " +
			"JOIN topics t ON t.topic_id = s.topic_id " +
			"JOIN answers a ON a.question_id = q.question_id " +
			"WHERE t.topic_id = " + topicId + " " +
			"GROUP BY q.question_id, q.question, s.subtopic_id, s.subtopic, t.topic_id, t.topic")
}

func RetrieveQuestionByTopicTopic(topicTopic string) []byte {
	return retrieveQuestion(
		"SELECT q.question_id, q.question, TO_JSON(s), TO_JSON(ARRAY_AGG(a)), TO_JSON(t) FROM questions q " +
			"JOIN subtopics s ON s.subtopic_id = q.subtopic_id " +
			"JOIN topics t ON t.topic_id = s.topic_id " +
			"JOIN answers a ON a.question_id = q.question_id " +
			"WHERE t.topic iLike '" + topicTopic + "' " +
			"GROUP BY q.question_id, q.question, s.subtopic_id, s.subtopic, t.topic_id, t.topic")
}
