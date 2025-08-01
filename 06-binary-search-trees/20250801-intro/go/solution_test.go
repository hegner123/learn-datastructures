package main

import (
	"reflect"
	"testing"
)

func TestBSTInsertAndSearch(t *testing.T) {
	bst := &BST{}
	
	// Test Case 1: Insert single element and search
	bst.Insert(5)
	if !bst.Search(5) {
		t.Errorf("Expected to find 5 in BST")
	}
	if bst.Search(3) {
		t.Errorf("Expected not to find 3 in BST")
	}
}

func TestBSTMultipleInserts(t *testing.T) {
	bst := &BST{}
	
	// Test Case 2: Insert multiple elements
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}
	
	for _, v := range values {
		if !bst.Search(v) {
			t.Errorf("Expected to find %d in BST", v)
		}
	}
}

func TestBSTInOrderTraversal(t *testing.T) {
	bst := &BST{}
	
	// Test Case 3: InOrder traversal should return sorted values
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}
	
	result := bst.InOrderTraversal()
	expected := []int{2, 3, 4, 5, 6, 7, 8}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBSTDeleteLeafNode(t *testing.T) {
	bst := &BST{}
	
	// Test Case 4: Delete leaf node
	values := []int{5, 3, 7, 2, 4}
	for _, v := range values {
		bst.Insert(v)
	}
	
	bst.Delete(2) // Delete leaf node
	if bst.Search(2) {
		t.Errorf("Expected not to find 2 after deletion")
	}
	
	expected := []int{3, 4, 5, 7}
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after deleting leaf, got %v", expected, result)
	}
}

func TestBSTDeleteNodeWithOneChild(t *testing.T) {
	bst := &BST{}
	
	// Test Case 5: Delete node with one child
	values := []int{5, 3, 7, 6}
	for _, v := range values {
		bst.Insert(v)
	}
	
	bst.Delete(7) // Delete node with one left child
	if bst.Search(7) {
		t.Errorf("Expected not to find 7 after deletion")
	}
	
	expected := []int{3, 5, 6}
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after deleting node with one child, got %v", expected, result)
	}
}

func TestBSTDeleteNodeWithTwoChildren(t *testing.T) {
	bst := &BST{}
	
	// Test Case 6: Delete node with two children
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}
	
	bst.Delete(3) // Delete node with two children
	if bst.Search(3) {
		t.Errorf("Expected not to find 3 after deletion")
	}
	
	expected := []int{2, 4, 5, 6, 7, 8}
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after deleting node with two children, got %v", expected, result)
	}
}

func TestBSTDeleteRoot(t *testing.T) {
	bst := &BST{}
	
	// Test Case 7: Delete root node
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}
	
	bst.Delete(5) // Delete root
	if bst.Search(5) {
		t.Errorf("Expected not to find 5 after deletion")
	}
	
	expected := []int{2, 3, 4, 6, 7, 8}
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after deleting root, got %v", expected, result)
	}
}

func TestBSTEmptyTree(t *testing.T) {
	bst := &BST{}
	
	// Test Case 8: Operations on empty tree
	if bst.Search(1) {
		t.Errorf("Expected not to find anything in empty tree")
	}
	
	result := bst.InOrderTraversal()
	if len(result) != 0 {
		t.Errorf("Expected empty traversal for empty tree, got %v", result)
	}
	
	bst.Delete(1) // Should not crash
}

func TestBSTDuplicateInserts(t *testing.T) {
	bst := &BST{}
	
	// Test Case 9: Insert duplicate values
	bst.Insert(5)
	bst.Insert(5) // Duplicate
	bst.Insert(3)
	bst.Insert(3) // Duplicate
	
	expected := []int{3, 5}
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v with no duplicates, got %v", expected, result)
	}
}

func TestBSTSingleNodeOperations(t *testing.T) {
	bst := &BST{}
	
	// Test Case 10: Operations on single node tree
	bst.Insert(42)
	
	if !bst.Search(42) {
		t.Errorf("Expected to find 42 in single node tree")
	}
	
	expected := []int{42}
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v for single node traversal, got %v", expected, result)
	}
	
	bst.Delete(42)
	if bst.Search(42) {
		t.Errorf("Expected not to find 42 after deleting single node")
	}
	
	result = bst.InOrderTraversal()
	if len(result) != 0 {
		t.Errorf("Expected empty tree after deleting single node, got %v", result)
	}
}