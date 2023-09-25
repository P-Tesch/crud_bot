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

func RegisterInterpreterHandler() {
	http.HandleFunc("/interpreters/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")

			if id != "" {
				result = services.RetrieveInterpreterById(id)

			} else if name != "" {
				result = services.RetrieveInterpreterByName(name)

			} else {
				result = services.RetrieveAllInterpreters()
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			interpreter := new(entities.Interpreter)
			json.Unmarshal(body, interpreter)
			result, err := services.CreateInterpreter(*interpreter.Name)

			if err == nil {
				w.WriteHeader(201)
				fmt.Fprintf(w, "{\"interpreter_id\": "+strconv.FormatInt(result, 10)+"}")
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "interpreters/")[1]
			err := services.DeleteInterpreter(id)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})

}
