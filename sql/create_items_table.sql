CREATE TABLE items (
	"item_id" serial PRIMARY KEY,
	"name" VARCHAR ( 50 ) UNIQUE NOT NULL,
    "description" VARCHAR ( 255 ) NOT NULL
);