/**
 * Stack implementation for integer values
 * Implements Last-In-First-Out (LIFO) behavior
 */

class Stack {
    // Properties
    items: Array<number> | null;

    // Constructor
    constructor() {
        this.items = null
    }

    push(data: number): void {
        if (!Array.isArray(this.items)) {
            this.items = [data]
        } else {
            this.items.push(data)
        }
    }

    pop(): number {
        if (!Array.isArray(this.items)) {
            throw new Error("items are empty")
        } else {
            let r = this.items.pop()
            if (r !== undefined) {
                return r
            }
            return 0
        }
    }
    peek(): number {
        if (!Array.isArray(this.items)) {
            throw new Error("items are empty")
        }
        const r = this.items[this.size() - 1]
        return r
    }

    isEmpty(): boolean {
        if (!Array.isArray(this.items)) {
            return true
        }
        return this.size() === 0
    }

    size(): number {
        const r = this.items?.length
        if (r !== undefined) {
            return r
        }
        return 0
    }

    clear(): void {
        this.items = null
    }

}

export { Stack }
