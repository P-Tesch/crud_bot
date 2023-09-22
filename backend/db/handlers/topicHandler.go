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

func RegisterTopicHandler() {
	http.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			topic := urlQuery.Get("topic")

			if id != "" {
				result = services.RetrieveTopicById(id)

			} else if topic != "" {
				result = services.RetrieveTopicByTopic(topic)

			} else {
				result = services.RetrieveAllTopics()
			}

			fmt.Fprintf(w, string(result))
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			topic := new(entities.Topic)
			json.Unmarshal(body, topic)
			result, err := services.CreateTopic(*topic.Topic)

			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"topic_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
