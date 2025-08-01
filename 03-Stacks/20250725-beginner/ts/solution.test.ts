import { Stack } from './solution';

describe('Stack', () => {
    test('should create a new empty stack', () => {
        const stack = new Stack();
        expect(stack).toBeDefined();
        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
    });

    test('should push a single element', () => {
        const stack = new Stack();
        stack.push(10);
        
        expect(stack.isEmpty()).toBe(false);
        expect(stack.size()).toBe(1);
        expect(stack.peek()).toBe(10);
    });

    test('should push multiple elements', () => {
        const stack = new Stack();
        const elements = [5, 15, 25, 35];
        
        elements.forEach((elem, index) => {
            stack.push(elem);
            expect(stack.size()).toBe(index + 1);
        });
        
        expect(stack.peek()).toBe(35); // Last pushed element
    });

    test('should pop a single element', () => {
        const stack = new Stack();
        stack.push(42);
        
        const popped = stack.pop();
        expect(popped).toBe(42);
        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
    });

    test('should pop multiple elements in LIFO order', () => {
        const stack = new Stack();
        const elements = [1, 2, 3, 4, 5];
        
        elements.forEach(elem => stack.push(elem));
        
        // Pop elements and verify LIFO order
        const expectedOrder = [5, 4, 3, 2, 1];
        expectedOrder.forEach((expected, index) => {
            const popped = stack.pop();
            expect(popped).toBe(expected);
            expect(stack.size()).toBe(elements.length - index - 1);
        });
    });

    test('should throw error when popping from empty stack', () => {
        const stack = new Stack();
        expect(() => stack.pop()).toThrow();
    });

    test('should throw error when peeking at empty stack', () => {
        const stack = new Stack();
        expect(() => stack.peek()).toThrow();
    });

    test('should peek without modifying the stack', () => {
        const stack = new Stack();
        const elements = [100, 200, 300];
        
        elements.forEach(elem => stack.push(elem));
        const initialSize = stack.size();
        
        // Peek multiple times
        for (let i = 0; i < 3; i++) {
            const top = stack.peek();
            expect(top).toBe(300);
            expect(stack.size()).toBe(initialSize);
        }
    });

    test('should clear all elements from stack', () => {
        const stack = new Stack();
        const elements = [7, 14, 21, 28, 35];
        
        elements.forEach(elem => stack.push(elem));
        expect(stack.isEmpty()).toBe(false);
        
        stack.clear();
        
        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
        
        // Verify operations work correctly after clear
        stack.push(99);
        expect(stack.size()).toBe(1);
    });

    test('should maintain LIFO behavior with mixed operations', () => {
        const stack = new Stack();
        
        // Test LIFO with mixed operations
        stack.push(1);
        stack.push(2);
        
        const first = stack.pop();
        expect(first).toBe(2);
        
        stack.push(3);
        stack.push(4);
        
        // Stack should now contain [1, 3, 4] (bottom to top)
        expect(stack.peek()).toBe(4);
        
        const second = stack.pop();
        const third = stack.pop();
        const fourth = stack.pop();
        
        expect(second).toBe(4);
        expect(third).toBe(3);
        expect(fourth).toBe(1);
        
        expect(stack.isEmpty()).toBe(true);
    });
});