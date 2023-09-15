CREATE TABLE questions (
	"question_id" serial PRIMARY KEY,
	"question" VARCHAR ( 512 ) UNIQUE NOT NULL,
    "subtopic_id" INTEGER NOT NULL,

    CONSTRAINT fk_subtopic_id FOREIGN KEY (subtopic_id) REFERENCES subtopics(subtopic_id)
);