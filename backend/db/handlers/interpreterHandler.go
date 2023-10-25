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
		username, password, _ := r.BasicAuth()
		switch r.Method {
		case "GET":
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")

			if id != "" {
				result = services.RetrieveInterpreterById(id, username, password)

			} else if name != "" {
				result = services.RetrieveInterpreterByName(name, username, password)

			} else {
				result = services.RetrieveAllInterpreters(username, password)
			}

			fmt.Fprintf(w, string(result))

		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			interpreter := new(entities.Interpreter)
			json.Unmarshal(body, interpreter)
			err = services.CreateInterpreter(*interpreter.Name, username, password)

			if err == nil {
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "DELETE":
			id := strings.Split(r.URL.Path, "interpreters/")[1]
			err := services.DeleteInterpreter(id, username, password)

			if err == nil {
				w.WriteHeader(204)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		case "PUT":
			id := strings.Split(r.URL.Path, "interpreters/")[1]
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Error reading body: %v\n", err)
			}

			interpreter := new(entities.Interpreter)
			err = json.Unmarshal(body, interpreter)
			if err != nil {
				fmt.Fprintf(w, "Error unmarshalling body: %v\n", err)
			}

			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error converting id: Error parsing from string to int64")
			}

			if *interpreter.Interpreter_id != idInt {
				fmt.Fprintf(w, "Error in path: Path id does not match object id")
			}

			err = services.UpdateInterpreter(*interpreter, username, password)
			if err == nil {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
			}
		}
	})

}
