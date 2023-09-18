package entities

type Botuser struct {
	Botuser_id *int64
	Discord_id *int64
	Currency   *int
	Score      *Score
	Items      *[]Item
}

func NewBotuser(id *int64, discord_id *int64, currency *int, score *Score, items *[]Item) *Botuser {
	botuser := new(Botuser)
	botuser.Botuser_id = id
	botuser.Discord_id = discord_id
	botuser.Currency = currency
	botuser.Score = score
	botuser.Items = items
	return botuser
}
