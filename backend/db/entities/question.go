package entities

type Question struct {
	Question_id *int64
	Question    *string
	Subtopic    *Subtopic
	Answers     *[]Answer
}

func NewQuestion(question_id *int64, question *string, subtopic *Subtopic, answers *[]Answer) *Question {
	newQuestion := new(Question)
	newQuestion.Question_id = question_id
	newQuestion.Question = question
	newQuestion.Subtopic = subtopic
	newQuestion.Answers = answers
	return newQuestion
}
