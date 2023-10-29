package migration

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate() error {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL")+"&user="+os.Getenv("ADM_USERNAME")+"&password="+os.Getenv("ADM_PASSWORD"))
	defer connection.Close()
	if err != nil {
		return err
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		return err
	}

	scripts := getScripts()

	for _, statement := range scripts {
		results, err := tx.Query(context.Background(), statement)
		if err != nil {
			return err
		}

		results.Close()
		err = results.Err()
		if err != nil {

		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func getScripts() []string {
	files, err := os.ReadDir("./assets/sql")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	scripts := make([]string, len(files))
	for i, file := range files {
		statement, err := os.ReadFile("./assets/sql/" + file.Name())
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		scripts[i] = string(statement)
	}

	return scripts
}
