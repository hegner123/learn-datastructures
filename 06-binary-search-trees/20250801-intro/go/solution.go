package main

// TODO: Implement a Binary Search Tree with the following operations:
//
// Step 1: Create a TreeNode struct
//   - Add a Value field of type int to store the node's data
//   - Add a Left field of type *TreeNode to point to the left child
//   - Add a Right field of type *TreeNode to point to the right child
//   - Initialize both Left and Right to nil for new nodes
//
// Step 2: Create a BST struct
//   - Add a Root field of type *TreeNode to point to the root of the tree
//   - Initialize Root to nil for an empty tree
//
// Step 3: Implement Insert method on BST
//   - Create a method signature: func (bst *BST) Insert(value int)
//   - If the tree is empty (Root is nil), create a new TreeNode and set it as Root
//   - Otherwise, call a helper function to recursively insert the value
//   - In the helper function:
//     * If current node is nil, create and return a new TreeNode with the value
//     * If value is less than current node's Value, recursively insert into Left subtree
//     * If value is greater than current node's Value, recursively insert into Right subtree
//     * If value equals current node's Value, do nothing (BST doesn't allow duplicates)
//     * Return the current node to maintain tree structure
//
// Step 4: Implement Search method on BST
//   - Create a method signature: func (bst *BST) Search(value int) bool
//   - Start from the Root node
//   - While current node is not nil:
//     * If value equals current node's Value, return true
//     * If value is less than current node's Value, move to Left child
//     * If value is greater than current node's Value, move to Right child
//   - If we reach nil without finding the value, return false
//
// Step 5: Implement InOrderTraversal method on BST
//   - Create a method signature: func (bst *BST) InOrderTraversal() []int
//   - Create an empty slice to store the result
//   - Use a helper function to recursively traverse:
//     * If current node is nil, return
//     * Recursively traverse Left subtree
//     * Add current node's Value to the result slice
//     * Recursively traverse Right subtree
//   - Return the result slice (will be in ascending order due to BST property)
//
// Step 6: Implement Delete method on BST
//   - Create a method signature: func (bst *BST) Delete(value int)
//   - Use a helper function to recursively delete the node
//   - Handle three cases in the helper function:
//     * Case 1 - Node with no children (leaf node):
//       - Simply remove the node by returning nil
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
// Step 7: Helper function to find minimum value in a subtree
//   - Create a helper function: func findMin(node *TreeNode) *TreeNode
//   - Start from the given node
//   - Keep going left until you reach a node with no left child
//   - Return that node (it contains the minimum value)
//
// Important BST Property to maintain:
// For every node in the tree:
// - All values in the left subtree are less than the node's value
// - All values in the right subtree are greater than the node's value
// - This property must be maintained after every Insert and Delete operation