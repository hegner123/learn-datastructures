// Tree represents a binary search tree with a root node
class Tree {
  root: Node | null;

  constructor(value: number) {
    // Step 1: set this.root equal to createNode(value)
    this.root = null;
  }

  // ValidateBST checks if the tree satisfies BST property
  validateBST(): boolean {
    // Step 1: check if this.root equals null, return true
    // Step 2: call this.root.validateHelper with -Infinity, Infinity
    // Step 3: return result from validateHelper
    return false;
  }

  // FindMin finds the minimum value in the tree
  findMin(): number {
    // Step 1: check if this.root equals null, return -1
    // Step 2: return this.root.findMin()
    return -1;
  }

  // FindMax finds the maximum value in the tree
  findMax(): number {
    // Step 1: check if this.root equals null, return -1
    // Step 2: return this.root.findMax()
    return -1;
  }
}

// Node represents a single node in the binary search tree
class Node {
  value: number;
  left: Node | null;
  right: Node | null;

  constructor(value: number) {
    // Step 1: set this.value equal to value parameter
    // Step 2: set this.left equal to null
    // Step 3: set this.right equal to null
    this.value = 0;
    this.left = null;
    this.right = null;
  }

  // validateHelper is a recursive helper for BST validation
  validateHelper(minVal: number, maxVal: number): boolean {
    // Step 1: check if this equals null, return true
    // Step 2: check if this.value <= minVal, return false
    // Step 3: check if this.value >= maxVal, return false
    // Step 4: create leftValid(boolean) equal to this.left?.validateHelper(minVal, this.value) ?? true
    // Step 5: create rightValid(boolean) equal to this.right?.validateHelper(this.value, maxVal) ?? true
    // Step 6: return leftValid && rightValid
    return false;
  }

  // findMin finds the minimum value in a BST starting from this node
  findMin(): number {
    // Step 1: create current(Node | null) equal to this
    // Step 2: write loop while current.left !== null
    // Step 3: set current equal to current.left
    // Step 4: return current.value
    return -1;
  }

  // findMax finds the maximum value in a BST starting from this node
  findMax(): number {
    // Step 1: create current(Node | null) equal to this
    // Step 2: write loop while current.right !== null
    // Step 3: set current equal to current.right
    // Step 4: return current.value
    return -1;
  }
}

// newTree creates a new BST with the given root value
export function newTree(value: number): Tree {
  // Step 1: create tree(Tree) equal to new Tree(value)
  // Step 2: return tree
  return new Tree(0);
}

// createNode creates a new BST node with the given value
export function createNode(value: number): Node {
  // Step 1: create newNode(Node) equal to new Node(value)
  // Step 2: return newNode
  return new Node(0);
}

export { Tree, Node };