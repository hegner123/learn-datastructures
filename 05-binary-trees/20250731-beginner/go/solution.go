package main

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree represents a binary tree structure
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree creates a new empty binary tree
func NewBinaryTree() *BinaryTree {
	// TODO: Implement this method
	return &BinaryTree{Root: nil}
}

// Insert adds a new node to the binary tree using level-order insertion
func (bt *BinaryTree) Insert(data int) {
	// TODO: Implement this method
	newNode := &TreeNode{Data: data, Left: nil, Right: nil}
	if bt.Root == nil {
		bt.Root = newNode
		return
	}
	queue := []*TreeNode{bt.Root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left == nil {
			current.Left = newNode
			return
		} else {
			queue = append(queue, current.Left)
		}
		if current.Right == nil {
			current.Right = newNode
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}

// Height returns the height of the binary tree
// Height is defined as the number of edges in the longest path from root to leaf
// An empty tree has height -1
func (bt *BinaryTree) Height() int {
	// TODO: Implement this method
	return bt.GetHeight(bt.Root)
}
func (bt *BinaryTree) GetHeight(node *TreeNode) int {
	if node == nil {
		return -1
	}
	leftHeight := bt.GetHeight(node.Left)
	rightHeight := bt.GetHeight(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Size returns the number of nodes in the binary tree
func (bt *BinaryTree) Size() int {
	// TODO: Implement this method
	return bt.GetSize(bt.Root)
}

func (bt *BinaryTree) GetSize(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + bt.GetSize(node.Left) + bt.GetSize(node.Right)
}

// IsEmpty returns true if the tree is empty
func (bt *BinaryTree) IsEmpty() bool {
	// TODO: Implement this method

	return bt.Root == nil
}
