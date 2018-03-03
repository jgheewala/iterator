package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	arr2 := make([]interface{}, 0)
	arr := []interface{}{1, 2, 3}
	it := NewIterator()
	it.Init(arr)
	for it.HasNext() {
		fmt.Println(it.Next())
	}
	arr2 = append(arr2, it)
	arr1 := []interface{}{4, 5}
	it2 := NewIterator()
	it2.Init(arr1)
	for it2.HasNext() {
		fmt.Println(it2.Next())
	}
	arr2 = append(arr2, it2)
	fmt.Println("done with two iterators")

	itC := NewIterator()
	itC.Init(arr2)
	/*
	   for itC.HasNext() {
	       fmt.Println(itC.Next())
	   }
	*/

	fl := NewFlatteningIterator(itC)
	for fl.HasNext() {
		fl.Next()
	}

	fmt.Println("done")
}
