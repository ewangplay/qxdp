package iterator

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
}

func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: aggregate,
		index:     0,
	}
}

func (b *ConcreteIterator) HasNext() bool {
	return b.index < b.aggregate.GetLength()
}

func (b *ConcreteIterator) Next() interface{} {
	item := b.aggregate.Get(b.index)
	b.index++
	return item
}

func (b *ConcreteIterator) HasPrevious() bool {
	return b.index > 0
}

func (b *ConcreteIterator) Previous() interface{} {
	item := b.aggregate.Get(b.index - 1)
	b.index--
	return item
}
