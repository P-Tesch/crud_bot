CREATE TABLE botusers_items (
	"botuser_id" INTEGER NOT NULL,
	"item_id" INTEGER NOT NULL,

    CONSTRAINT fk_botuser_id FOREIGN KEY (botuser_id) REFERENCES botusers(botuser_id),
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(item_id)
);