package entities

type Subtopic struct {
	Subtopic_id *int64
	Subtopic    *string
	Topic       *Topic
}

func NewSubtopic(subtopic_id *int64, subtopic *string, topic *Topic) *Subtopic {
	newSubtopic := new(Subtopic)
	newSubtopic.Subtopic_id = subtopic_id
	newSubtopic.Subtopic = subtopic
	newSubtopic.Topic = topic
	return newSubtopic
}
