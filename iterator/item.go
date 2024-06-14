package iterator

type Item struct {
	name string
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (b *Item) GetName() string {
	return b.name
}
