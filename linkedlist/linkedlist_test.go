package linkedlist

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	list := NewSingly[int]()
	if !list.IsEmpty() {
		t.Errorf("expected empty list")
	}

	list.Append(1)
	if list.IsEmpty() {
		t.Errorf("expected non-empty list after Append()")
	}
}

func TestNewList(t *testing.T) {
	list := NewSingly[int]()

	if !list.IsEmpty() {
		t.Errorf("New list should be empty")
	}

	if list.Size != 0 {
		t.Errorf("expected size 0, got %d", list.Size)
	}

	if list.head != nil || list.tail != nil {
		t.Errorf("expected head and tail to be nil")
	}
}

func TestAppend(t *testing.T) {
	list := NewSingly[int]()
	list.Append(1)

	if list.GetFirst() != list.GetLast() {
		t.Errorf("head and tail should be the same node after Append()")
	}

	if list.IsEmpty() {
		t.Errorf("List shouldn't be empty")
	}

	list.Append(2)
	list.Append(3)

	if list.Size != 3 {
		t.Errorf("expected size 3, got %d", list.Size)
	}

	first := list.GetFirst()
	if first == nil || first.Value != 1 {
		t.Errorf("expected first element to be 1, got %v", first)
	}

	last := list.GetLast()
	if last == nil || last.Value != 3 {
		t.Errorf("expected last element to be 3, got %v", last)
	}
}

func TestPrepend(t *testing.T) {
	list := NewSingly[int]()
	list.Prepend(20)
	list.Prepend(10)

	if list.Size != 2 {
		t.Errorf("expected size 2, got %d", list.Size)
	}

	first := list.GetFirst()
	if first == nil || first.Value != 10 {
		t.Errorf("expected first element 10, got %v", first)
	}

	last := list.GetLast()
	if last == nil || last.Value != 20 {
		t.Errorf("expected last element 20, got %v", last)
	}
}

func TestGet(t *testing.T) {
	list := NewSingly[int]()

	first := list.GetFirst()

	if first != nil {
		t.Errorf("expected first to be nil, got %v", first)
	}

	last := list.GetLast()
	if last != nil {
		t.Errorf("expected last to be nil, got %v", last)
	}

	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Prepend(2)
	list.Prepend(1)

	first = list.GetFirst()
	if first == nil || first.Value != 1 {
		t.Errorf("expected first element to be 1, got %v", first)
	}

	last = list.GetLast()
	if last == nil || last.Value != 5 {
		t.Errorf("expected last element to be 5, got %v", last)
	}

	mid, isPresent := list.Get(2)
	if mid == nil || mid.Value != 3 || !isPresent {
		t.Errorf("expected value at index 2 = 3, got %v", mid)
	}

	_, isPresent = list.Get(100)

	if isPresent {
		t.Errorf("expected Get(100) to be absent")
	}

}

func TestInsertAt(t *testing.T) {
	list := NewSingly[int]()
	list.Append(3)
	list.Prepend(1)

	_, err := list.InsertAt(1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if list.Size != 3 {
		t.Errorf("expected size 3, got %d", list.Size)
	}

	mid, _ := list.Get(1)
	if mid == nil || mid.Value != 2 {
		t.Errorf("expected value at index 1 = 2, got %v", mid)
	}

	_, insertErr := list.InsertAt(1000, 1)

	if insertErr == nil {
		t.Errorf("expected InsertAt(1000, 1) to return an error")
	}
	list.InsertAt(0, 0)
	list.InsertAt(4, 4)

	first := list.GetFirst().Value
	if first != 0 {
		t.Errorf("expected first element to be 0, got %d", first)
	}

	last := list.GetLast()
	if last.Value != 4 {
		t.Errorf("expected last element to be 4, got %d", last.Value)
	}

}

func TestPopFrontAndBack(t *testing.T) {
	list := NewSingly[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)

	deleted := list.PopFront()
	if deleted != 10 {
		t.Errorf("expected deleted value to be 10, got %d", deleted)
	}

	if list.Size != 4 {
		t.Errorf("expected size 2 after PopFront, got %d", list.Size)
	}

	first := list.GetFirst()
	if first.Value != 20 {
		t.Errorf("expected first 20, got %d", first.Value)
	}

	deleted = list.PopBack()
	if deleted != 50 {
		t.Errorf("expected deleted value to be 50, got %d", deleted)
	}

	if list.Size != 3 {
		t.Errorf("expected size 3 after PopBack, got %d", list.Size)
	}

	t.Log(list)
	last := list.GetLast()
	if last.Value != 40 {
		t.Errorf("expected last 40, got %d", last.Value)
	}
}
