CREATE TABLE subtopics (
	"subtopic_id" serial PRIMARY KEY,
	"subtopic" VARCHAR ( 50 ) UNIQUE NOT NULL,
    "topic_id" INTEGER NOT NULL,

    CONSTRAINT fk_topic_id FOREIGN KEY (topic_id) REFERENCES topics(topic_id)
);