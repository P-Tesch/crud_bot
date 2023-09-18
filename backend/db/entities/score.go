package entities

type Score struct {
	Score_id        *int64
	Musicle_total   *int
	Musicle_win     *int
	Quiz_total      *int
	Quiz_win        *int
	Tictactoe_total *int
	Tictactoe_win   *int
	Chess_total     *int
	Chess_win       *int
}

func NewScore(id *int64, musicle_total *int, musicle_win *int, quiz_total *int, quiz_win *int, tictactoe_total *int, tictactoe_win *int, chess_total *int, chess_win *int) *Score {
	score := new(Score)
	score.Score_id = id
	score.Musicle_total = musicle_total
	score.Musicle_win = musicle_win
	score.Quiz_total = quiz_total
	score.Quiz_win = quiz_win
	score.Tictactoe_total = tictactoe_total
	score.Tictactoe_win = tictactoe_win
	score.Chess_total = chess_total
	score.Chess_win = chess_win
	return score
}
