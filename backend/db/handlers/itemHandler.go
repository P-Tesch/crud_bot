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

func RegisterItemHandler() {
	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := r.BasicAuth()
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")
			botuserId := urlQuery.Get("botuser_id")

			if id != "" {
				result = services.RetrieveItemById(id, username, password)

			} else if name != "" {
				result = services.RetrieveItemByName(name, username, password)

			} else if botuserId != "" {
				result = services.RetrieveItemByBotuserId(botuserId, username, password)

			} else {
				result = services.RetrieveAllItems(username, password)
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			item := new(entities.Item)
			json.Unmarshal(body, item)
			err = services.CreateItem(*item.Name, *item.Description, username, password)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "items/")[1]
			err := services.DeleteItem(id, username, password)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "PUT":
			id := strings.Split(r.URL.Path, "items/")[1]
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			item := new(entities.Item)
			err = json.Unmarshal(body, item)
			if err != nil {
				fmt.Fprintf(w, "Error unmarshalling body: %v\n", err)
			}

			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error converting id: Error parsing from string to int64")
			}

			if *item.Item_id != idInt {
				fmt.Fprintf(w, "Error in path: Path id does not match object id")
			}

			err = services.UpdateItem(*item, username, password)
			if err == nil {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})
}
