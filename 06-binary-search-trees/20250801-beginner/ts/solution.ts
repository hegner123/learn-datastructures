// TreeNode represents a node in the binary search tree
class TreeNode {
    value: number;
    left: TreeNode | null;
    right: TreeNode | null;

    constructor(value: number) {
        this.value = value;
        this.left = null;
        this.right = null;
    }
}

// BST represents the binary search tree
export class BST {
    private root: TreeNode | null;

    constructor() {
        // TODO: Initialize the BST
        this.root = null;
    }

    // Insert adds a new value to the BST
    insert(value: number): void {
        // TODO: Implement insertion logic
        // Hint: Use recursive insertion maintaining BST property
    }

    // Search finds a value in the BST and returns true if found
    search(value: number): boolean {
        // TODO: Implement search logic
        // Hint: Use BST property to navigate left/right
        return false;
    }

    // InOrderTraversal returns all values in the BST in sorted order
    inOrderTraversal(): number[] {
        // TODO: Implement in-order traversal
        // Hint: Left -> Root -> Right gives sorted order
        return [];
    }

    // Helper method for recursive insertion (optional)
    private insertRecursive(node: TreeNode | null, value: number): TreeNode {
        // TODO: Implement recursive insertion helper
        return new TreeNode(value);
    }

    // Helper method for recursive search (optional)
    private searchRecursive(node: TreeNode | null, value: number): boolean {
        // TODO: Implement recursive search helper
        return false;
    }

    // Helper method for recursive in-order traversal (optional)
    private inOrderRecursive(node: TreeNode | null, result: number[]): void {
        // TODO: Implement recursive in-order traversal helper
    }
}