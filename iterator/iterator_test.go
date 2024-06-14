package iterator_test

import (
	"testing"

	"github.com/ewangplay/qxdp/iterator"
)

func TestIterator(t *testing.T) {
	aggregate := iterator.NewConcreteAggregate(10)
	aggregate.Add(iterator.NewItem("四世同堂"))
	aggregate.Add(iterator.NewItem("骆驼祥子"))
	aggregate.Add(iterator.NewItem("茶馆"))
	aggregate.Add(iterator.NewItem("正红旗下"))
	aggregate.Add(iterator.NewItem("离婚"))
	aggregate.Add(iterator.NewItem("不是问题的问题"))
	aggregate.Add(iterator.NewItem("二马"))

	it := aggregate.Iterator()
	for it.HasNext() {
		item := it.Next().(*iterator.Item)
		t.Log(item.GetName())
	}

	for it.HasPrevious() {
		item := it.Previous().(*iterator.Item)
		t.Log(item.GetName())
	}
}
