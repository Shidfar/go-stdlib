package atomic

type item[T any] struct {
	value T
}

type queue[T any] struct {
	items   []item[T]
	lessFun LessFun[T]
}

func (q *queue[T]) Len() int {
	return len(q.items)
}

func (q *queue[T]) Head() any {
	return q.items[0]
}

func (q *queue[T]) Less(i, j int) bool {
	return q.lessFun(q.items[i].value, q.items[j].value)
}

func (pq *queue[T]) Pop() any {
	old := pq.items
	n := len(old)
	it := old[n-1]
	pq.items = old[0 : n-1]
	return it
}

func (pq *queue[T]) Push(x any) {
	v := x.(T)
	it := item[T]{v}
	pq.items = append(pq.items, it)
}

func (pq queue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}
