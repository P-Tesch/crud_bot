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

func RegisterSubtopicHandler() {
	http.HandleFunc("/subtopics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			subtopic := urlQuery.Get("subtopic")
			topicTopic := urlQuery.Get("topic_topic")
			topicId := urlQuery.Get("topic_id")

			if id != "" {
				result = services.RetrieveSubtopicById(id)

			} else if subtopic != "" {
				result = services.RetrieveSubtopicBySubtopic(subtopic)

			} else if topicTopic != "" {
				result = services.RetrieveSubtopicByTopicTopic(topicTopic)

			} else if topicId != "" {
				result = services.RetrieveSubtopicByTopicId(topicId)

			} else {
				result = services.RetrieveAllSubtopics()
			}

			fmt.Fprintf(w, string(result))
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			subtopic := new(entities.Subtopic)
			json.Unmarshal(body, subtopic)
			result := services.CreateSubtopic(*subtopic.Subtopic, *subtopic.Topic)

			w.WriteHeader(201)
			fmt.Fprintf(w, "{\"subtopic_id\": "+strconv.FormatInt(result, 10)+"}")
		}
	})
}
