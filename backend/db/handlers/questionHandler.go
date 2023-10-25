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
		username, password, _ := r.BasicAuth()
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
				result = services.RetrieveQuestionById(id, username, password)

			} else if subtopicId != "" {
				result = services.RetrieveQuestionBySubtopicId(subtopicId, username, password)

			} else if subtopicSubtopic != "" {
				result = services.RetrieveQuestionBySubtopicSubtopic(subtopicSubtopic, username, password)

			} else if topicId != "" {
				result = services.RetrieveQuestionByTopicId(topicId, username, password)

			} else if topicTopic != "" {
				result = services.RetrieveQuestionByTopicTopic(topicTopic, username, password)

			} else {
				result = services.RetrieveAllQuestions(username, password)
			}

			fmt.Fprintf(w, string(result))
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			question := new(entities.Question)
			json.Unmarshal(body, question)
			err = services.CreateQuestion(*question.Question, *question.Subtopic, *question.Answers, username, password)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "questions/")[1]
			err := services.DeleteQuestion(id, username, password)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "PUT":
			id := strings.Split(r.URL.Path, "questions/")[1]
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			question := new(entities.Question)
			err = json.Unmarshal(body, question)
			if err != nil {
				fmt.Fprintf(w, "Error unmarshalling body: %v\n", err)
			}

			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error converting id: Error parsing from string to int64")
			}

			if *question.Question_id != idInt {
				fmt.Fprintf(w, "Error in path: Path id does not match object id")
			}

			err = services.UpdateQuestion(*question, username, password)
			if err == nil {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
