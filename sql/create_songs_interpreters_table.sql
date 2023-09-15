CREATE TABLE songs_interpreters (
	"song_id" INTEGER NOT NULL,
    "interpreter_id" INTEGER NOT NULL,

    CONSTRAINT fk_song_id FOREIGN KEY (song_id) REFERENCES songs(song_id),
    CONSTRAINT fk_interpreter_id FOREIGN KEY (interpreter_id) REFERENCES interpreters(interpreter_id)
);