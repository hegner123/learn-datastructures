// TODO: Implement a Binary Search Tree with the following operations:
//
// Step 1: Create a TreeNode class
//   - Add a value property of type number to store the node's data
//   - Add a left property of type TreeNode | null to point to the left child
//   - Add a right property of type TreeNode | null to point to the right child
//   - In the constructor, initialize left and right to null
//   - Example: constructor(value: number) { this.value = value; this.left = null; this.right = null; }
//
// Step 2: Create a BST class
//   - Add a root property of type TreeNode | null to point to the root of the tree
//   - In the constructor, initialize root to null for an empty tree
//   - Example: constructor() { this.root = null; }
//
// Step 3: Implement insert method in BST class
//   - Create a method signature: insert(value: number): void
//   - If the tree is empty (root is null), create a new TreeNode and set it as root
//   - Otherwise, call a private helper method to recursively insert the value
//   - In the private helper method:
//     * If current node is null, create and return a new TreeNode with the value
//     * If value is less than current node's value, recursively insert into left subtree
//     * If value is greater than current node's value, recursively insert into right subtree
//     * If value equals current node's value, do nothing (BST doesn't allow duplicates)
//     * Return the current node to maintain tree structure
//
// Step 4: Implement search method in BST class
//   - Create a method signature: search(value: number): boolean
//   - Start from the root node
//   - While current node is not null:
//     * If value equals current node's value, return true
//     * If value is less than current node's value, move to left child
//     * If value is greater than current node's value, move to right child
//   - If we reach null without finding the value, return false
//
// Step 5: Implement inOrderTraversal method in BST class
//   - Create a method signature: inOrderTraversal(): number[]
//   - Create an empty array to store the result
//   - Use a private helper method to recursively traverse:
//     * If current node is null, return
//     * Recursively traverse left subtree
//     * Add current node's value to the result array
//     * Recursively traverse right subtree
//   - Return the result array (will be in ascending order due to BST property)
//
// Step 6: Implement delete method in BST class
//   - Create a method signature: delete(value: number): void
//   - Use a private helper method to recursively delete the node
//   - Handle three cases in the helper method:
//     * Case 1 - Node with no children (leaf node):
//       - Simply remove the node by returning null
//     * Case 2 - Node with one child:
//       - Replace the node with its child by returning the child
//     * Case 3 - Node with two children:
//       - Find the inorder successor (smallest value in right subtree)
//       - Replace the node's value with the successor's value
//       - Recursively delete the successor node
//   - To find inorder successor:
//     * Go to right subtree
//     * Keep going left until you find a node with no left child
//
// Step 7: Helper method to find minimum value in a subtree
//   - Create a private helper method: private findMin(node: TreeNode): TreeNode
//   - Start from the given node
//   - Keep going left until you reach a node with no left child
//   - Return that node (it contains the minimum value)
//
// Important TypeScript considerations:
// - Use proper type annotations (number, TreeNode | null, etc.)
// - Handle null checks properly with optional chaining or explicit null checks
// - Mark helper methods as private
// - Export the BST class for testing
//
// Important BST Property to maintain:
// For every node in the tree:
// - All values in the left subtree are less than the node's value
// - All values in the right subtree are greater than the node's value
// - This property must be maintained after every insert and delete operation

export { }; // TODO: Replace with your actual exports