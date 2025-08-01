// TODO: Implement a simple BST Node structure and basic validation
// This is a placeholder - implement your solution here

// Node represents a single node in the binary search tree
class Node {
  // TODO: Add properties for value, left child, right child
  
  constructor(value: number) {
    // TODO: Initialize the node with the given value
  }
}

// ValidateBST checks if a given binary tree satisfies BST property
export function validateBST(root: Node | null): boolean {
  // TODO: Implement BST validation logic
  // Return true if the tree maintains BST property, false otherwise
  return false;
}

// CreateNode creates a new BST node with the given value
export function createNode(value: number): Node {
  // TODO: Create and return a new node with the specified value
  return new Node(value);
}

// FindMin finds the minimum value in a BST
export function findMin(root: Node | null): number {
  // TODO: Find and return the minimum value in the BST
  // Return -1 if tree is empty
  return -1;
}

// FindMax finds the maximum value in a BST
export function findMax(root: Node | null): number {
  // TODO: Find and return the maximum value in the BST
  // Return -1 if tree is empty
  return -1;
}

export { Node };