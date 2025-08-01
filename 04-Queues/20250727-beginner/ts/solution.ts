// TODO: Implement a Queue data structure with the following operations:
// - enqueue(item): Add an item to the back of the queue
// - dequeue(): Remove and return the item from the front of the queue
// - front(): Return the front item without removing it
// - isEmpty(): Check if the queue is empty
// - size(): Return the number of items in the queue

export class Queue<T> {
  // TODO: Define your queue properties

  constructor() {
    // TODO: Initialize your queue
  }

  // TODO: Implement enqueue method
  enqueue(item: T): void {
    // TODO: Add item to the back of the queue
  }

  // TODO: Implement dequeue method
  dequeue(): T | null {
    // TODO: Remove and return the front item
    // Return null if queue is empty
    return null;
  }

  // TODO: Implement front method
  front(): T | null {
    // TODO: Return the front item without removing it
    // Return null if queue is empty
    return null;
  }

  // TODO: Implement isEmpty method
  isEmpty(): boolean {
    // TODO: Return true if queue is empty, false otherwise
    return false;
  }

  // TODO: Implement size method
  size(): number {
    // TODO: Return the number of items in the queue
    return 0;
  }
}