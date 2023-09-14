package entities

type genre struct {
	name *string
}

func newGenre(name *string) *genre {
	genre := new(genre)
	genre.name = name
	return genre
}
