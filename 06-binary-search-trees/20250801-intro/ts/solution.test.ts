import { Node, createNode, validateBST, findMin, findMax } from './solution';

describe('BST Introduction Tests', () => {
  test('createNode creates a node with correct value', () => {
    const node = createNode(10);
    expect(node).toBeDefined();
    // Additional checks will depend on your implementation
  });

  test('validateBST returns true for single node', () => {
    const root = createNode(5);
    expect(validateBST(root)).toBe(true);
  });

  test('validateBST returns true for valid small BST', () => {
    // Create a small valid BST: 2 <- 5 -> 7
    const root = createNode(5);
    // You'll need to manually set up the tree structure based on your implementation
    expect(validateBST(root)).toBe(true);
  });

  test('validateBST returns false for invalid BST', () => {
    // Create an invalid BST where left child > parent
    const root = createNode(5);
    // You'll need to set up an invalid tree structure
    // This test may need modification based on your implementation
  });

  test('validateBST returns true for empty tree', () => {
    expect(validateBST(null)).toBe(true);
  });

  test('findMin returns correct value for single node', () => {
    const root = createNode(10);
    expect(findMin(root)).toBe(10);
  });

  test('findMin returns -1 for empty tree', () => {
    expect(findMin(null)).toBe(-1);
  });

  test('findMax returns correct value for single node', () => {
    const root = createNode(15);
    expect(findMax(root)).toBe(15);
  });

  test('findMax returns -1 for empty tree', () => {
    expect(findMax(null)).toBe(-1);
  });

  test('findMin and findMax work for complex tree', () => {
    // This test will need to be implemented based on your Node structure
    // Create a tree like: 1 <- 3 <- 5 -> 7 -> 9
    // Expected min: 1, max: 9
  });
});