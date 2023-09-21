package handlers

import (
	"crud_bot/db/entities"
	"crud_bot/db/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func RegisterAnswerHandler() {
	http.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
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
			result := services.CreateAnswer(*answer.Answer, *answer.Correct, *answer.Question_id)

			w.WriteHeader(201)
			fmt.Fprintf(w, "{\"answer_id\": "+strconv.FormatInt(result, 10)+"}")

		}
	})
}