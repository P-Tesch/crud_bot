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

func RegisterAnswerHandler() {
	http.HandleFunc("/answers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			questionId := urlQuery.Get("question_id")

			if id != "" {
				result = services.RetrieveAnswerById(id)

			} else if questionId != "" {
				result = services.RetrieveAnswerByQuestionId(questionId)

			} else {
				result = services.RetrieveAllAnswers()
			}

			fmt.Fprintf(w, string(result))
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			answer := new(entities.Answer)
			json.Unmarshal(body, answer)
			result, err := services.CreateAnswer(*answer.Answer, *answer.Correct, *answer.Question_id)
			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"answer_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "answers/")[1]
			err := services.DeleteAnswer(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
