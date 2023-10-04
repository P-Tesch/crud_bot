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

func RegisterSubtopicHandler() {
	http.HandleFunc("/subtopics/", func(w http.ResponseWriter, r *http.Request) {
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
			err = services.CreateSubtopic(*subtopic.Subtopic, *subtopic.Topic)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "subtopics/")[1]
			err := services.DeleteSubtopic(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "PUT":
			id := strings.Split(r.URL.Path, "subtopics/")[1]
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			subtopic := new(entities.Subtopic)
			err = json.Unmarshal(body, subtopic)
			if err != nil {
				fmt.Fprintf(w, "Error unmarshalling body: %v\n", err)
			}

			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error converting id: Error parsing from string to int64")
			}

			if *subtopic.Subtopic_id != idInt {
				fmt.Fprintf(w, "Error in path: Path id does not match object id")
			}

			err = services.UpdateSubtopic(*subtopic)
			if err == nil {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
