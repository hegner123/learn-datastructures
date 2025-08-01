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

// TODO: Implement append method to add elements to the end of the list
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

// TODO: Implement toSlice method to convert the linked list to a slice
func (ll *LinkedList) ToSlice() []int {
	out := make([]int, 0)
	if ll.head == nil {
		return out
	}
	current := ll.head
	for current != nil {
		out = append(out, current.data)
		current = current.next
	}
	return out
}

// TODO: Implement createFromSlice method to create a linked list from a slice
func CreateFromSlice(data []int) *LinkedList {
	list := &LinkedList{head: nil}
	for _, item := range data {
		list.Append(item)
	}
	return list
}

// TODO: Implement mergeSortedLists method to merge two sorted linked lists
func MergeSortedLists(list1, list2 *LinkedList) *LinkedList {
	// TODO: Handle edge cases (empty lists)
	// TODO: Create a new result list
	// TODO: Use two pointers to traverse both lists
	// TODO: Compare elements and add the smaller one to the result
	// TODO: Continue until both lists are exhausted
	// TODO: Return the merged list
	if list1.head == nil && list2.head == nil {
		return &LinkedList{head: nil}
	}
	sortInputted := make([]int, 0)
	var currentA *Node
	var currentB *Node
	if list1.head != nil {
		currentA = list1.head
	}
	if list2.head != nil {

		currentB = list2.head
	}

	for currentA != nil || currentB != nil {
		var a int
		var b int
		if currentA != nil {
			a = currentA.data
		}
		if currentB != nil {
			b = currentB.data
		}
		if currentA != nil && currentB == nil {
			sortInputted = append(sortInputted, a)
			currentA = currentA.next
			continue
		}
		if currentB != nil && currentA == nil {
			sortInputted = append(sortInputted, b)
			currentB = currentB.next
			continue
		}

		if a < b {
			sortInputted = append(sortInputted, a)
			currentA = currentA.next
		} else {
			sortInputted = append(sortInputted, b)
			currentB = currentB.next
		}
	}
	list := CreateFromSlice(sortInputted)

	return list
}
