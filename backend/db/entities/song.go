package entities

type song struct {
	name   *string
	url    *string
	author *string
	genre  *genre
}

func newSong(name *string, url *string, author *string, genre *genre) *song {
	song := new(song)
	song.name = name
	song.url = url
	song.author = author
	song.genre = genre
	return song
}
