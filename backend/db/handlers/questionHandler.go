package handlers

import (
	"crud_bot/db/entities"
	"crud_bot/db/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func RegisterQuestionHandler() {
	http.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			subtopicId := urlQuery.Get("subtopic_id")
			subtopicSubtopic := urlQuery.Get("subtopic_subtopic")
			topicTopic := urlQuery.Get("topic_topic")
			topicId := urlQuery.Get("topic_id")

			if id != "" {
				result = services.RetrieveQuestionById(id)

			} else if subtopicId != "" {
				result = services.RetrieveQuestionBySubtopicId(subtopicId)

			} else if subtopicSubtopic != "" {
				result = services.RetrieveQuestionBySubtopicSubtopic(subtopicSubtopic)

			} else if topicId != "" {
				result = services.RetrieveQuestionByTopicId(topicId)

			} else if topicTopic != "" {
				result = services.RetrieveQuestionByTopicTopic(topicTopic)

			} else {
				result = services.RetrieveAllQuestions()
			}

			fmt.Fprintf(w, string(result))
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			question := new(entities.Question)
			json.Unmarshal(body, question)
			result, err := services.CreateQuestion(*question.Question, *question.Subtopic, *question.Answers)

			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"question_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "questions/")[1]
			err := services.DeleteQuestion(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
