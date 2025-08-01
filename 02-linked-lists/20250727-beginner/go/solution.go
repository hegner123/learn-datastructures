package main

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list structure
type LinkedList struct {
	head *Node
}

// TODO: Implement the append method
// append adds a new node with the given data at the end of the list
func (ll *LinkedList) append(data int) {
	// TODO: Implement this method
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// TODO: Implement the findMiddle method
// findMiddle returns the data of the middle node in the linked list
// For lists with even number of elements, return the first of the two middle elements
func (ll *LinkedList) findMiddle() (int, bool) {
	// TODO: Implement this method
	// Return the middle element's data and true if list is not empty
	// Return 0 and false if list is empty
	if ll.head == nil {
		return 0, false
	}
	size := ll.size()
	if size == 1 {
		return ll.head.data, true
	}
	target := (size - 1) / 2
	current := ll.head
	for range target {
		current = current.next
	}
	return current.data, true

}

// TODO: Implement the toSlice method
// toSlice returns all elements in the list as a slice
func (ll *LinkedList) toSlice() []int {
	// TODO: Implement this method
	if ll.head == nil {
		return []int{}
	}
	out := make([]int, 0)
	current := ll.head
	for current != nil {
		out = append(out, current.data)
		current = current.next
	}
	return out
}

// TODO: Implement the size method
// size returns the total number of nodes in the list
func (ll *LinkedList) size() int {
	// TODO: Implement this method
	count := 0
	current := ll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

