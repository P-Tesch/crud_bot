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

func RegisterItemHandler() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
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
			result := services.CreateItem(*item.Name, *item.Description)

			w.WriteHeader(201)
			fmt.Fprintf(w, "{\"item_id\": "+strconv.FormatInt(result, 10)+"}")
		}
	})
}
