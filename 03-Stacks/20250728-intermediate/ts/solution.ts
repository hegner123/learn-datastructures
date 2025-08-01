// MinStack with Undo History - Intermediate Level

/**
 * Operation interface represents a recorded operation for undo functionality
 */
interface Operation {
  // TODO: Define the interface for storing operation information
  // Consider what you need to store to enable undo:
  // - Operation type (push/pop)  
  // - Value involved in the operation
}

/**
 * MinStack class that maintains a stack of integers while tracking the minimum value
 * at each state, combined with an undo history feature that allows reversing operations.
 */
export class MinStack {
  // TODO: Define the properties needed for your MinStack implementation
  // Consider what data structures you need:
  // - A stack for storing values
  // - A stack for tracking minimums
  // - An array for storing operation history

  /**
   * Creates a new empty MinStack with empty history
   */
  constructor() {
    // TODO: Initialize all necessary data structures
  }

  /**
   * Adds an element to the top of the stack and records the operation
   * @param value - The integer value to push
   */
  push(value: number): void {
    // TODO: Implement push operation
    // 1. Add value to the main stack
    // 2. Update minimum tracking
    // 3. Record the operation in history
  }

  /**
   * Removes and returns the top element, recording the operation
   * @returns The popped value
   * @throws Error if stack is empty
   */
  pop(): number {
    // TODO: Implement pop operation
    // 1. Check if stack is empty
    // 2. Remove from main stack and update minimum tracking
    // 3. Record the operation in history (store popped value for undo)
    // 4. Return the popped value
    throw new Error('Stack is empty');
  }

  /**
   * Returns the top element without removing it
   * @returns The top value
   * @throws Error if stack is empty
   */
  peek(): number {
    // TODO: Implement peek operation
    // Return the top element without modifying the stack
    throw new Error('Stack is empty');
  }

  /**
   * Returns the minimum element in the stack
   * @returns The minimum value
   * @throws Error if stack is empty
   */
  getMin(): number {
    // TODO: Implement get minimum operation
    // Return the current minimum value efficiently (O(1))
    throw new Error('Stack is empty');
  }

  /**
   * Checks if the stack is empty
   * @returns True if empty, false otherwise
   */
  isEmpty(): boolean {
    // TODO: Implement empty check
    return false;
  }

  /**
   * Returns the number of elements in the stack
   * @returns The size of the stack
   */
  size(): number {
    // TODO: Implement size operation
    return 0;
  }

  /**
   * Reverses the last operation (push or pop)
   * @returns True if successful, false if no history available
   */
  undo(): boolean {
    // TODO: Implement undo operation
    // 1. Check if there's any operation to undo
    // 2. Get the last operation from history
    // 3. Reverse the operation:
    //    - If last was Push: remove the top element
    //    - If last was Pop: restore the popped element
    // 4. Remove the operation from history
    // 5. Return true if successful, false if no history
    return false;
  }

  /**
   * Returns an array of operation strings for debugging/testing
   * @returns Array of operation descriptions
   */
  getHistory(): string[] {
    // TODO: Implement get history operation
    // Return an array of strings describing each recorded operation
    // Format: "Push(value)" for push operations, "Pop()" for pop operations
    return [];
  }

  /**
   * Clears the operation history but leaves stack contents unchanged
   */
  clearHistory(): void {
    // TODO: Implement clear history operation
    // Clear the history array but don't modify the stack contents
  }
}