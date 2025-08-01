# Stack Implementation - Beginner Level

## Problem Description

Implement a stack data structure that supports the fundamental stack operations. A stack follows the Last-In-First-Out (LIFO) principle, where elements are added and removed from the same end (the "top" of the stack).

Your implementation should handle integer values and provide all the core stack operations with proper error handling for edge cases.

## Requirements

Implement a `Stack` class/struct with the following methods:

1. **Constructor/Initialization**: Create a new empty stack
2. **Push**: Add an element to the top of the stack
3. **Pop**: Remove and return the top element (throw/return error if empty)
4. **Peek**: Return the top element without removing it (throw/return error if empty)
5. **IsEmpty**: Check if the stack is empty
6. **Size**: Return the number of elements in the stack
7. **Clear**: Remove all elements from the stack

## Example Usage

### Go Example
```go
stack := NewStack()

// Push elements
stack.Push(10)
stack.Push(20)
stack.Push(30)

fmt.Println(stack.Size())    // Output: 3
fmt.Println(stack.IsEmpty()) // Output: false

// Peek at top element
top, _ := stack.Peek()
fmt.Println(top)            // Output: 30

// Pop elements (LIFO order)
val1, _ := stack.Pop()      // Returns 30
val2, _ := stack.Pop()      // Returns 20
val3, _ := stack.Pop()      // Returns 10

fmt.Println(stack.IsEmpty()) // Output: true
```

### TypeScript Example
```typescript
const stack = new Stack();

// Push elements
stack.push(10);
stack.push(20);
stack.push(30);

console.log(stack.size());    // Output: 3
console.log(stack.isEmpty()); // Output: false

// Peek at top element
console.log(stack.peek());    // Output: 30

// Pop elements (LIFO order)
console.log(stack.pop());     // Output: 30
console.log(stack.pop());     // Output: 20
console.log(stack.pop());     // Output: 10

console.log(stack.isEmpty()); // Output: true
```

## Constraints

- The stack should handle integer values only
- Operations on an empty stack (pop, peek) should return an error or throw an exception
- The stack should be able to grow dynamically (no fixed size limit)
- All operations should be efficient:
  - Push: O(1) time complexity
  - Pop: O(1) time complexity
  - Peek: O(1) time complexity
  - IsEmpty: O(1) time complexity
  - Size: O(1) time complexity
  - Clear: O(1) time complexity

## Test Cases

Your implementation will be tested against the following 10 test cases:

### Test Case 1: New Stack Creation
- **Input**: Create a new stack
- **Expected**: Stack should be empty with size 0

### Test Case 2: Push Single Element
- **Input**: Push(10)
- **Expected**: Stack size = 1, isEmpty = false, peek = 10

### Test Case 3: Push Multiple Elements
- **Input**: Push(5), Push(15), Push(25), Push(35)
- **Expected**: Stack size = 4, peek = 35 (last pushed)

### Test Case 4: Pop Single Element
- **Input**: Push(42), Pop()
- **Expected**: Pop returns 42, stack becomes empty (size = 0)

### Test Case 5: Pop Multiple Elements (LIFO Order)
- **Input**: Push(1), Push(2), Push(3), Push(4), Push(5), then Pop all
- **Expected**: Pop order should be 5, 4, 3, 2, 1

### Test Case 6: Pop from Empty Stack
- **Input**: Pop() on empty stack
- **Expected**: Should return an error/throw exception

### Test Case 7: Peek at Empty Stack
- **Input**: Peek() on empty stack
- **Expected**: Should return an error/throw exception

### Test Case 8: Peek Without Modification
- **Input**: Push(100), Push(200), Push(300), then Peek multiple times
- **Expected**: Peek always returns 300, stack size remains 3

### Test Case 9: Clear Operation
- **Input**: Push several elements, then Clear()
- **Expected**: Stack becomes empty, can be used normally after clear

### Test Case 10: Mixed Operations (LIFO Behavior)
- **Input**: Push(1), Push(2), Pop(), Push(3), Push(4)
- **Expected**: 
  - First pop returns 2
  - Final stack contains [1, 3, 4] (bottom to top)
  - Subsequent pops return 4, 3, 1 in that order

## Implementation Notes

- Use an array/slice as the underlying data structure
- Keep track of the current size for efficient size() operation
- Remember that stack operations work on the "top" end only
- Ensure proper error handling for empty stack operations
- The stack should be able to handle any number of elements (within memory limits)

## Getting Started

1. **Go**: Navigate to the `go/` directory and run `make test` to see the failing tests
2. **TypeScript**: Navigate to the `ts/` directory, run `npm install`, then `npm test` to see the failing tests

Complete the TODO sections in the respective `solution.go` and `solution.ts` files to make all tests pass.