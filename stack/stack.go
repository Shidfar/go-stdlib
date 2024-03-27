package stack

const popOnEmptyStackMessage = "popping on empty stack"

type Stack[V any] interface {
	Push(val V)
	Pop() (val V)
	Head() (val V)
	IsEmpty() bool
	Size() int
}
