package main

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	node := CreateNode(10)
	if node == nil {
		t.Fatal("CreateNode should return a non-nil node")
	}
	// Additional checks will depend on your implementation
}

func TestValidateBST_SingleNode(t *testing.T) {
	root := CreateNode(5)
	if !ValidateBST(root) {
		t.Error("Single node should be a valid BST")
	}
}

func TestValidateBST_ValidSmallBST(t *testing.T) {
	// Create a small valid BST: 2 <- 5 -> 7
	root := CreateNode(5)
	// You'll need to manually set up the tree structure based on your implementation
	if !ValidateBST(root) {
		t.Error("Valid BST should return true")
	}
}

func TestValidateBST_InvalidBST(t *testing.T) {
	// Create an invalid BST where left child > parent
	// You'll need to set up an invalid tree structure
	// This test may need modification based on your implementation
}

func TestValidateBST_NilRoot(t *testing.T) {
	if !ValidateBST(nil) {
		t.Error("Empty tree (nil root) should be considered a valid BST")
	}
}

func TestFindMin_SingleNode(t *testing.T) {
	root := CreateNode(10)
	min := FindMin(root)
	if min != 10 {
		t.Errorf("Expected minimum to be 10, got %d", min)
	}
}

func TestFindMin_EmptyTree(t *testing.T) {
	min := FindMin(nil)
	if min != -1 {
		t.Errorf("Expected -1 for empty tree, got %d", min)
	}
}

func TestFindMax_SingleNode(t *testing.T) {
	root := CreateNode(15)
	max := FindMax(root)
	if max != 15 {
		t.Errorf("Expected maximum to be 15, got %d", max)
	}
}

func TestFindMax_EmptyTree(t *testing.T) {
	max := FindMax(nil)
	if max != -1 {
		t.Errorf("Expected -1 for empty tree, got %d", max)
	}
}

func TestFindMinMax_ComplexTree(t *testing.T) {
	// This test will need to be implemented based on your Node structure
	// Create a tree like: 1 <- 3 <- 5 -> 7 -> 9
	// Expected min: 1, max: 9
}