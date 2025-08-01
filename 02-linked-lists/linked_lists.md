# Linked Lists

## What is a Linked List?

A linked list is a linear data structure where elements (nodes) are stored in sequence, but unlike arrays, elements are not stored in contiguous memory locations. Each node contains data and a reference (or link) to the next node in the sequence.

## Key Properties

- **Dynamic size**: Can grow or shrink during runtime
- **Sequential access**: Elements must be accessed sequentially from the head
- **Non-contiguous memory**: Nodes can be stored anywhere in memory
- **Efficient insertion/deletion**: O(1) at known positions

## Types of Linked Lists

- **Singly Linked List**: Each node points to the next node
- **Doubly Linked List**: Each node has pointers to both next and previous nodes
- **Circular Linked List**: The last node points back to the first node

## Time Complexity

- **Access**: O(n)
- **Search**: O(n)
- **Insertion**: O(1) at head, O(n) at arbitrary position
- **Deletion**: O(1) at head, O(n) at arbitrary position

## Go Implementation

### Basic Singly Linked List

```go
package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
    data int
    next *Node
}

// LinkedList represents the linked list structure
type LinkedList struct {
    head *Node
}

// Insert adds a new node at the beginning of the list
func (ll *LinkedList) Insert(data int) {
    newNode := &Node{data: data, next: ll.head}
    ll.head = newNode
}

// Append adds a new node at the end of the list
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

// Delete removes the first occurrence of the specified data
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

// Search finds if a value exists in the list
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

// Display prints all elements in the list
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

// Size returns the number of nodes in the list
func (ll *LinkedList) Size() int {
    count := 0
    current := ll.head
    for current != nil {
        count++
        current = current.next
    }
    return count
}

func main() {
    ll := &LinkedList{}
    
    // Insert elements
    ll.Insert(10)
    ll.Insert(20)
    ll.Insert(30)
    
    fmt.Println("After insertions:")
    ll.Display() // 30 -> 20 -> 10 -> nil
    
    // Append elements
    ll.Append(5)
    ll.Append(1)
    
    fmt.Println("After appending:")
    ll.Display() // 30 -> 20 -> 10 -> 5 -> 1 -> nil
    
    // Search for elements
    fmt.Println("Search 20:", ll.Search(20)) // true
    fmt.Println("Search 100:", ll.Search(100)) // false
    
    // Delete an element
    ll.Delete(20)
    fmt.Println("After deleting 20:")
    ll.Display() // 30 -> 10 -> 5 -> 1 -> nil
    
    fmt.Println("Size:", ll.Size()) // 4
}
```

### Doubly Linked List in Go

```go
package main

import "fmt"

// DoublyNode represents a node in a doubly linked list
type DoublyNode struct {
    data int
    next *DoublyNode
    prev *DoublyNode
}

// DoublyLinkedList represents the doubly linked list structure
type DoublyLinkedList struct {
    head *DoublyNode
    tail *DoublyNode
}

// InsertAtBeginning adds a new node at the beginning
func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
    newNode := &DoublyNode{data: data}
    
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
        return
    }
    
    newNode.next = dll.head
    dll.head.prev = newNode
    dll.head = newNode
}

// InsertAtEnd adds a new node at the end
func (dll *DoublyLinkedList) InsertAtEnd(data int) {
    newNode := &DoublyNode{data: data}
    
    if dll.tail == nil {
        dll.head = newNode
        dll.tail = newNode
        return
    }
    
    dll.tail.next = newNode
    newNode.prev = dll.tail
    dll.tail = newNode
}

// DisplayForward prints the list from head to tail
func (dll *DoublyLinkedList) DisplayForward() {
    current := dll.head
    for current != nil {
        fmt.Printf("%d <-> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

// DisplayBackward prints the list from tail to head
func (dll *DoublyLinkedList) DisplayBackward() {
    current := dll.tail
    for current != nil {
        fmt.Printf("%d <-> ", current.data)
        current = current.prev
    }
    fmt.Println("nil")
}

func main() {
    dll := &DoublyLinkedList{}
    
    dll.InsertAtBeginning(10)
    dll.InsertAtBeginning(20)
    dll.InsertAtEnd(30)
    dll.InsertAtEnd(40)
    
    fmt.Println("Forward:")
    dll.DisplayForward() // 20 <-> 10 <-> 30 <-> 40 <-> nil
    
    fmt.Println("Backward:")
    dll.DisplayBackward() // 40 <-> 30 <-> 10 <-> 20 <-> nil
}
```

## TypeScript Implementation

### Basic Singly Linked List

```typescript
// Node class represents a single node
class Node<T> {
    data: T;
    next: Node<T> | null;
    
    constructor(data: T) {
        this.data = data;
        this.next = null;
    }
}

// LinkedList class
class LinkedList<T> {
    private head: Node<T> | null;
    
    constructor() {
        this.head = null;
    }
    
    // Insert at the beginning
    insert(data: T): void {
        const newNode = new Node<T>(data);
        newNode.next = this.head;
        this.head = newNode;
    }
    
    // Append at the end
    append(data: T): void {
        const newNode = new Node<T>(data);
        
        if (!this.head) {
            this.head = newNode;
            return;
        }
        
        let current = this.head;
        while (current.next) {
            current = current.next;
        }
        current.next = newNode;
    }
    
    // Delete first occurrence of data
    delete(data: T): boolean {
        if (!this.head) return false;
        
        if (this.head.data === data) {
            this.head = this.head.next;
            return true;
        }
        
        let current = this.head;
        while (current.next) {
            if (current.next.data === data) {
                current.next = current.next.next;
                return true;
            }
            current = current.next;
        }
        return false;
    }
    
    // Search for a value
    search(data: T): boolean {
        let current = this.head;
        while (current) {
            if (current.data === data) {
                return true;
            }
            current = current.next;
        }
        return false;
    }
    
    // Display all elements
    display(): void {
        const elements: T[] = [];
        let current = this.head;
        while (current) {
            elements.push(current.data);
            current = current.next;
        }
        console.log(elements.join(' -> ') + ' -> null');
    }
    
    // Get size of the list
    size(): number {
        let count = 0;
        let current = this.head;
        while (current) {
            count++;
            current = current.next;
        }
        return count;
    }
    
    // Convert to array
    toArray(): T[] {
        const result: T[] = [];
        let current = this.head;
        while (current) {
            result.push(current.data);
            current = current.next;
        }
        return result;
    }
    
    // Get first element
    getFirst(): T | null {
        return this.head ? this.head.data : null;
    }
    
    // Get last element
    getLast(): T | null {
        if (!this.head) return null;
        
        let current = this.head;
        while (current.next) {
            current = current.next;
        }
        return current.data;
    }
    
    // Check if list is empty
    isEmpty(): boolean {
        return this.head === null;
    }
}

// Example usage
const ll = new LinkedList<number>();

// Insert elements
ll.insert(10);
ll.insert(20);
ll.insert(30);

console.log("After insertions:");
ll.display(); // 30 -> 20 -> 10 -> null

// Append elements
ll.append(5);
ll.append(1);

console.log("After appending:");
ll.display(); // 30 -> 20 -> 10 -> 5 -> 1 -> null

// Search for elements
console.log("Search 20:", ll.search(20)); // true
console.log("Search 100:", ll.search(100)); // false

// Delete an element
ll.delete(20);
console.log("After deleting 20:");
ll.display(); // 30 -> 10 -> 5 -> 1 -> null

console.log("Size:", ll.size()); // 4
console.log("Array:", ll.toArray()); // [30, 10, 5, 1]
console.log("First:", ll.getFirst()); // 30
console.log("Last:", ll.getLast()); // 1
console.log("Is empty:", ll.isEmpty()); // false
```

### Doubly Linked List in TypeScript

```typescript
// DoublyNode class represents a node in a doubly linked list
class DoublyNode<T> {
    data: T;
    next: DoublyNode<T> | null;
    prev: DoublyNode<T> | null;
    
    constructor(data: T) {
        this.data = data;
        this.next = null;
        this.prev = null;
    }
}

// DoublyLinkedList class
class DoublyLinkedList<T> {
    private head: DoublyNode<T> | null;
    private tail: DoublyNode<T> | null;
    
    constructor() {
        this.head = null;
        this.tail = null;
    }
    
    // Insert at beginning
    insertAtBeginning(data: T): void {
        const newNode = new DoublyNode<T>(data);
        
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
            return;
        }
        
        newNode.next = this.head;
        this.head.prev = newNode;
        this.head = newNode;
    }
    
    // Insert at end
    insertAtEnd(data: T): void {
        const newNode = new DoublyNode<T>(data);
        
        if (!this.tail) {
            this.head = newNode;
            this.tail = newNode;
            return;
        }
        
        this.tail.next = newNode;
        newNode.prev = this.tail;
        this.tail = newNode;
    }
    
    // Insert at specific position (0-indexed)
    insertAtPosition(data: T, position: number): boolean {
        if (position < 0) return false;
        
        if (position === 0) {
            this.insertAtBeginning(data);
            return true;
        }
        
        const newNode = new DoublyNode<T>(data);
        let current = this.head;
        let currentIndex = 0;
        
        while (current && currentIndex < position) {
            current = current.next;
            currentIndex++;
        }
        
        if (currentIndex === position) {
            if (!current) {
                // Insert at end
                this.insertAtEnd(data);
                return true;
            }
            
            newNode.next = current;
            newNode.prev = current.prev;
            
            if (current.prev) {
                current.prev.next = newNode;
            } else {
                this.head = newNode;
            }
            
            current.prev = newNode;
            return true;
        }
        
        return false;
    }
    
    // Display forward
    displayForward(): void {
        const elements: T[] = [];
        let current = this.head;
        while (current) {
            elements.push(current.data);
            current = current.next;
        }
        console.log(elements.join(' <-> ') + ' <-> null');
    }
    
    // Display backward
    displayBackward(): void {
        const elements: T[] = [];
        let current = this.tail;
        while (current) {
            elements.push(current.data);
            current = current.prev;
        }
        console.log(elements.join(' <-> ') + ' <-> null');
    }
    
    // Delete node with specific data
    delete(data: T): boolean {
        if (!this.head) return false;
        
        let current = this.head;
        
        while (current) {
            if (current.data === data) {
                if (current.prev) {
                    current.prev.next = current.next;
                } else {
                    this.head = current.next;
                }
                
                if (current.next) {
                    current.next.prev = current.prev;
                } else {
                    this.tail = current.prev;
                }
                
                return true;
            }
            current = current.next;
        }
        
        return false;
    }
    
    // Search for a value
    search(data: T): boolean {
        let current = this.head;
        while (current) {
            if (current.data === data) {
                return true;
            }
            current = current.next;
        }
        return false;
    }
    
    // Get size of the list
    size(): number {
        let count = 0;
        let current = this.head;
        while (current) {
            count++;
            current = current.next;
        }
        return count;
    }
    
    // Convert to array (forward)
    toArray(): T[] {
        const result: T[] = [];
        let current = this.head;
        while (current) {
            result.push(current.data);
            current = current.next;
        }
        return result;
    }
    
    // Convert to array (backward)
    toArrayReverse(): T[] {
        const result: T[] = [];
        let current = this.tail;
        while (current) {
            result.push(current.data);
            current = current.prev;
        }
        return result;
    }
    
    // Get first element
    getFirst(): T | null {
        return this.head ? this.head.data : null;
    }
    
    // Get last element
    getLast(): T | null {
        return this.tail ? this.tail.data : null;
    }
    
    // Check if list is empty
    isEmpty(): boolean {
        return this.head === null;
    }
    
    // Clear the list
    clear(): void {
        this.head = null;
        this.tail = null;
    }
}

// Example usage
const dll = new DoublyLinkedList<number>();

dll.insertAtBeginning(10);
dll.insertAtBeginning(20);
dll.insertAtEnd(30);
dll.insertAtEnd(40);

console.log("Forward:");
dll.displayForward(); // 20 <-> 10 <-> 30 <-> 40 <-> null

console.log("Backward:");
dll.displayBackward(); // 40 <-> 30 <-> 10 <-> 20 <-> null

// Insert at position
dll.insertAtPosition(25, 2);
console.log("After inserting 25 at position 2:");
dll.displayForward(); // 20 <-> 10 <-> 25 <-> 30 <-> 40 <-> null

// Search for elements
console.log("Search 25:", dll.search(25)); // true
console.log("Search 100:", dll.search(100)); // false

dll.delete(10);
console.log("After deleting 10:");
dll.displayForward(); // 20 <-> 25 <-> 30 <-> 40 <-> null

console.log("Size:", dll.size()); // 4
console.log("Array:", dll.toArray()); // [20, 25, 30, 40]
console.log("Array (reverse):", dll.toArrayReverse()); // [40, 30, 25, 20]
console.log("First:", dll.getFirst()); // 20
console.log("Last:", dll.getLast()); // 40
console.log("Is empty:", dll.isEmpty()); // false

// String example
const stringList = new DoublyLinkedList<string>();
stringList.insertAtEnd("hello");
stringList.insertAtEnd("world");
stringList.insertAtBeginning("hi");

console.log("String list:");
stringList.displayForward(); // hi <-> hello <-> world <-> null
```

## Common Use Cases

- **Dynamic data storage**: When size is unknown at compile time
- **Implementation of other data structures**: Stacks, queues, graphs
- **Undo functionality**: Each node can represent a state
- **Music playlists**: Navigate forward/backward through songs
- **Browser history**: Navigate through visited pages
- **Memory management**: Free memory blocks in operating systems

## Advantages

- Dynamic size adjustment
- Efficient insertion/deletion at known positions
- Memory efficiency (allocate only what's needed)
- Easy implementation of complex data structures

## Disadvantages

- No random access (must traverse from head)
- Extra memory overhead for storing pointers
- Not cache-friendly due to non-contiguous memory
- More complex than arrays for simple operations