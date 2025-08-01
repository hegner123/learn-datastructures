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
class LinkedList {
    private head: Node | null;

    constructor() {
        this.head = null;
    }

    // TODO: Implement append method to add elements to the end of the list
    append(data: number): void {
        const newNode = new Node(data)
        if (this.head === null) {
            this.head = newNode
            return
        }
        let current: Node = this.head
        while (current.next !== null) {
            current = current.next
        }
        current.next = newNode
    }

    // TODO: Implement toArray method to convert the linked list to an array
    toArray(): number[] {
        // TODO: Create an empty array
        // TODO: Traverse the list and add each node's data to the array
        // TODO: Return the array
        if (this.head === null) {
            return []
        }
        const out = []
        let current: Node | null = this.head
        while (current !== null) {
            out.push(current.data)
            current = current.next
        }
        return out
    }

    // TODO: Implement static method to create a linked list from an array
    static createFromArray(data: number[]): LinkedList {
        // TODO: Create a new LinkedList
        // TODO: Iterate through the array and append each element to the list
        // TODO: Return the LinkedList
        const output = new LinkedList()
        if (data.length === 0) {
            return output
        }
        for (let i = 0; i < data.length; i++) {
            output.append(data[i])
        }
        return output;
    }

    getHead(): Node | null {
        return this.head
    }
}

// TODO: Implement function to merge two sorted linked lists
export function mergeSortedLists(list1: LinkedList, list2: LinkedList): LinkedList {
    // TODO: Handle edge cases (empty lists)
    // TODO: Create a new result list
    // TODO: Use two pointers to traverse both lists
    // TODO: Compare elements and add the smaller one to the result
    // TODO: Continue until both lists are exhausted
    // TODO: Return the merged list
    const out = new LinkedList()
    const sorted = []

    let currentA = list1.getHead()
    let currentB = list2.getHead()
    if (currentA === null && currentB === null) {
        return out
    }

    while (currentA !== null || currentB !== null) {
        if (currentA === null) {
            sorted.push(currentB!.data);
            currentB = currentB!.next;
        } else if (currentB === null) {
            sorted.push(currentA!.data);
            currentA = currentA!.next;
        } else if (currentA.data < currentB.data) {
            sorted.push(currentA.data);
            currentA = currentA.next;
        } else {
            sorted.push(currentB.data);
            currentB = currentB.next;
        }
    }
    return LinkedList.createFromArray(sorted)
}

export { LinkedList, Node };
