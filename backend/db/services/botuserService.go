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

func CreateBotuser(discord_id int64, currency int, score entities.Score, items []entities.Item) error {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		return err
	}

	tx, err := connection.Begin(context.Background())
	defer tx.Rollback(context.Background())
	if err != nil {
		return err
	}

	resultsBotusers, err := tx.Query(context.Background(),
		"INSERT INTO botusers (discord_id, currency, score_id) "+
			"VALUES ('"+strconv.FormatInt(discord_id, 10)+"', '"+strconv.Itoa(currency)+"', '"+strconv.FormatInt(*score.Score_id, 10)+"') "+
			"RETURNING botuser_id")
	if err != nil {
		return err
	}

	var id int64
	if resultsBotusers.Next() {
		resultsBotusers.Scan(&id)
		resultsBotusers.Close()
	}
	err = resultsBotusers.Err()
	if err != nil {
		return err
	}

	for i := range items {
		resultsJoin, err := tx.Query(context.Background(),
			"INSERT INTO botusers_items (botuser_id, item_id) "+
				"VALUES ('"+strconv.FormatInt(id, 10)+"', '"+strconv.FormatInt(*items[i].Item_id, 10)+"')")
		if err != nil {
			return err
		}
		resultsJoin.Close()
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func DeleteBotuser(id string) error {
	return executeQuery("DELETE FROM botusers WHERE botuser_id = " + id)
}

func retrieveBotuser(query string) []byte {
	connection, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	defer connection.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	results, err := connection.Query(context.Background(), query)
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

func RetrieveAllBotusers() []byte {
	return retrieveBotuser(
		"SELECT b.botuser_id, b.discord_id, b.currency, TO_JSON(s), TO_JSON(ARRAY_AGG(i)) FROM botusers b " +
			"JOIN scores s ON s.score_id = b.score_id " +
			"JOIN botusers_items bi ON b.botuser_id = bi.botuser_id " +
			"JOIN items i ON i.item_id = bi.item_id " +
			"GROUP BY b.botuser_id, b.discord_id, b.currency, s.score_id")
}

func RetrieveBotuserById(id string) []byte {
	return retrieveBotuser(
		"SELECT b.botuser_id, b.discord_id, b.currency, TO_JSON(s), TO_JSON(ARRAY_AGG(i)) FROM botusers b " +
			"JOIN scores s ON s.score_id = b.score_id " +
			"JOIN botusers_items bi ON b.botuser_id = bi.botuser_id " +
			"JOIN items i ON i.item_id = bi.item_id " +
			"WHERE b.botuser_id = " + id + " " +
			"GROUP BY b.botuser_id, b.discord_id, b.currency, s.score_id")
}

func RetrieveBotuserByDiscordId(discordId string) []byte {
	return retrieveBotuser(
		"SELECT b.botuser_id, b.discord_id, b.currency, TO_JSON(s), TO_JSON(ARRAY_AGG(i)) FROM botusers b " +
			"JOIN scores s ON s.score_id = b.score_id " +
			"JOIN botusers_items bi ON b.botuser_id = bi.botuser_id " +
			"JOIN items i ON i.item_id = bi.item_id " +
			"WHERE b.discord_id = " + discordId + " " +
			"GROUP BY b.botuser_id, b.discord_id, b.currency, s.score_id")
}

func RetrieveBotuserByScoreId(scoreId string) []byte {
	return retrieveBotuser(
		"SELECT b.botuser_id, b.discord_id, b.currency, TO_JSON(s), TO_JSON(ARRAY_AGG(i)) FROM botusers b " +
			"JOIN scores s ON s.score_id = b.score_id " +
			"JOIN botusers_items bi ON b.botuser_id = bi.botuser_id " +
			"JOIN items i ON i.item_id = bi.item_id " +
			"WHERE b.score_id = " + scoreId + " " +
			"GROUP BY b.botuser_id, b.discord_id, b.currency, s.score_id")
}
