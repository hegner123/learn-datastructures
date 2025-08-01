
// TreeNode class represents a node in the binary tree
export class Node {
    public data: number;
    public left: Node | null;
    public right: Node | null;

    constructor(data: number) {
        this.data = data;
        this.left = null;
        this.right = null;
    }
}

// BinaryTree class represents a binary tree structure
export class BinaryTree {
    public root: Node | null;

    constructor() {
        this.root = null;
    }

    // Insert adds a new node to the binary tree using level-order insertion
    public insert(data: number): void {
        const newNode = new Node(data)

        if (this.root === null) {
            this.root = newNode
            return
        }
        let queue = [this.root]
        while (queue.length > 0) {
            let current = queue[0]
            queue = queue.slice(1)
            if (current.left === null) {
                current.left = newNode
                return
            } else {
                queue.push(current.left)
            }
            if (current.right === null) {
                current.right = newNode
                return
            } else {
                queue.push(current.right)
            }
        }
    }

    // Height returns the height of the binary tree
    // Height is defined as the number of edges in the longest path from root to leaf
    // An empty tree has height -1
    public height(current: Node | null): number {
        if (current === null) {
            return -1
        }
        let leftHeight = this.height(current.left)
        let rightHeight = this.height(current.right)
        if (leftHeight > rightHeight) {
            return leftHeight + 1
        }
        return rightHeight + 1
    }

    // Size returns the number of nodes in the binary tree
    public size(): number {
        if (this.root === null) {
            return 0
        }
        let size = 0
        let queue = [this.root]
        while (queue.length > 0) {
            let current = queue[0]
            queue = queue.slice(1)
            size++
            if (current.left !== null) {
                queue.push(current.left)
            }
            if (current.right !== null) {
                queue.push(current.right)
            }
        }
        return size;
    }

    // IsEmpty returns true if the tree is empty
    public isEmpty(): boolean {
        if (this.root === null) {
            return true;
        }
        return false;
    }
}
