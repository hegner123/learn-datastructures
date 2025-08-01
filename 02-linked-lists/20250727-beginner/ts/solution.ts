// Node class represents a single node in the linked list
class Node {
    data: number;
    next: Node | null;

    constructor(data: number) {
        this.data = data;
        this.next = null;
    }
}

// LinkedList class
export class LinkedList {
    private head: Node | null;

    constructor() {
        this.head = null;
    }

    // TODO: Implement the append method
    // append adds a new node with the given data at the end of the list
    append(data: number): void {
        // TODO: Implement this method
        const newNode = new Node(data)
        if (this.head === null) {
            this.head = newNode
            return
        }
        let current = this.head
        while (current.next !== null) {
            current = current.next
        }
        current.next = newNode
    }

    // TODO: Implement the findMiddle method
    // findMiddle returns the data of the middle node in the linked list
    // For lists with even number of elements, return the first of the two middle elements
    findMiddle(): { data: number; found: boolean } {
        // TODO: Implement this method
        // Return the middle element's data and found=true if list is not empty
        // Return data=0 and found=false if list is empty
        if (this.head == null || this.size() === 0) {
            return { data: 0, found: false };
        }
        const size = this.size()
        const target = size % 2 === 0 ? Math.floor(size / 2) - 1 : Math.floor(size / 2)
        let current = this.head
        for (let i = 0; i < target; i++) {
            current = current.next as Node
        }
        return { data: current.data, found: true };
    }

    // TODO: Implement the toArray method
    // toArray returns all elements in the list as an array
    toArray(): number[] {
        // TODO: Implement this method
        let out = []
        let current = this.head
        while (current != null) {
            out.push(current.data)
            current = current.next
        }

        return out;
    }

    // TODO: Implement the size method
    // size returns the total number of nodes in the list
    size(): number {
        // TODO: Implement this method
        let count = 0
        let current = this.head
        while (current != null) {
            count++
            current = current.next
        }
        return count;
    }
}
