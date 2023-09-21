package main

import (
	"crud_bot/db/handlers"
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
	handlers.RegisterBotuserHandler()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
