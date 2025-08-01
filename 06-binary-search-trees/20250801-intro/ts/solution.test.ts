import { describe, test, expect } from '@jest/globals';
// TODO: Import your BST and TreeNode classes here
// import { BST, TreeNode } from './solution';

describe('Binary Search Tree', () => {
  test('Test Case 1: Insert single element and search', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // bst.insert(5);
    // expect(bst.search(5)).toBe(true);
    // expect(bst.search(3)).toBe(false);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 2: Insert multiple elements', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // const values = [5, 3, 7, 2, 4, 6, 8];
    // values.forEach(v => bst.insert(v));
    // values.forEach(v => {
    //   expect(bst.search(v)).toBe(true);
    // });
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 3: InOrder traversal should return sorted values', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // const values = [5, 3, 7, 2, 4, 6, 8];
    // values.forEach(v => bst.insert(v));
    // const result = bst.inOrderTraversal();
    // const expected = [2, 3, 4, 5, 6, 7, 8];
    // expect(result).toEqual(expected);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 4: Delete leaf node', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // const values = [5, 3, 7, 2, 4];
    // values.forEach(v => bst.insert(v));
    // bst.delete(2); // Delete leaf node
    // expect(bst.search(2)).toBe(false);
    // const expected = [3, 4, 5, 7];
    // const result = bst.inOrderTraversal();
    // expect(result).toEqual(expected);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 5: Delete node with one child', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // const values = [5, 3, 7, 6];
    // values.forEach(v => bst.insert(v));
    // bst.delete(7); // Delete node with one left child
    // expect(bst.search(7)).toBe(false);
    // const expected = [3, 5, 6];
    // const result = bst.inOrderTraversal();
    // expect(result).toEqual(expected);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 6: Delete node with two children', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // const values = [5, 3, 7, 2, 4, 6, 8];
    // values.forEach(v => bst.insert(v));
    // bst.delete(3); // Delete node with two children
    // expect(bst.search(3)).toBe(false);
    // const expected = [2, 4, 5, 6, 7, 8];
    // const result = bst.inOrderTraversal();
    // expect(result).toEqual(expected);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 7: Delete root node', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // const values = [5, 3, 7, 2, 4, 6, 8];
    // values.forEach(v => bst.insert(v));
    // bst.delete(5); // Delete root
    // expect(bst.search(5)).toBe(false);
    // const expected = [2, 3, 4, 6, 7, 8];
    // const result = bst.inOrderTraversal();
    // expect(result).toEqual(expected);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 8: Operations on empty tree', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // expect(bst.search(1)).toBe(false);
    // const result = bst.inOrderTraversal();
    // expect(result).toEqual([]);
    // bst.delete(1); // Should not crash
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 9: Insert duplicate values', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // bst.insert(5);
    // bst.insert(5); // Duplicate
    // bst.insert(3);
    // bst.insert(3); // Duplicate
    // const expected = [3, 5];
    // const result = bst.inOrderTraversal();
    // expect(result).toEqual(expected);
    expect(true).toBe(true); // Remove this line when implementing
  });

  test('Test Case 10: Operations on single node tree', () => {
    // TODO: Uncomment and fix imports
    // const bst = new BST();
    // bst.insert(42);
    // expect(bst.search(42)).toBe(true);
    // const expected = [42];
    // let result = bst.inOrderTraversal();
    // expect(result).toEqual(expected);
    // bst.delete(42);
    // expect(bst.search(42)).toBe(false);
    // result = bst.inOrderTraversal();
    // expect(result).toEqual([]);
    expect(true).toBe(true); // Remove this line when implementing
  });
});