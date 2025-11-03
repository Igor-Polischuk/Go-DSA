package queue

import "go-dsa/linkedlist"

type Queue[T any] struct {
	list linkedlist.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.list.Append(value)
}

func (queue *Queue[T]) Dequeue() T {
	return queue.list.PopFront()
}

func (queue *Queue[T]) Peek() T {
	return queue.list.GetFirst().Value
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *Queue[T]) Size() int {
	return queue.list.Size
}
