CREATE TABLE IF NOT EXISTS interpreters (
	"interpreter_id" serial PRIMARY KEY,
	"name" VARCHAR ( 50 ) UNIQUE NOT NULL
);