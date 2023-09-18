package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"crud_bot/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RetrieveAllBotusers() []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), "select b.botuser_id, b.discord_id, b.currency, to_json(s), to_json(array_agg(i)) from botusers b join scores s on s.score_id = b.score_id join botusers_items bi on b.botuser_id = bi.botuser_id join items i on i.item_id = bi.item_id group by b.botuser_id, b.discord_id, b.currency, s.score_id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable execute query: %v\n", err)
	}

	var resultSet []*entities.Botuser
	for results.Next() {
		var id int64
		var discord_id int64
		var currency int
		var score *entities.Score
		var items *[]entities.Item
		var scoreByte []byte
		var itemsByte []byte

		results.Scan(&id, &discord_id, &currency, &score, &items)
		json.Unmarshal(scoreByte, score)
		json.Unmarshal(itemsByte, items)
		resultSet = append(resultSet, entities.NewBotuser(&id, &discord_id, &currency, score, items))
	}

	jsonResult, err := json.Marshal(resultSet)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable parse JSON: %v\n", err)
	}
	return jsonResult
}
