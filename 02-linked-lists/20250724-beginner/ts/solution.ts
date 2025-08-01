// TODO: Implement the Node class
class Node<T> {
    data: T;
    next: Node<T> | null
    constructor(data: T) {
        this.data = data
        this.next = null
    }
}
// TODO: Implement the LinkedList class
class LinkedList<T> {
    private head: Node<T> | null
    constructor() {
        this.head = null
    }
    // TODO: Implement insert method to add element at beginning
    insert(data: T): void {
        const newNode = new Node<T>(data)
        newNode.next = this.head
        this.head = newNode
    }
    // TODO: Implement append method to add element at end
    append(data: T): void {
        const newNode = new Node<T>(data)
        if (!this.head) {
            this.head = newNode
            return
        }

        let current = this.head
        while (current.next) {
            current = current.next
        }
        current.next = newNode
    }
    // TODO: Implement delete method to remove first occurrence of value
    delete(data: T): boolean {
        if (!this.head) return false

        if (this.head.data === data) {
            this.head = this.head.next
            return true
        }

        let current = this.head
        while (current.next) {
            if (current.next.data === data) {
                current.next = current.next.next
                return true
            }
            current = current.next
        }
        return false
    }
    // TODO: Implement search method to check if value exists
    search(data: T): boolean {
        let current = this.head
        while (current) {
            if (current.data === data) {
                return true
            }
            current = current.next
        }
        return false
    }
    // TODO: Implement size method to return number of elements
    size(): number {
        let count = 0;
        let current = this.head
        while (current) {
            count++
            current = current.next
        }

        return count
    }
    // TODO: Implement toArray method to return all elements as array
    toArray(): T[] {
        const result: T[] = []
        let current = this.head
        while (current) {
            result.push(current.data)
            current = current.next
        }
        return result
    }
}

export { LinkedList };
