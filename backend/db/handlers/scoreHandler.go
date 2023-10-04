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

func RegisterScoreHandler() {
	http.HandleFunc("/scores/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			botuserId := urlQuery.Get("botuser_id")

			if id != "" {
				result = services.RetrieveScoreById(id)

			} else if botuserId != "" {
				result = services.RetrieveScoreByBotuserId(botuserId)

			} else {
				result = services.RetrieveAllScores()
			}

			fmt.Fprintf(w, string(result))
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			score := new(entities.Score)
			json.Unmarshal(body, score)
			err = services.CreateScore(*score.Musicle_total, *score.Musicle_win, *score.Quiz_total, *score.Quiz_win, *score.Tictactoe_total, *score.Tictactoe_win, *score.Chess_total, *score.Chess_win)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "scores/")[1]
			err := services.DeleteScore(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "PUT":
			id := strings.Split(r.URL.Path, "scores/")[1]
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			score := new(entities.Score)
			err = json.Unmarshal(body, score)
			if err != nil {
				fmt.Fprintf(w, "Error unmarshalling body: %v\n", err)
			}

			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error converting id: Error parsing from string to int64")
			}

			if *score.Score_id != idInt {
				fmt.Fprintf(w, "Error in path: Path id does not match object id")
			}

			err = services.UpdateScore(*score)
			if err == nil {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
