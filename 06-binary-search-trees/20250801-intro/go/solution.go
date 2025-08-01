package main

// TreeNode represents a node in a binary search tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Insert adds a new node with the given value to the BST
// TODO: Implement the insert function
// The function should:
// 1. Create a new TreeNode with the given value if the tree is empty
// 2. Recursively insert the value in the correct position to maintain BST property
// 3. Return the root of the modified tree
func Insert(root *TreeNode, val int) *TreeNode {
	// TODO: Your implementation here
	return nil
}

// Search looks for a value in the BST
// TODO: Implement the search function
// The function should:
// 1. Return true if the value exists in the tree, false otherwise
// 2. Use the BST property to efficiently navigate the tree
func Search(root *TreeNode, val int) bool {
	// TODO: Your implementation here
	return false
}

// InorderTraversal returns a slice of values from in-order traversal
// TODO: Implement the inorder traversal function
// The function should:
// 1. Traverse left subtree
// 2. Visit root node
// 3. Traverse right subtree
// 4. Return values in sorted order (left -> root -> right)
func InorderTraversal(root *TreeNode) []int {
	// TODO: Your implementation here
	return nil
}