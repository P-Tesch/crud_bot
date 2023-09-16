package entities

type Answer struct {
	Answer_id   *int64
	Answer      *string
	Correct     *bool
	Question_id *int64
}

func NewAnswer(answer_id *int64, answer *string, correct *bool, question_id *int64) *Answer {
	newAnswer := new(Answer)
	newAnswer.Answer_id = answer_id
	newAnswer.Answer = answer
	newAnswer.Correct = correct
	newAnswer.Question_id = question_id
	return newAnswer
}
