# Queue Implementation

## Problem Description

Implement a basic Queue data structure that follows the First-In-First-Out (FIFO) principle. A queue is a linear data structure where elements are added to the back (rear) and removed from the front. Think of it like a line of people waiting - the first person to join the line is the first person to be served.

## Requirements

Implement a Queue class/struct with the following operations:

1. **Enqueue(item)**: Add an item to the back of the queue
2. **Dequeue()**: Remove and return the item from the front of the queue
3. **Front()**: Return the front item without removing it
4. **IsEmpty()**: Check if the queue is empty
5. **Size()**: Return the number of items in the queue

## Example Usage

### Go
```go
q := NewQueue()
q.Enqueue("first")
q.Enqueue("second")
q.Enqueue("third")

fmt.Println(q.Front())    // Output: "first"
fmt.Println(q.Size())     // Output: 3
fmt.Println(q.Dequeue())  // Output: "first"
fmt.Println(q.Size())     // Output: 2
fmt.Println(q.IsEmpty())  // Output: false
```

### TypeScript
```typescript
const queue = new Queue<string>();
queue.enqueue("first");
queue.enqueue("second");
queue.enqueue("third");

console.log(queue.front());    // Output: "first"
console.log(queue.size());     // Output: 3
console.log(queue.dequeue());  // Output: "first"
console.log(queue.size());     // Output: 2
console.log(queue.isEmpty());  // Output: false
```

## Constraints

- The queue should handle any data type (use `interface{}` in Go and generics in TypeScript)
- Return `nil` (Go) or `null` (TypeScript) when attempting to dequeue from an empty queue
- Return `nil` (Go) or `null` (TypeScript) when attempting to get the front of an empty queue
- Size should accurately track the number of elements in the queue
- The implementation should be efficient for the basic operations

## Test Cases

1. **New Empty Queue**: Create a new queue and verify it's empty with size 0
2. **Enqueue Single Item**: Add one item and verify queue is not empty, size is 1, and front returns the item
3. **Enqueue Multiple Items**: Add multiple items and verify size increases and front returns the first item
4. **Dequeue from Non-Empty Queue**: Remove an item and verify it returns the correct item and size decreases
5. **Dequeue from Empty Queue**: Attempt to dequeue from empty queue and verify it returns nil/null
6. **Front on Empty Queue**: Call front on empty queue and verify it returns nil/null
7. **IsEmpty on Empty Queue**: Verify isEmpty returns true for empty queue
8. **IsEmpty on Non-Empty Queue**: Add an item and verify isEmpty returns false
9. **Size Tracking**: Verify size correctly tracks as items are added and removed
10. **FIFO Behavior**: Enqueue multiple items and dequeue them to verify first-in-first-out order

## Implementation Notes

- You can use an array/slice or linked list as the underlying data structure
- Consider the trade-offs between different implementations (array vs linked list)
- Make sure your implementation handles edge cases properly (empty queue operations)
- The queue should maintain FIFO ordering at all times

## Expected Time Complexity

- Enqueue: O(1) amortized
- Dequeue: O(1)
- Front: O(1)
- IsEmpty: O(1)
- Size: O(1)