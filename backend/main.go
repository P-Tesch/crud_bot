package main

import (
	"crud_bot/db/handlers"
	"crud_bot/db/services"
	"fmt"
	"net/http"
)

func main() {
	handlers.RegisterGenreHandler()
	handlers.RegisterInterpreterHandler()
	handlers.RegisterSongHandler()
	handlers.RegisterTopicHandler()
	handlers.RegisterSubtopicHandler()
	handlers.RegisterQuestionHandler()
	handlers.RegisterAnswerHandler()
	handlers.RegisterScoreHandler()
	handlers.RegisterItemHandler()

	http.HandleFunc("/botusers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			discordId := urlQuery.Get("discord_id")
			scoreId := urlQuery.Get("score_id")

			if id != "" {
				result = services.RetrieveBotuserById(id)

			} else if discordId != "" {
				result = services.RetrieveBotuserByDiscordId(discordId)

			} else if scoreId != "" {
				result = services.RetrieveBotuserByScoreId(scoreId)

			} else {
				result = services.RetrieveAllBotusers()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
