package dequeue

import "sync"

type dequeue[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int64
	mutex  *sync.RWMutex
}

func (dq dequeue[T]) Head() (val Iterator[T]) {
	dq.mutex.RLock()
	val = dq.head
	dq.mutex.RUnlock()
	return
}

func (dq dequeue[T]) Tail() (val Iterator[T]) {
	dq.mutex.RLock()
	val = dq.tail
	dq.mutex.RUnlock()
	return
}

func (dq dequeue[T]) Length() (res int64) {
	dq.mutex.RLock()
	res = dq.length
	dq.mutex.RUnlock()
	return
}

func (dq *dequeue[T]) SafePopFront() (val T) {
	dq.mutex.Lock()
	val = dq.PopFront()
	dq.mutex.Unlock()
	return
}

func (dq *dequeue[T]) PopFront() (val T) {
	val = dq.head.value
	if dq.length == 1 {
		dq.length = 0
		dq.head = nil
		dq.tail = nil
		return
	}
	dq.head = dq.head.next
	dq.head.prev = nil
	dq.length--
	return
}

func (dq *dequeue[T]) SafePopBack() (val T) {
	dq.mutex.Lock()
	val = dq.PopBack()
	dq.mutex.Unlock()
	return
}

func (dq *dequeue[T]) PopBack() (val T) {
	val = dq.tail.value
	if dq.length == 1 {
		dq.length = 0
		dq.head = nil
		dq.tail = nil
		return
	}

	dq.tail = dq.tail.prev
	dq.tail.next = nil
	dq.length--
	return
}

func (dq *dequeue[T]) SafePushFront(val T) {
	dq.mutex.Lock()
	dq.PushFront(val)
	dq.mutex.Unlock()
}

func (dq *dequeue[T]) PushFront(val T) {
	n := new(node[T])
	n.value = val
	if dq.length == 0 {
		dq.head = n
		dq.tail = n
		dq.length++
		return
	}

	n.next = dq.head
	dq.head.prev = n
	dq.head = n
	dq.length++
}

func (dq *dequeue[T]) SafePushBack(val T) {
	dq.mutex.Lock()
	dq.PushBack(val)
	dq.mutex.Unlock()
}

func (dq *dequeue[T]) PushBack(val T) {
	n := new(node[T])
	n.value = val
	if dq.length == 0 {
		dq.head = n
		dq.tail = n
		dq.length++
		return
	}

	n.prev = dq.tail
	dq.tail.next = n
	dq.tail = n
	dq.length++
}
