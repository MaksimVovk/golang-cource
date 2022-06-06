package queue

import "testing"

func TextAddQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}

		if !q.Append(i) {
			t.Errorf("Falid to append item %v to queue", i)
		}
	}

	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()

		if !ok {
			t.Errorf("should be able to get item from queue")
		}

		if item != i {
			t.Errorf("got item in wrong order: %v, want %v", item, i)
		}

	}

	item, ok := q.Next()

	if ok {
		t.Errorf("should be not any item in queue, got %v", item)
	}
}
