package entities

type Song struct {
	Song_id      *int64
	Name         *string
	Url          *string
	Interpreters *[]Interpreter
	Genre        *Genre
}

func NewSong(id *int64, name *string, url *string, interpreters *[]Interpreter, genre *Genre) *Song {
	song := new(Song)
	song.Song_id = id
	song.Name = name
	song.Url = url
	song.Interpreters = interpreters
	song.Genre = genre
	return song
}
