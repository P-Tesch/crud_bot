package handlers

import (
	"crud_bot/db/entities"
	"crud_bot/db/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func RegisterItemHandler() {
	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")
			botuserId := urlQuery.Get("botuser_id")

			if id != "" {
				result = services.RetrieveItemById(id)

			} else if name != "" {
				result = services.RetrieveItemByName(name)

			} else if botuserId != "" {
				result = services.RetrieveItemByBotuserId(botuserId)

			} else {
				result = services.RetrieveAllItems()
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			item := new(entities.Item)
			json.Unmarshal(body, item)
			err = services.CreateItem(*item.Name, *item.Description)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "items/")[1]
			err := services.DeleteItem(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
