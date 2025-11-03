package linkedlist

import "fmt"

type listNode[T any] struct {
	value T
	next  *listNode[T]
}

type linkedList[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	Size int
}

func NewSingly[T any]() linkedList[T] {
	return linkedList[T]{}
}

func (list linkedList[T]) String() string {
	listString := ""

	curr := list.head

	for curr != nil {
		listString += fmt.Sprintf("%v", curr.value)
		if curr.next != nil {
			listString += " -> "
		}
		curr = curr.next
	}

	return listString
}

func (list *linkedList[T]) Append(val T) int {
	if list.Size == 0 {
		return list.addToEmpty(val)
	}

	node := listNode[T]{value: val}
	list.tail.next = &node
	list.tail = &node
	list.Size += 1

	return list.Size
}

func (list *linkedList[T]) Prepend(val T) int {
	if list.Size == 0 {
		return list.addToEmpty(val)
	}

	node := listNode[T]{value: val}
	node.next = list.head
	list.head = &node
	list.Size += 1

	return list.Size
}

func (list *linkedList[T]) InsertAt(index int, val T) (int, error) {
	if list.isOutOfRange(index) {
		return 0, fmt.Errorf("index out of list's range")
	}

	if index == 0 {
		return list.Prepend(val), nil
	}

	if index == list.Size {
		return list.Append(val), nil
	}

	node := listNode[T]{value: val}
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

func (list *linkedList[T]) PopFront() T {
	if list.Size == 1 {
		return list.removeOnlyItem()
	}

	val := list.head.value
	list.head = list.head.next
	list.Size--
	return val
}

func (list *linkedList[T]) PopBack() T {
	if list.Size == 1 {
		return list.removeOnlyItem()
	}

	nodeBeforeTail, _ := list.Get(list.Size - 2)
	val := list.tail.value
	nodeBeforeTail.next = nil
	list.tail = nodeBeforeTail
	list.Size--

	return val
}

func (list *linkedList[T]) RemoveAt(index int) int {
	return list.Size
}

func (list *linkedList[T]) Get(index int) (*listNode[T], bool) {
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

func (list *linkedList[T]) GetFirst() *listNode[T] {
	return list.head
}

func (list *linkedList[T]) GetLast() *listNode[T] {
	return list.tail
}

func (list *linkedList[T]) IsEmpty() bool {
	return list.Size == 0
}

func (list *linkedList[T]) addToEmpty(val T) int {
	node := listNode[T]{value: val}
	list.head = &node
	list.tail = &node
	list.Size += 1

	return list.Size
}

func (list *linkedList[T]) isOutOfRange(index int) bool {
	return index < 0 || index > list.Size
}

func (list *linkedList[T]) removeOnlyItem() T {
	val := list.head.value
	list.tail = nil
	list.head = nil
	list.Size = 0

	return val
}
