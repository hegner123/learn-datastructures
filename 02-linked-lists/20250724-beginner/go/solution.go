package main

// TODO: Implement the Node struct
type Node struct {
	data int
	next *Node
}

// TODO: Implement the LinkedList struct
type LinkedList struct {
	head *Node
}

// TODO: Implement Insert method to add element at beginning
func (ll *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

// TODO: Implement Append method to add element at end
func (ll *LinkedList) Append(data int) {
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

// TODO: Implement Delete method to remove first occurrence of value
func (ll *LinkedList) Delete(data int) bool {
	if ll.head == nil {
		return false
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return true
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// TODO: Implement Search method to check if value exists
func (ll *LinkedList) Search(data int) bool {
	current := ll.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// TODO: Implement Size method to return number of elements

func (ll *LinkedList) Size() int {
	count := 0

	current := ll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// TODO: Implement ToSlice method to return all elements as slice
func (ll *LinkedList) ToSlice() []int {
	r := make([]int,0)
	current := ll.head
	for current != nil {
		r = append(r, current.data)
		current = current.next
	}
	return r
}
