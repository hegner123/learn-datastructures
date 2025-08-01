import { 
  Node, 
  createNode, 
  validateBST, 
  findMin, 
  findMax, 
  insert, 
  deleteNode, 
  search, 
  inorderTraversal, 
  preorderTraversal, 
  postorderTraversal, 
  height, 
  size 
} from './solution';

describe('BST Introduction Tests', () => {
  // Test 1: Node Creation
  test('createNode creates a node with correct value', () => {
    const node = createNode(10);
    expect(node).toBeDefined();
    expect(node.value).toBe(10);
    expect(node.left).toBeNull();
    expect(node.right).toBeNull();
  });

  // Test 2: BST Validation - Single Node
  test('validateBST returns true for single node', () => {
    const root = createNode(5);
    expect(validateBST(root)).toBe(true);
  });

  // Test 3: BST Validation - Valid Small BST
  test('validateBST returns true for valid small BST', () => {
    // Create a valid BST: 3 <- 5 -> 7
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    
    expect(validateBST(root)).toBe(true);
  });

  // Test 4: BST Validation - Invalid BST
  test('validateBST returns false for invalid BST', () => {
    // Create an invalid BST: 8 <- 5 -> 7 (left child > parent)
    const root = createNode(5);
    root.left = createNode(8);
    root.right = createNode(7);
    
    expect(validateBST(root)).toBe(false);
  });

  // Test 5: BST Validation - Empty Tree
  test('validateBST returns true for empty tree', () => {
    expect(validateBST(null)).toBe(true);
  });

  // Test 6: Find Min - Single Node
  test('findMin returns correct value for single node', () => {
    const root = createNode(10);
    expect(findMin(root)).toBe(10);
  });

  // Test 7: Find Min - Empty Tree
  test('findMin returns -1 for empty tree', () => {
    expect(findMin(null)).toBe(-1);
  });

  // Test 8: Find Max - Single Node
  test('findMax returns correct value for single node', () => {
    const root = createNode(15);
    expect(findMax(root)).toBe(15);
  });

  // Test 9: Find Max - Empty Tree
  test('findMax returns -1 for empty tree', () => {
    expect(findMax(null)).toBe(-1);
  });

  // Test 10: Find Min/Max - Complex Tree
  test('findMin and findMax work for complex tree', () => {
    // Create tree: 1 <- 3 <- 5 -> 7 -> 9
    const root = createNode(5);
    root.left = createNode(3);
    root.left.left = createNode(1);
    root.right = createNode(7);
    root.right.right = createNode(9);
    
    expect(findMin(root)).toBe(1);
    expect(findMax(root)).toBe(9);
  });

  // Additional Tests for Extended BST Features

  // Test 11: Insert Operation
  test('insert operation works correctly', () => {
    let root: Node | null = null;
    root = insert(root, 5);
    root = insert(root, 3);
    root = insert(root, 7);
    
    expect(root).not.toBeNull();
    expect(root!.value).toBe(5);
    expect(root!.left).not.toBeNull();
    expect(root!.left!.value).toBe(3);
    expect(root!.right).not.toBeNull();
    expect(root!.right!.value).toBe(7);
  });

  // Test 12: Search Operation
  test('search operation works correctly', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    
    expect(search(root, 5)).toBe(true);
    expect(search(root, 3)).toBe(true);
    expect(search(root, 7)).toBe(true);
    expect(search(root, 10)).toBe(false);
  });

  // Test 13: Delete Operation
  test('delete operation works correctly', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    
    const updatedRoot = deleteNode(root, 3);
    expect(search(updatedRoot, 3)).toBe(false);
    expect(search(updatedRoot, 5)).toBe(true);
    expect(search(updatedRoot, 7)).toBe(true);
  });

  // Test 14: Inorder Traversal
  test('inorder traversal returns sorted values', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    root.left.left = createNode(1);
    root.right.right = createNode(9);
    
    const result = inorderTraversal(root);
    expect(result).toEqual([1, 3, 5, 7, 9]);
  });

  // Test 15: Preorder Traversal
  test('preorder traversal returns root-first values', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    
    const result = preorderTraversal(root);
    expect(result).toEqual([5, 3, 7]);
  });

  // Test 16: Postorder Traversal
  test('postorder traversal returns children-first values', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    
    const result = postorderTraversal(root);
    expect(result).toEqual([3, 7, 5]);
  });

  // Test 17: Height Calculation
  test('height calculation works correctly', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    root.left.left = createNode(1);
    
    expect(height(root)).toBe(2);
    expect(height(null)).toBe(-1);
  });

  // Test 18: Size Calculation
  test('size calculation works correctly', () => {
    const root = createNode(5);
    root.left = createNode(3);
    root.right = createNode(7);
    root.left.left = createNode(1);
    
    expect(size(root)).toBe(4);
    expect(size(null)).toBe(0);
  });
});