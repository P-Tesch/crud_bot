package entities

type Item struct {
	Item_id     *int64
	Name        *string
	Description *string
}

func NewItem(id *int64, name *string, description *string) *Item {
	item := new(Item)
	item.Item_id = id
	item.Name = name
	item.Description = description
	return item
}
