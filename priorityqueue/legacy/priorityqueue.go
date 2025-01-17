package legacy

import (
	"container/heap"
	"time"
)

type PriorityQueue struct {
	q       queue
	itemCap int
}

type Item struct {
	Key          string
	LastModified time.Time
	Index        int // The index of the item in the heap.
}

type queue []Item

func NewPriorityQueue(itemCap int) PriorityQueue {
	return PriorityQueue{itemCap: itemCap}
}

func (pq *PriorityQueue) Push(x interface{}) {
	heap.Push(&pq.q, x)
	if ln := pq.q.Len(); ln > pq.itemCap {
		heap.Pop(&pq.q)
	}
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(&pq.q)
}

func (pq *PriorityQueue) Len() int {
	return pq.q.Len()
}

func (pq queue) Len() int {
	return len(pq)
}

func (pq queue) Less(i, j int) bool {
	return pq[i].LastModified.Before(pq[j].LastModified)
}

func (pq *queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *queue) Push(x interface{}) {
	n := len(*pq)
	item := x.(Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
