package stack

type Stack[T any] struct {
	element []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(value T) {
	stack.element = append(stack.element, value)
}

func (stack *Stack[T]) Pop() T {
	if stack.IsEmpty() {
		var val T
		return val
	}

	val, index := stack.getLastValueAndIndex()
	stack.element = stack.element[:index]

	return val
}

func (stack Stack[T]) Peek() T {
	if stack.IsEmpty() {
		var val T
		return val
	}

	val, _ := stack.getLastValueAndIndex()

	return val
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.element) == 0
}

func (stack *Stack[T]) Size() int {
	return len(stack.element)
}

func (stack Stack[T]) getLastValueAndIndex() (value T, index int) {
	index = len(stack.element) - 1
	value = stack.element[index]
	return
}
