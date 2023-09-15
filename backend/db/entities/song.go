package entities

type Song struct {
	Id     *int64
	Name   *string
	Url    *string
	Author *string
	Genre  *Genre
}

func NewSong(id *int64, name *string, url *string, author *string, genre *Genre) *Song {
	song := new(Song)
	song.Id = id
	song.Name = name
	song.Url = url
	song.Author = author
	song.Genre = genre
	return song
}
