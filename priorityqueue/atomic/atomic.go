package atomic

import (
	"github.com/Shidfar/go-stdlib/priorityqueue/heap"
	"sync"
)

type LessFun[T any] func(T, T) bool

type PriorityQueue[T any] struct {
	q       queue[T]
	itemCap int
	lock    sync.RWMutex
}

func NewAtomicPriorityQueue[T any](cap int, less LessFun[T]) PriorityQueue[T] {
	return PriorityQueue[T]{
		itemCap: cap,
		q: queue[T]{
			lessFun: less,
			items:   []item[T]{},
		},
	}
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	heap.Push[T](&pq.q, x)
	if ln := pq.q.Len(); ln > pq.itemCap {
		heap.Pop[T](&pq.q)
	}
}

func (pq *PriorityQueue[T]) Pop() T {
	pq.lock.Lock()
	x := heap.Pop[T](&pq.q)
	pq.lock.Unlock()
	it := x.(item[T])
	return it.value
}

func (pq *PriorityQueue[T]) Len() int {
	pq.lock.RLock()
	defer pq.lock.RUnlock()
	return pq.q.Len()
}

func (pq *PriorityQueue[T]) Head() T {
	pq.lock.RLock()
	x := pq.q.Head()
	pq.lock.RUnlock()
	it := x.(item[T])
	return it.value
}
