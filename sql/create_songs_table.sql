CREATE TABLE songs (
	"song_id" serial PRIMARY KEY,
	"name" VARCHAR ( 50 ) NOT NULL,
	"url" VARCHAR ( 255 ) UNIQUE NOT NULL,
	"genre_id" INTEGER NOT NULL,

    CONSTRAINT fk_genre_id FOREIGN KEY (genre_id) REFERENCES genres(genre_id)
);