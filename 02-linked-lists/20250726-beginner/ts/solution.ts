// Node class represents a single node in the linked list
class Node {
    data: number;
    next: Node | null;

    constructor(data: number) {
        this.data = data;
        this.next = null;
    }
}

// LinkedList class represents the linked list structure
export class LinkedList {
    private head: Node | null;

    constructor(head?: Node) {
        this.head = head ? head : null;
    }

    // Append adds a new node with the given data at the end of the list
    append(data: number): void {
        const newNode = new Node(data)
        if (this.head === null) {
            this.head = newNode
            return
        }
        let current = this.head
        while (current != null) {
            console.log("append loop  start")
            if (current.next != null) {
                current = current.next
            } else {
                current.next = newNode
                return
            }
        }
    }

    // toArray returns all elements in the list as an array
    toArray(): number[] {
        if (this.head === null) {
            return []
        }
        let current: Node | null = this.head
        let out = []
        while (current !== null) {
            console.log("array loop  start")
            out.push(current.data)
            current = current.next
        }
        console.log(out)
        return out
    }

    // reverse reverses the linked list in-place
    reverse(): void {
        // TODO: Implement in-place reversal of the linked list
        // This should modify the existing list structure, not create a new one
        if (this.head === null) {
            return
        }
        let next: Node | null = null
        let current: Node | null = this.head
        let previous: Node | null = null
        while (current !== null) {
            console.log("reverse loop start")
            next = current.next
            current.next = previous
            previous = current
            current = next
            
        }
        this.head = previous
        return
    }

    // size returns the number of elements in the list
    size(): number {
        let count = 0
        let current: Node | null = this.head
        while (current !== null) {
            count ++
            current = current.next
            
        }
        return  count
    }
}
