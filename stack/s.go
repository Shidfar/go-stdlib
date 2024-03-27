package stack

type stack[V any] struct {
	head *node[V]
	size int
}

type node[V any] struct {
	next *node[V]
	val  V
}

func New[V any]() Stack[V] {
	return &stack[V]{}
}

func (s *stack[V]) Push(val V) {
	curr := node[V]{
		val: val,
	}
	if s.size == 0 {
		s.head = &curr
		s.size++
	} else {
		s.size++
		curr.next = s.head
		s.head = &curr
	}
}

func (s *stack[V]) Pop() (val V) {
	if s.size == 0 {
		panic(popOnEmptyStackMessage)
	}
	s.size--
	val = s.head.val
	s.head = s.head.next
	return
}

func (s *stack[V]) Head() (val V) {
	if s.head != nil {
		val = s.head.val
	}
	return
}

func (s *stack[V]) IsEmpty() bool {
	return s.size == 0
}

func (s *stack[V]) Size() int {
	return s.size
}
