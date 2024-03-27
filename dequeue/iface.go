package dequeue

import "sync"

type Dequeue[T any] interface {
	Tail() (val Iterator[T])
	Head() (val Iterator[T])
	Length() (res int64)
	SafePopFront() (val T)
	PopFront() (val T)
	SafePopBack() (val T)
	PopBack() (val T)
	SafePushFront(val T)
	PushFront(val T)
	SafePushBack(val T)
	PushBack(val T)
}

type Iterator[T any] interface {
	Next() Iterator[T]
	Prev() Iterator[T]
	Value() T
}

func New[T any]() Dequeue[T] {
	return &dequeue[T]{
		mutex: &sync.RWMutex{},
	}
}
