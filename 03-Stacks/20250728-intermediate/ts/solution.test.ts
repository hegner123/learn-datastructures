import { MinStack } from './solution';

describe('MinStack', () => {
  test('Test Case 1: Basic MinStack Operations', () => {
    const minStack = new MinStack();
    
    // Push operations
    minStack.push(3);
    minStack.push(1);
    minStack.push(4);
    
    // Test GetMin
    expect(minStack.getMin()).toBe(1);
    
    // Test Peek
    expect(minStack.peek()).toBe(4);
    
    // Test Size
    expect(minStack.size()).toBe(3);
  });

  test('Test Case 2: Pop and Min Update', () => {
    const minStack = new MinStack();
    
    minStack.push(5);
    minStack.push(2);
    minStack.push(7);
    
    // Pop operation
    expect(minStack.pop()).toBe(7);
    
    // Check min after pop
    expect(minStack.getMin()).toBe(2);
  });

  test('Test Case 3: Multiple Same Minimum Values', () => {
    const minStack = new MinStack();
    
    minStack.push(2);
    minStack.push(3);
    minStack.push(2);
    minStack.push(1);
    minStack.push(2);
    
    // Pop the top 2
    expect(minStack.pop()).toBe(2);
    
    // Min should still be 1
    expect(minStack.getMin()).toBe(1);
  });

  test('Test Case 4: Basic Undo Push Operation', () => {
    const minStack = new MinStack();
    
    minStack.push(10);
    minStack.push(5);
    
    // Undo the last push
    expect(minStack.undo()).toBe(true);
    
    // Check result
    expect(minStack.peek()).toBe(10);
    expect(minStack.size()).toBe(1);
  });

  test('Test Case 5: Basic Undo Pop Operation', () => {
    const minStack = new MinStack();
    
    minStack.push(1);
    minStack.push(2);
    minStack.pop();
    
    // Undo the pop
    expect(minStack.undo()).toBe(true);
    
    // Check result
    expect(minStack.peek()).toBe(2);
    expect(minStack.size()).toBe(2);
  });

  test('Test Case 6: Multiple Undo Operations', () => {
    const minStack = new MinStack();
    
    minStack.push(1);
    minStack.push(2);
    minStack.push(3);
    
    // Undo twice
    minStack.undo();
    minStack.undo();
    
    // Check result
    expect(minStack.peek()).toBe(1);
    expect(minStack.getMin()).toBe(1);
  });

  test('Test Case 7: Undo on Empty History', () => {
    const minStack = new MinStack();
    
    // Try to undo with no history
    expect(minStack.undo()).toBe(false);
    
    // Stack should remain empty
    expect(minStack.isEmpty()).toBe(true);
  });

  test('Test Case 8: History Tracking', () => {
    const minStack = new MinStack();
    
    minStack.push(5);
    minStack.push(3);
    minStack.pop();
    minStack.push(7);
    
    const history = minStack.getHistory();
    const expected = ['Push(5)', 'Push(3)', 'Pop()', 'Push(7)'];
    
    expect(history).toEqual(expected);
  });

  test('Test Case 9: Clear History Functionality', () => {
    const minStack = new MinStack();
    
    minStack.push(1);
    minStack.push(2);
    minStack.pop();
    
    // Clear history
    minStack.clearHistory();
    
    // History should be empty
    expect(minStack.getHistory()).toEqual([]);
    
    // But stack content should be unchanged
    expect(minStack.size()).toBe(1);
    expect(minStack.peek()).toBe(1);
  });

  test('Test Case 10: Complex Mixed Operations with Undo', () => {
    const minStack = new MinStack();
    
    minStack.push(4);
    minStack.push(1);
    minStack.push(6);
    
    // Pop
    minStack.pop();
    // Stack: [4, 1], min: 1
    expect(minStack.getMin()).toBe(1);
    
    // Undo the pop
    minStack.undo();
    // Stack: [4, 1, 6], min: 1
    expect(minStack.getMin()).toBe(1);
    
    // Pop again
    minStack.pop();
    // Stack: [4, 1], min: 1
    expect(minStack.getMin()).toBe(1);
    
    // Undo the pop again
    minStack.undo();
    // Stack: [4, 1, 6], min: 1
    expect(minStack.getMin()).toBe(1);
  });

  test('Error Cases: Empty Stack Operations', () => {
    const minStack = new MinStack();
    
    // Pop from empty stack
    expect(() => minStack.pop()).toThrow('Stack is empty');
    
    // Peek at empty stack
    expect(() => minStack.peek()).toThrow('Stack is empty');
    
    // GetMin on empty stack
    expect(() => minStack.getMin()).toThrow('Stack is empty');
  });
});