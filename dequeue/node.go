package dequeue

type node[T any] struct {
	next  *node[T]
	prev  *node[T]
	value T
}

func (n node[T]) Next() Iterator[T] {
	return n.next
}

func (n node[T]) Prev() Iterator[T] {
	return n.prev
}

func (n node[T]) Value() T {
	return n.value
}
