import { LinkedList } from './solution';

describe('LinkedList Reverse', () => {
    // Test Case 1: Empty list
    test('Empty list', () => {
        const list = new LinkedList();
        list.reverse();
        expect(list.toArray()).toEqual([]);
    });

    // Test Case 2: Single element
    test('Single element', () => {
        const list = new LinkedList();
        list.append(5);
        list.reverse();
        expect(list.toArray()).toEqual([5]);
    });

    // Test Case 3: Two elements
    test('Two elements', () => {
        const list = new LinkedList();
        list.append(1);
        list.append(2);
        list.reverse();
        expect(list.toArray()).toEqual([2, 1]);
    });

    // Test Case 4: Three elements
    test('Three elements', () => {
        const list = new LinkedList();
        list.append(1);
        list.append(2);
        list.append(3);
        list.reverse();
        expect(list.toArray()).toEqual([3, 2, 1]);
    });

    // Test Case 5: Multiple elements
    test('Multiple elements', () => {
        const list = new LinkedList();
        for (let i = 1; i <= 5; i++) {
            list.append(i);
        }
        list.reverse();
        expect(list.toArray()).toEqual([5, 4, 3, 2, 1]);
    });

    // Test Case 6: Negative numbers
    test('Negative numbers', () => {
        const list = new LinkedList();
        list.append(-1);
        list.append(-2);
        list.append(-3);
        list.reverse();
        expect(list.toArray()).toEqual([-3, -2, -1]);
    });

    // Test Case 7: Mixed positive and negative
    test('Mixed positive and negative', () => {
        const list = new LinkedList();
        list.append(-2);
        list.append(0);
        list.append(3);
        list.append(-1);
        list.reverse();
        expect(list.toArray()).toEqual([-1, 3, 0, -2]);
    });

    // Test Case 8: Duplicate elements
    test('Duplicate elements', () => {
        const list = new LinkedList();
        list.append(1);
        list.append(2);
        list.append(2);
        list.append(1);
        list.reverse();
        expect(list.toArray()).toEqual([1, 2, 2, 1]);
    });

    // Test Case 9: All same elements
    test('All same elements', () => {
        const list = new LinkedList();
        for (let i = 0; i < 4; i++) {
            list.append(7);
        }
        list.reverse();
        expect(list.toArray()).toEqual([7, 7, 7, 7]);
    });

    // Test Case 10: Large numbers
    test('Large numbers', () => {
        const list = new LinkedList();
        list.append(1000000);
        list.append(2000000);
        list.append(3000000);
        list.reverse();
        expect(list.toArray()).toEqual([3000000, 2000000, 1000000]);
    });
});

describe('LinkedList Size', () => {
    // Test size function with different scenarios
    test('Size after reverse operations', () => {
        const list = new LinkedList();
        
        // Empty list
        expect(list.size()).toBe(0);
        
        // Add elements
        list.append(1);
        list.append(2);
        list.append(3);
        expect(list.size()).toBe(3);
        
        // Reverse should not change size
        list.reverse();
        expect(list.size()).toBe(3);
    });
});