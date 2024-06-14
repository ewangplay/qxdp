package iterator

type ConcreteAggregate struct {
	items []*Item
}

func NewConcreteAggregate(capacity int) *ConcreteAggregate {
	return &ConcreteAggregate{
		items: make([]*Item, 0, capacity),
	}
}

func (b *ConcreteAggregate) Add(item *Item) {
	b.items = append(b.items, item)
}

func (b *ConcreteAggregate) Get(index int) *Item {
	return b.items[index]
}

func (b *ConcreteAggregate) GetLength() int {
	return len(b.items)
}

func (b *ConcreteAggregate) Iterator() Iterator {
	return NewConcreteIterator(b)
}
