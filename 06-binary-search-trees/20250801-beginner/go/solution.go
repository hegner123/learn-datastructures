package main

// TreeNode represents a node in the binary search tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BST represents the binary search tree
type BST struct {
	Root *TreeNode
}

// NewBST creates a new empty binary search tree
func NewBST() *BST {
	// TODO: Initialize and return a new BST
	return nil
}

// Insert adds a new value to the BST
func (bst *BST) Insert(value int) {
	// TODO: Implement insertion logic
	// Hint: Use recursive insertion maintaining BST property
}

// Search finds a value in the BST and returns true if found
func (bst *BST) Search(value int) bool {
	// TODO: Implement search logic
	// Hint: Use BST property to navigate left/right
	return false
}

// InOrderTraversal returns all values in the BST in sorted order
func (bst *BST) InOrderTraversal() []int {
	// TODO: Implement in-order traversal
	// Hint: Left -> Root -> Right gives sorted order
	return nil
}

// helper function for recursive insertion (optional)
func (bst *BST) insertRecursive(node *TreeNode, value int) *TreeNode {
	// TODO: Implement recursive insertion helper
	return nil
}

// helper function for recursive search (optional)
func (bst *BST) searchRecursive(node *TreeNode, value int) bool {
	// TODO: Implement recursive search helper
	return false
}

// helper function for recursive in-order traversal (optional)
func (bst *BST) inOrderRecursive(node *TreeNode, result *[]int) {
	// TODO: Implement recursive in-order traversal helper
}