package queue

type queue[V any] struct {
	head *node[V]
	tail *node[V]
	size int
}

type node[V any] struct {
	next *node[V]
	val  V
}

func New[V any]() Queue[V] {
	return &queue[V]{}
}

func (q *queue[V]) Push(val V) {
	curr := node[V]{
		val: val,
	}
	if q.size == 0 {
		q.head = &curr
		q.tail = &curr
		q.size++
	} else {
		q.size++
		q.tail.next = &curr
		q.tail = &curr
	}
}

func (q *queue[V]) Pop() (val V) {
	if q.size == 0 {
		panic(popOnEmptyQueueMessage)
	}
	q.size--
	val = q.head.val
	q.head = q.head.next
	return
}

func (q *queue[V]) IsEmpty() bool {
	return q.size == 0
}

func (q *queue[V]) Size() int {
	return q.size
}
