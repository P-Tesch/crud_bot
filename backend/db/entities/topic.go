package entities

type Topic struct {
	Topic_id *int64
	Topic    *string
}

func NewTopic(topic_id *int64, topic *string) *Topic {
	newTopic := new(Topic)
	newTopic.Topic_id = topic_id
	newTopic.Topic = topic
	return newTopic
}
