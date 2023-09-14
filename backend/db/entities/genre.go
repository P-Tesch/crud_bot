package entities

type Genre struct {
	Id   *int64
	Name *string
}

func NewGenre(id *int64, name *string) *Genre {
	genre := new(Genre)
	genre.Name = name
	genre.Id = id
	return genre
}
