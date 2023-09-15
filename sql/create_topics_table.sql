CREATE TABLE topics (
	"topic_id" serial PRIMARY KEY,
	"topic" VARCHAR ( 50 ) UNIQUE NOT NULL
);