CREATE TABLE IF NOT EXISTS genres (
	"genre_id" serial PRIMARY KEY,
	"name" VARCHAR ( 50 ) UNIQUE NOT NULL
);