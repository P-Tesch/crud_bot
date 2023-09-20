package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateInterpreter(name string) int64 {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return 0
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to begin transaction: %v\n", err)
		return 0
	}

	results, err := tx.Exec(context.Background(), "INSERT INTO interpreters (name) VALUES ('"+name+"')")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
		return 0
	}

	err = tx.Commit(context.Background())
	return results.RowsAffected()
}

func retrieveInterpreter(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
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

func RetrieveAllInterpreters() []byte {
	return retrieveInterpreter("SELECT * FROM interpreters")
}

func RetrieveInterpreterById(id string) []byte {
	return retrieveInterpreter("SELECT * FROM interpreters i WHERE i.interpreter_id = " + id)
}

func RetrieveInterpreterByName(name string) []byte {
	return retrieveInterpreter("SELECT * FROM interpreters i WHERE i.name iLike '" + name + "'")
}
