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

func RegisterBotuserHandler() {
	http.HandleFunc("/botusers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
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

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			botuser := new(entities.Botuser)
			json.Unmarshal(body, botuser)
			result, err := services.CreateBotuser(*botuser.Discord_id, *botuser.Currency, *botuser.Score, *botuser.Items)

			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"botuser_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "botusers/")[1]
			err := services.DeleteBotuser(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
