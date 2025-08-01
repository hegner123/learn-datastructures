// TreeNode represents a node in a binary search tree
export class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;

  constructor(val: number, left: TreeNode | null = null, right: TreeNode | null = null) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

// Insert adds a new node with the given value to the BST
// TODO: Implement the insert function
// The function should:
// 1. Create a new TreeNode with the given value if the tree is empty
// 2. Recursively insert the value in the correct position to maintain BST property
// 3. Return the root of the modified tree
export function insert(root: TreeNode | null, val: number): TreeNode | null {
  // TODO: Your implementation here
  return null;
}

// Search looks for a value in the BST
// TODO: Implement the search function
// The function should:
// 1. Return true if the value exists in the tree, false otherwise
// 2. Use the BST property to efficiently navigate the tree
export function search(root: TreeNode | null, val: number): boolean {
  // TODO: Your implementation here
  return false;
}

// InorderTraversal returns an array of values from in-order traversal
// TODO: Implement the inorder traversal function
// The function should:
// 1. Traverse left subtree
// 2. Visit root node
// 3. Traverse right subtree
// 4. Return values in sorted order (left -> root -> right)
export function inorderTraversal(root: TreeNode | null): number[] {
  // TODO: Your implementation here
  return [];
}