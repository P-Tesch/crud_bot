CREATE TABLE IF NOT EXISTS botusers (
	"botuser_id" serial PRIMARY KEY,
    "discord_id" BIGINT UNIQUE,
    "currency" INTEGER,
	"score_id" INTEGER UNIQUE,

    CONSTRAINT fk_score_id FOREIGN KEY (score_id) REFERENCES scores(score_id)
);