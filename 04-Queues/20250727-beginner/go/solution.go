package main

// TODO: Implement a Queue data structure with the following operations:
// - Enqueue(item): Add an item to the back of the queue
// - Dequeue(): Remove and return the item from the front of the queue
// - Front(): Return the front item without removing it
// - IsEmpty(): Check if the queue is empty
// - Size(): Return the number of items in the queue

// TODO: Implement your Queue structure here
type Queue struct {
	// TODO: Define your queue fields
	data []any
}

// TODO: Implement NewQueue constructor
func NewQueue() *Queue {
	// TODO: Initialize and return a new queue
	queue := &Queue{
		data: nil,
	}
	return queue
}

// TODO: Implement Enqueue method
func (q *Queue) Enqueue(item any) {
	// TODO: Add item to the back of the queue
	if q.data == nil {
		return
	}
	q.data = append(q.data, item)
}

// TODO: Implement Dequeue method
func (q *Queue) Dequeue() any {
	// TODO: Remove and return the front item
	// Return nil if queue is empty
	if q.Size() == 0 {
		return nil
	}
	return nil
}

// TODO: Implement Front method
func (q *Queue) Front() any {
	// TODO: Return the front item without removing it
	// Return nil if queue is empty
	if q.Size() == 0 {
		return nil
	}
	return nil
}

// TODO: Implement IsEmpty method
func (q *Queue) IsEmpty() bool {
	// TODO: Return true if queue is empty, false otherwise
	return q.Size() == 0
}

// TODO: Implement Size method
func (q *Queue) Size() int {
	// TODO: Return the number of items in the queue
	return len(q.data)
}

