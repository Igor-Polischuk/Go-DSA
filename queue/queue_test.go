package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := New[int]()

	if !queue.IsEmpty() {
		t.Errorf("expected queue to be empty, got %d", queue.Size())
	}

	if queue.Size() != 0 {
		t.Errorf("expected size 0, got %d", queue.Size())
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if queue.Peek() != 1 {
		t.Errorf("expected peek to be 1, got %d", queue.Peek())
	}

	t.Run("Test size", func(t *testing.T) {
		if queue.Size() != 5 {
			t.Errorf("expected size 3, got %d", queue.Size())
		}
	})

	if queue.IsEmpty() {
		t.Errorf("expected queue to be not empty, got %d", queue.Size())
	}

	dequeue := queue.Dequeue()

	if dequeue != 1 {
		t.Errorf("expected dequeue to be 1, got %d", 1)
	}

	result := []int{}
	for !queue.IsEmpty() {
		result = append(result, queue.Dequeue())
	}

	if len(result) != 4 {
		t.Errorf("expected result to be 4, got %d", len(result))
	}

	if result[0] != 2 || result[1] != 3 || result[2] != 4 || result[3] != 5 {
		t.Errorf("expected result to be 2, 3, 4, 5, got %d, %d, %d, %d", result[0], result[1], result[2], result[3])
	}
}
