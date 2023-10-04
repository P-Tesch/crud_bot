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

func CreateItem(name string, description string) error {
	return executeQuery(
		"INSERT INTO items (name, description) " +
			"VALUES ('" + name + "', '" + description + "') ")
}

func DeleteItem(id string) error {
	return executeQuery("DELETE FROM items WHERE item_id = " + id)
}

func UpdateItem(item entities.Item) error {
	return executeQuery(
		"INSERT INTO items (item_id, name, description) " +
			"VALUES ('" + strconv.FormatInt(*item.Item_id, 10) + "', '" + *item.Name + "', '" + *item.Description + "') " +
			"ON CONFLICT (item_id) DO UPDATE SET name = '" + *item.Name + "', description = '" + *item.Description + "'")
}

func retrieveItem(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Item
	for results.Next() {
		var id int64
		var name string
		var description string
		results.Scan(&id, &name, &description)
		resultSet = append(resultSet, entities.NewItem(&id, &name, &description))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}

func RetrieveAllItems() []byte {
	return retrieveItem("SELECT * FROM items")
}

func RetrieveItemById(id string) []byte {
	return retrieveItem(
		"SELECT * FROM items i " +
			"WHERE i.item_id = " + id)
}

func RetrieveItemByName(name string) []byte {
	return retrieveItem(
		"SELECT * FROM items i " +
			"WHERE i.name iLike '" + name + "'")
}

func RetrieveItemByBotuserId(botuserId string) []byte {
	return retrieveItem(
		"SELECT i.item_id, i.name, i.description FROM items i " +
			"JOIN botusers_items bi ON bi.item_id = i.item_id " +
			"WHERE bi.botuser_id = " + botuserId)
}
