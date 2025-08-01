import { BST } from './solution';

describe('BST Basic Operations', () => {
    
    test('Test 1: Insert single value and search', () => {
        const bst = new BST();
        bst.insert(5);
        expect(bst.search(5)).toBe(true);
    });

    test('Test 2: Insert multiple values and search existing', () => {
        const bst = new BST();
        bst.insert(10);
        bst.insert(5);
        bst.insert(15);
        
        expect(bst.search(5)).toBe(true);
        expect(bst.search(10)).toBe(true);
        expect(bst.search(15)).toBe(true);
    });

    test('Test 3: Search for non-existing value', () => {
        const bst = new BST();
        bst.insert(10);
        bst.insert(20);
        
        expect(bst.search(15)).toBe(false);
    });

    test('Test 4: In-order traversal of simple tree', () => {
        const bst = new BST();
        bst.insert(10);
        bst.insert(5);
        bst.insert(15);
        
        expect(bst.inOrderTraversal()).toEqual([5, 10, 15]);
    });

    test('Test 5: In-order traversal with more complex tree', () => {
        const bst = new BST();
        const values = [50, 30, 20, 40, 70];
        values.forEach(val => bst.insert(val));
        
        expect(bst.inOrderTraversal()).toEqual([20, 30, 40, 50, 70]);
    });

    test('Test 6: Insert duplicate values (should not create duplicates)', () => {
        const bst = new BST();
        bst.insert(10);
        bst.insert(10);
        bst.insert(10);
        
        expect(bst.inOrderTraversal()).toEqual([10]);
    });

    test('Test 7: Left-skewed tree operations', () => {
        const bst = new BST();
        bst.insert(30);
        bst.insert(20);
        bst.insert(10);
        
        expect(bst.search(20)).toBe(true);
        expect(bst.inOrderTraversal()).toEqual([10, 20, 30]);
    });

    test('Test 8: Right-skewed tree operations', () => {
        const bst = new BST();
        bst.insert(10);
        bst.insert(20);
        bst.insert(30);
        
        expect(bst.search(25)).toBe(false);
        expect(bst.inOrderTraversal()).toEqual([10, 20, 30]);
    });

    test('Test 9: Empty tree operations', () => {
        const bst = new BST();
        
        expect(bst.search(10)).toBe(false);
        expect(bst.inOrderTraversal()).toEqual([]);
    });

    test('Test 10: Large set insertion and search', () => {
        const bst = new BST();
        const values = [50, 25, 75, 12, 37, 62, 87];
        values.forEach(val => bst.insert(val));
        
        expect(bst.search(37)).toBe(true);
        expect(bst.search(100)).toBe(false);
        expect(bst.inOrderTraversal()).toEqual([12, 25, 37, 50, 62, 75, 87]);
    });
});

describe('BST Edge Cases', () => {
    test('Empty tree search and traversal', () => {
        const bst = new BST();
        
        expect(bst.search(42)).toBe(false);
        expect(bst.inOrderTraversal()).toEqual([]);
    });

    test('Single node operations', () => {
        const bst = new BST();
        bst.insert(42);
        
        expect(bst.search(42)).toBe(true);
        expect(bst.search(10)).toBe(false);
        expect(bst.inOrderTraversal()).toEqual([42]);
    });
});