# MinStack with Undo History - Intermediate Level

## Problem Description

Implement a MinStack data structure that maintains a stack of integers while tracking the minimum value at each state, combined with an undo history feature that allows reversing operations. This problem combines the LIFO principle of stacks with array-based history tracking, demonstrating how multiple data structures work together to solve complex requirements.

A MinStack should support all standard stack operations while also providing efficient access to the minimum element. Additionally, it should maintain a history of operations that can be undone, restoring the stack to previous states.

## Requirements

Implement a `MinStack` class/struct with the following methods:

1. **Constructor/Initialization**: Create a new empty MinStack with empty history
2. **Push(value)**: Add an element to the top of the stack and record the operation
3. **Pop()**: Remove and return the top element (error if empty), record the operation  
4. **Peek()**: Return the top element without removing it (error if empty)
5. **GetMin()**: Return the minimum element in the stack (error if empty)
6. **IsEmpty()**: Check if the stack is empty
7. **Size()**: Return the number of elements in the stack
8. **Undo()**: Reverse the last operation (push or pop), return true if successful, false if no history
9. **GetHistory()**: Return an array of operation strings for debugging/testing
10. **ClearHistory()**: Clear the operation history (but not the stack contents)

## Example Usage

### Go Example
```go
minStack := NewMinStack()

// Push operations
minStack.Push(5)    // History: ["Push(5)"]
minStack.Push(2)    // History: ["Push(5)", "Push(2)"]  
minStack.Push(8)    // History: ["Push(5)", "Push(2)", "Push(8)"]

fmt.Println(minStack.GetMin())  // Output: 2
fmt.Println(minStack.Peek())    // Output: 8
fmt.Println(minStack.Size())    // Output: 3

// Pop operation
val, _ := minStack.Pop()        // Returns 8, History: ["Push(5)", "Push(2)", "Push(8)", "Pop()"]
fmt.Println(minStack.GetMin())  // Output: 2

// Undo the pop operation
success := minStack.Undo()      // Returns true, restores 8 to stack
fmt.Println(success)            // Output: true
fmt.Println(minStack.Peek())    // Output: 8

// Undo the push operation
minStack.Undo()                 // Removes 8 from stack
fmt.Println(minStack.GetMin())  // Output: 2
```

### TypeScript Example  
```typescript
const minStack = new MinStack();

// Push operations
minStack.push(5);    // History: ["Push(5)"]
minStack.push(2);    // History: ["Push(5)", "Push(2)"]
minStack.push(8);    // History: ["Push(5)", "Push(2)", "Push(8)"]

console.log(minStack.getMin());  // Output: 2
console.log(minStack.peek());    // Output: 8
console.log(minStack.size());    // Output: 3

// Pop operation  
console.log(minStack.pop());     // Output: 8, History: ["Push(5)", "Push(2)", "Push(8)", "Pop()"]
console.log(minStack.getMin());  // Output: 2

// Undo the pop operation
console.log(minStack.undo());    // Output: true, restores 8 to stack
console.log(minStack.peek());    // Output: 8

// Undo the push operation
minStack.undo();                 // Removes 8 from stack  
console.log(minStack.getMin());  // Output: 2
```

## Data Structures Used

This problem combines two fundamental data structures:

1. **Stack (Current Subject)**: The main data structure for LIFO operations
   - Stores the actual integer values
   - Maintains minimum value tracking
   - Supports push, pop, peek operations

2. **Array (Previous Subject)**: For operation history tracking
   - Stores operation history as strings
   - Supports append for new operations
   - Enables undo functionality by reversing operations

## Algorithm Approach

**MinStack Design**:
- Use a main stack for storing values
- Maintain a separate minimum stack to track minimums efficiently
- For each push: add to main stack and update minimum stack
- For each pop: remove from both stacks appropriately

**History Tracking Design**:  
- Use a dynamic array to store operation descriptions
- For Push operations: store "Push(value)"
- For Pop operations: store "Pop()" with the popped value for undo
- Undo reverses the last operation using stored information

## Constraints

- The stack should handle integer values ranging from -1,000,000 to 1,000,000  
- Maximum stack size is 10,000 elements
- Maximum history size is 10,000 operations
- All core operations should be efficient:
  - Push: O(1) time complexity
  - Pop: O(1) time complexity  
  - Peek: O(1) time complexity
  - GetMin: O(1) time complexity
  - Undo: O(1) time complexity
  - GetHistory: O(n) where n is number of operations
- Space complexity: O(n) for n elements plus O(h) for h history operations

## Test Cases

Your implementation will be tested against these 10 test cases:

### Test Case 1: Basic MinStack Operations
- **Input**: Push(3), Push(1), Push(4), GetMin(), Peek()
- **Expected**: GetMin() returns 1, Peek() returns 4, Size() = 3

### Test Case 2: Pop and Min Update  
- **Input**: Push(5), Push(2), Push(7), Pop(), GetMin()
- **Expected**: Pop() returns 7, GetMin() returns 2

### Test Case 3: Multiple Same Minimum Values
- **Input**: Push(2), Push(3), Push(2), Push(1), Push(2), Pop(), GetMin()  
- **Expected**: Pop() returns 2, GetMin() returns 1

### Test Case 4: Basic Undo Push Operation
- **Input**: Push(10), Push(5), Undo(), Peek(), Size()
- **Expected**: After undo, Peek() returns 10, Size() = 1

### Test Case 5: Basic Undo Pop Operation  
- **Input**: Push(1), Push(2), Pop(), Undo(), Peek(), Size()
- **Expected**: After undo, Peek() returns 2, Size() = 2

### Test Case 6: Multiple Undo Operations
- **Input**: Push(1), Push(2), Push(3), Undo(), Undo(), Peek(), GetMin()
- **Expected**: After two undos, Peek() returns 1, GetMin() returns 1

### Test Case 7: Undo on Empty History
- **Input**: Create empty MinStack, Undo()  
- **Expected**: Undo() returns false, stack remains empty

### Test Case 8: History Tracking
- **Input**: Push(5), Push(3), Pop(), Push(7)
- **Expected**: GetHistory() returns ["Push(5)", "Push(3)", "Pop()", "Push(7)"]

### Test Case 9: Clear History Functionality  
- **Input**: Push(1), Push(2), Pop(), ClearHistory(), GetHistory()
- **Expected**: GetHistory() returns empty array, but stack content unchanged

### Test Case 10: Complex Mixed Operations with Undo
- **Input**: Push(4), Push(1), Push(6), Pop(), Undo(), Pop(), GetMin(), Undo(), GetMin()
- **Expected**: 
  - After Pop(): stack has [4, 1], min is 1
  - After Undo(): stack has [4, 1, 6], min is 1  
  - After Pop(): stack has [4, 1], min is 1
  - After Undo(): stack has [4, 1, 6], min is 1

## Implementation Notes

### Stack Design Tips
- Consider using two stacks: one for values, one for minimums
- When pushing: always push to value stack, conditionally push to min stack
- When popping: always pop from value stack, conditionally pop from min stack

### History Design Tips  
- Store operation type and necessary data for reversal
- For Push undo: remove the top element
- For Pop undo: restore the popped element
- Consider storing popped values to enable proper undo

### Error Handling
- Pop/Peek on empty stack should return error
- GetMin on empty stack should return error  
- Undo with no history should return false
- Handle integer overflow/underflow appropriately

## Getting Started

1. **Go**: Navigate to the `go/` directory and run `make test` to see the failing tests
2. **TypeScript**: Navigate to the `ts/` directory, run `npm install`, then `npm test` to see the failing tests  

Complete the TODO sections in the respective `solution.go` and `solution.ts` files to make all tests pass.

## Learning Objectives

This problem helps you understand:
- **Stack Data Structure**: LIFO operations and minimum tracking
- **Array Data Structure**: Dynamic storage for operation history  
- **Data Structure Composition**: How multiple structures solve complex problems
- **Operation Reversal**: Implementing undo functionality with history
- **Efficiency Trade-offs**: Balancing time complexity with space for features
- **State Management**: Maintaining consistent state across operations and undos