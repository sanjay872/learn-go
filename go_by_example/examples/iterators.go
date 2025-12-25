package examples

import (
	"fmt"
	"iter"
)

func (lst *List[T]) push(t T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: t}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: t}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for i := lst.head; i != nil; i = i.next {
			if !yield(i.val) {
				return
			}
		}

	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func Iterators() {
	lst := List[int]{}

	lst.push(1)
	lst.push(2)
	lst.push(3)
	lst.push(4)
	lst.push(5)

	for i := range lst.All() {
		fmt.Println(i)
	}

	for i := range genFib() {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}
}
