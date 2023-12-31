package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateInterpreter(name string, username string, password string) error {
	return executeQuery(
		"INSERT INTO interpreters (name) "+
			"VALUES ('"+name+"') ", username, password)
}

func DeleteInterpreter(id string, username string, password string) error {
	return executeQuery("DELETE FROM interpreters WHERE interpreter_id = "+id, username, password)
}

func UpdateInterpreter(interpreter entities.Interpreter, username string, password string) error {
	return executeQuery(
		"INSERT INTO interpreters (interpreter_id, name) "+
			"VALUES ('"+strconv.FormatInt(*interpreter.Interpreter_id, 10)+"', '"+*interpreter.Name+"') "+
			"ON CONFLICT (interpreter_id) DO UPDATE SET name = '"+*interpreter.Name+"'", username, password)
}

func retrieveInterpreter(query string, username string, password string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"&user="+username+"&password="+password)
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Interpreter
	for results.Next() {
		var id int64
		var name string
		results.Scan(&id, &name)
		resultSet = append(resultSet, entities.NewInterpreter(&id, &name))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}

func RetrieveAllInterpreters(username string, password string) []byte {
	return retrieveInterpreter("SELECT * FROM interpreters", username, password)
}

func RetrieveInterpreterById(id string, username string, password string) []byte {
	return retrieveInterpreter("SELECT * FROM interpreters i WHERE i.interpreter_id = "+id, username, password)
}

func RetrieveInterpreterByName(name string, username string, password string) []byte {
	return retrieveInterpreter("SELECT * FROM interpreters i WHERE i.name iLike '"+name+"'", username, password)
}
