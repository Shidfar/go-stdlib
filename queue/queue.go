package queue

const popOnEmptyQueueMessage = "popping on empty queue"

type Queue[V any] interface {
	Push(val V)
	Pop() (val V)
	IsEmpty() bool
	Size() int
}
