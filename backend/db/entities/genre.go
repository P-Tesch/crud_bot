package entities

type Genre struct {
	Genre_id *int64
	Name     *string
}

func NewGenre(genre_id *int64, name *string) *Genre {
	genre := new(Genre)
	genre.Name = name
	genre.Genre_id = genre_id
	return genre
}
