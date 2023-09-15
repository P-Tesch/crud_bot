CREATE TABLE answers (
	"answer_id" serial PRIMARY KEY,
	"answer" VARCHAR ( 255 ) NOT NULL,
    "correct" BOOLEAN NOT NULL,
    "question_id" INTEGER NOT NULL,

    CONSTRAINT fk_question_id FOREIGN KEY (question_id) REFERENCES questions(question_id)
);