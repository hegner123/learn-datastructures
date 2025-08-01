package main

import "math"

// Tree represents a binary search tree with a root node
type Tree struct {
	Root *Node
}

// Node represents a single node in the binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewTree creates a new BST with the given root value
func NewTree(value int) *Tree {
	// Step 1: create tree(*Tree) equal to &Tree{}
	// Step 2: set tree.Root equal to CreateNode(value)
	// Step 3: return tree
	return &Tree{}
}

// CreateNode creates a new BST node with the given value
func CreateNode(value int) *Node {
	// Step 1: create newNode(*Node) equal to &Node{}
	// Step 2: set newNode.Value equal to value parameter
	// Step 3: set newNode.Left equal to nil
	// Step 4: set newNode.Right equal to nil
	// Step 5: return newNode
	return &Node{}
}

// ValidateBST checks if the tree satisfies BST property
func (t *Tree) ValidateBST() bool {
	// Step 1: check if t.Root equals nil, return true
	// Step 2: call t.Root.validateHelper with math.Inf(-1), math.Inf(1)
	// Step 3: return result from validateHelper
	return false
}

// validateHelper is a recursive helper for BST validation
func (n *Node) validateHelper(minVal float64, maxVal float64) bool {
	// Step 1: check if n equals nil, return true
	// Step 2: check if float64(n.Value) <= minVal, return false
	// Step 3: check if float64(n.Value) >= maxVal, return false
	// Step 4: create leftValid(bool) equal to n.Left.validateHelper(minVal, float64(n.Value))
	// Step 5: create rightValid(bool) equal to n.Right.validateHelper(float64(n.Value), maxVal)
	// Step 6: return leftValid && rightValid
	return false
}

// FindMin finds the minimum value in the tree
func (t *Tree) FindMin() int {
	// Step 1: check if t.Root equals nil, return -1
	// Step 2: return t.Root.FindMin()
	return -1
}

// FindMin finds the minimum value in a BST starting from this node
func (n *Node) FindMin() int {
	// Step 1: check if n equals nil, return -1
	// Step 2: create current(*Node) equal to n
	// Step 3: write loop while current.Left != nil
	// Step 4: set current equal to current.Left
	// Step 5: return current.Value
	return -1
}

// FindMax finds the maximum value in the tree
func (t *Tree) FindMax() int {
	// Step 1: check if t.Root equals nil, return -1
	// Step 2: return t.Root.FindMax()
	return -1
}

// FindMax finds the maximum value in a BST starting from this node
func (n *Node) FindMax() int {
	// Step 1: check if n equals nil, return -1
	// Step 2: create current(*Node) equal to n
	// Step 3: write loop while current.Right != nil
	// Step 4: set current equal to current.Right
	// Step 5: return current.Value
	return -1
}