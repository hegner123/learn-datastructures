package main

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("NewQueue should return a non-nil queue")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.Size() != 0 {
		t.Error("New queue should have size 0")
	}
}

func TestEnqueueSingleItem(t *testing.T) {
	q := NewQueue()
	q.Enqueue("hello")
	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}
	if q.Size() != 1 {
		t.Error("Queue size should be 1 after enqueuing one item")
	}
	if q.Front() != "hello" {
		t.Error("Front should return the enqueued item")
	}
}

func TestEnqueueMultipleItems(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Error("Queue size should be 3 after enqueuing three items")
	}
	if q.Front() != 1 {
		t.Error("Front should return the first enqueued item")
	}
}

func TestDequeueFromNonEmptyQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue("first")
	q.Enqueue("second")
	item := q.Dequeue()
	if item != "first" {
		t.Error("Dequeue should return the first enqueued item")
	}
	if q.Size() != 1 {
		t.Error("Queue size should decrease after dequeue")
	}
	if q.Front() != "second" {
		t.Error("Front should now return the second item")
	}
}

func TestDequeueFromEmptyQueue(t *testing.T) {
	q := NewQueue()
	item := q.Dequeue()
	if item != nil {
		t.Error("Dequeue from empty queue should return nil")
	}
}

func TestFrontOnEmptyQueue(t *testing.T) {
	q := NewQueue()
	front := q.Front()
	if front != nil {
		t.Error("Front on empty queue should return nil")
	}
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() {
		t.Error("Empty queue should return true for IsEmpty")
	}
}

func TestIsEmptyOnNonEmptyQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue("item")
	if q.IsEmpty() {
		t.Error("Non-empty queue should return false for IsEmpty")
	}
}

func TestSizeTracking(t *testing.T) {
	q := NewQueue()
	// Size should be 0 initially
	if q.Size() != 0 {
		t.Error("Initial size should be 0")
	}
	
	// Add items and check size
	q.Enqueue(1)
	if q.Size() != 1 {
		t.Error("Size should be 1 after one enqueue")
	}
	
	q.Enqueue(2)
	if q.Size() != 2 {
		t.Error("Size should be 2 after two enqueues")
	}
	
	// Remove items and check size
	q.Dequeue()
	if q.Size() != 1 {
		t.Error("Size should be 1 after one dequeue")
	}
	
	q.Dequeue()
	if q.Size() != 0 {
		t.Error("Size should be 0 after dequeuing all items")
	}
}

func TestFIFOBehavior(t *testing.T) {
	q := NewQueue()
	items := []string{"first", "second", "third", "fourth"}
	
	// Enqueue all items
	for _, item := range items {
		q.Enqueue(item)
	}
	
	// Dequeue all items and verify FIFO order
	for _, expected := range items {
		actual := q.Dequeue()
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
	
	// Queue should be empty now
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all items")
	}
}