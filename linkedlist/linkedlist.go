package linkedlist

import "fmt"

type listNode[T any] struct {
	Value T
	next  *listNode[T]
}

type LinkedList[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	Size int
}

func NewSingly[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list LinkedList[T]) String() string {
	listString := ""

	curr := list.head

	for curr != nil {
		listString += fmt.Sprintf("%v", curr.Value)
		if curr.next != nil {
			listString += " -> "
		}
		curr = curr.next
	}

	return listString
}

func (list *LinkedList[T]) Append(val T) int {
	if list.Size == 0 {
		return list.addToEmpty(val)
	}

	node := listNode[T]{Value: val}
	list.tail.next = &node
	list.tail = &node
	list.Size += 1

	return list.Size
}

func (list *LinkedList[T]) Prepend(val T) int {
	if list.Size == 0 {
		return list.addToEmpty(val)
	}

	node := listNode[T]{Value: val}
	node.next = list.head
	list.head = &node
	list.Size += 1

	return list.Size
}

func (list *LinkedList[T]) InsertAt(index int, val T) (int, error) {
	if list.isOutOfRange(index) {
		return 0, fmt.Errorf("index out of list's range")
	}

	if index == 0 {
		return list.Prepend(val), nil
	}

	if index == list.Size {
		return list.Append(val), nil
	}

	node := listNode[T]{Value: val}
	nodeBefore, found := list.Get(index - 1)
	if !found || nodeBefore == nil {
		return 0, fmt.Errorf("failed to get node at index %d", index-1)
	}
	nodeAfter := nodeBefore.next
	node.next = nodeAfter
	nodeBefore.next = &node

	list.Size += 1

	return list.Size, nil
}

func (list *LinkedList[T]) PopFront() T {
	if list.Size == 1 {
		return list.removeOnlyItem()
	}

	val := list.head.Value
	list.head = list.head.next
	list.Size--
	return val
}

func (list *LinkedList[T]) PopBack() T {
	if list.Size == 1 {
		return list.removeOnlyItem()
	}

	nodeBeforeTail, _ := list.Get(list.Size - 2)
	val := list.tail.Value
	nodeBeforeTail.next = nil
	list.tail = nodeBeforeTail
	list.Size--

	return val
}

func (list *LinkedList[T]) RemoveAt(index int) int {
	return list.Size
}

func (list *LinkedList[T]) Get(index int) (*listNode[T], bool) {
	if index < 0 || index >= list.Size {
		return nil, false
	}

	curr := list.head

	for i := 0; i < list.Size; i++ {
		if i == index {
			return curr, true
		}

		curr = curr.next
	}

	return nil, false
}

func (list *LinkedList[T]) GetFirst() *listNode[T] {
	return list.head
}

func (list *LinkedList[T]) GetLast() *listNode[T] {
	return list.tail
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.Size == 0
}

func (list *LinkedList[T]) addToEmpty(val T) int {
	node := listNode[T]{Value: val}
	list.head = &node
	list.tail = &node
	list.Size += 1

	return list.Size
}

func (list *LinkedList[T]) isOutOfRange(index int) bool {
	return index < 0 || index > list.Size
}

func (list *LinkedList[T]) removeOnlyItem() T {
	val := list.head.Value
	list.tail = nil
	list.head = nil
	list.Size = 0

	return val
}
