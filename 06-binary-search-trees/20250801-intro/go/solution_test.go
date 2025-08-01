package main

import (
	"reflect"
	"testing"
)

// Test 1: Node Creation
func TestCreateNode(t *testing.T) {
	node := CreateNode(10)
	if node == nil {
		t.Fatal("CreateNode should return a non-nil node")
	}
	if node.Value != 10 {
		t.Errorf("Expected node value to be 10, got %d", node.Value)
	}
	if node.Left != nil {
		t.Error("Expected left child to be nil")
	}
	if node.Right != nil {
		t.Error("Expected right child to be nil")
	}
}

// Test 2: BST Validation - Single Node
func TestValidateBST_SingleNode(t *testing.T) {
	root := CreateNode(5)
	if !ValidateBST(root) {
		t.Error("Single node should be a valid BST")
	}
}

// Test 3: BST Validation - Valid Small BST
func TestValidateBST_ValidSmallBST(t *testing.T) {
	// Create a valid BST: 3 <- 5 -> 7
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	
	if !ValidateBST(root) {
		t.Error("Valid BST should return true")
	}
}

// Test 4: BST Validation - Invalid BST
func TestValidateBST_InvalidBST(t *testing.T) {
	// Create an invalid BST: 8 <- 5 -> 7 (left child > parent)
	root := CreateNode(5)
	root.Left = CreateNode(8)
	root.Right = CreateNode(7)
	
	if ValidateBST(root) {
		t.Error("Invalid BST should return false")
	}
}

// Test 5: BST Validation - Empty Tree
func TestValidateBST_NilRoot(t *testing.T) {
	if !ValidateBST(nil) {
		t.Error("Empty tree (nil root) should be considered a valid BST")
	}
}

// Test 6: Find Min - Single Node
func TestFindMin_SingleNode(t *testing.T) {
	root := CreateNode(10)
	min := FindMin(root)
	if min != 10 {
		t.Errorf("Expected minimum to be 10, got %d", min)
	}
}

// Test 7: Find Min - Empty Tree
func TestFindMin_EmptyTree(t *testing.T) {
	min := FindMin(nil)
	if min != -1 {
		t.Errorf("Expected -1 for empty tree, got %d", min)
	}
}

// Test 8: Find Max - Single Node
func TestFindMax_SingleNode(t *testing.T) {
	root := CreateNode(15)
	max := FindMax(root)
	if max != 15 {
		t.Errorf("Expected maximum to be 15, got %d", max)
	}
}

// Test 9: Find Max - Empty Tree
func TestFindMax_EmptyTree(t *testing.T) {
	max := FindMax(nil)
	if max != -1 {
		t.Errorf("Expected -1 for empty tree, got %d", max)
	}
}

// Test 10: Find Min/Max - Complex Tree
func TestFindMinMax_ComplexTree(t *testing.T) {
	// Create tree: 1 <- 3 <- 5 -> 7 -> 9
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Left.Left = CreateNode(1)
	root.Right = CreateNode(7)
	root.Right.Right = CreateNode(9)
	
	min := FindMin(root)
	max := FindMax(root)
	
	if min != 1 {
		t.Errorf("Expected minimum to be 1, got %d", min)
	}
	if max != 9 {
		t.Errorf("Expected maximum to be 9, got %d", max)
	}
}

// Additional Tests for Extended BST Features

// Test 11: Insert Operation
func TestInsert(t *testing.T) {
	var root *Node
	root = Insert(root, 5)
	root = Insert(root, 3)
	root = Insert(root, 7)
	
	if root == nil || root.Value != 5 {
		t.Error("Root should be 5")
	}
	if root.Left == nil || root.Left.Value != 3 {
		t.Error("Left child should be 3")
	}
	if root.Right == nil || root.Right.Value != 7 {
		t.Error("Right child should be 7")
	}
}

// Test 12: Search Operation
func TestSearch(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	
	if !Search(root, 5) {
		t.Error("Should find value 5")
	}
	if !Search(root, 3) {
		t.Error("Should find value 3")
	}
	if Search(root, 10) {
		t.Error("Should not find value 10")
	}
}

// Test 13: Delete Operation
func TestDelete(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	
	root = Delete(root, 3)
	if Search(root, 3) {
		t.Error("Value 3 should be deleted")
	}
}

// Test 14: Inorder Traversal
func TestInorderTraversal(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	root.Left.Left = CreateNode(1)
	root.Right.Right = CreateNode(9)
	
	result := InorderTraversal(root)
	expected := []int{1, 3, 5, 7, 9}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test 15: Preorder Traversal
func TestPreorderTraversal(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	
	result := PreorderTraversal(root)
	expected := []int{5, 3, 7}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test 16: Postorder Traversal
func TestPostorderTraversal(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	
	result := PostorderTraversal(root)
	expected := []int{3, 7, 5}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test 17: Height Calculation
func TestHeight(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	root.Left.Left = CreateNode(1)
	
	height := Height(root)
	if height != 2 {
		t.Errorf("Expected height to be 2, got %d", height)
	}
	
	// Test empty tree
	if Height(nil) != -1 {
		t.Error("Expected height of empty tree to be -1")
	}
}

// Test 18: Size Calculation
func TestSize(t *testing.T) {
	root := CreateNode(5)
	root.Left = CreateNode(3)
	root.Right = CreateNode(7)
	root.Left.Left = CreateNode(1)
	
	size := Size(root)
	if size != 4 {
		t.Errorf("Expected size to be 4, got %d", size)
	}
	
	// Test empty tree
	if Size(nil) != 0 {
		t.Error("Expected size of empty tree to be 0")
	}
}