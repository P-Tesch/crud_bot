package entities

type song struct {
	id     *int64
	name   *string
	url    *string
	author *string
	genre  *Genre
}

func NewSong(id *int64, name *string, url *string, author *string, genre *Genre) *song {
	song := new(song)
	song.id = id
	song.name = name
	song.url = url
	song.author = author
	song.genre = genre
	return song
}
