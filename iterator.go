package iterator

import (
	"fmt"
)

/*
Part 1: Here's the interface for an iterator. Implement it.

class Iterator {
    // candidate should insert constructor here
    boolean hasNext();
    Object next();
}

it = Iterator([1, 2, 3])
while it.hasNext():
    print it.next()

# 1
# 2
# 3
*/
type Iterator struct {
	data       []interface{}
	currentIdx uint64
}

func (i *Iterator) reset() {
	i.currentIdx = 0
}

func (i *Iterator) Next() interface{} {
	value := i.data[i.currentIdx]
	i.currentIdx++
	return value
}

func (i *Iterator) HasNext() bool {
	if i.currentIdx < uint64(len(i.data)) {
		return true
	}
	defer i.reset()
	return false
}

func NewIterator() *Iterator {
	i := Iterator{}
	return &i
}

func (i *Iterator) Init(arr []interface{}) {
	i.data = arr
	i.currentIdx = uint64(0)
}

type FlatteningIterator struct {
	it         *Iterator
	currentIdx uint64
}

func NewFlatteningIterator(it *Iterator) *FlatteningIterator {
	fl := FlatteningIterator{
		it:         it,
		currentIdx: uint64(0),
	}
	return &fl
}

func (f *FlatteningIterator) HasNext() bool {
	return f.it.HasNext()
}

func (f *FlatteningIterator) Next() {
	// next is func (i *Iterator) Next() interface{} {
	// so now its a function and we need to print the values in this function
	// we can do this by calling this function...
	next := f.it.Next()
	it := next.(*Iterator)
	for it.HasNext() {
		fmt.Println(it.Next())
	}

	return
}
