// TODO: Implement a function that checks if parentheses, brackets, and braces are balanced using a stack
//
// The function should:
// 1. Use a stack (array) to keep track of opening brackets
// 2. For each opening bracket '(', '[', '{', push it onto the stack
// 3. For each closing bracket ')', ']', '}', check if it matches the most recent opening bracket
// 4. Return true if all brackets are properly balanced, false otherwise
//
// Function signature:
// export function isBalanced(expression: string): boolean
//
// You may use an array as a stack with push() and pop() methods.
// Remember: A stack follows LIFO (Last In, First Out) principle.

class Stack {
    data: Array<string>
    length: number
    constructor() {
        this.data = []
        this.length = 0

    }

    Push(expression: string) {
        this.data.push(expression)
        this.length++
    }

    Pop(): string {
        if (this.length === 0) {
            return ""
        }
        let r = this.data.pop()
        if (r !== undefined) {
            this.length--
            return r
        } else {
            return ""

        }

    }
    Peek(): string {
        if (this.length === 0) {
            return ""
        }
        return this.data[this.length - 1]

    }

}

export function isBalanced(expression: string): boolean {
    // TODO: Implement this function
    const s = new Stack()
    for (let i = 0; i < expression.length; i++) {
        if (isOpeningBracket(expression[i])) {
            s.Push(expression[i])
        } else if (isClosingBracket(expression[i])) {
            if (s.length === 0 || !isMatchingPair(s.Pop(), expression[i])) {
                return false
            }

        }
    }
    return s.length === 0

}

// Helper function - you may implement this if needed
function isOpeningBracket(char: string): boolean {
    // TODO: Check if the character is an opening bracket
    const valid = ["(", "[", "{"]
    return valid.includes(char);
}

// Helper function - you may implement this if needed
function isClosingBracket(char: string): boolean {
    // TODO: Check if the character is a closing bracket
    const valid = [")", "]", "}"]
    return valid.includes(char);
}

// Helper function - you may implement this if needed
function isMatchingPair(opening: string, closing: string): boolean {
    // TODO: Check if opening and closing brackets form a valid pair
    switch (opening) {
        case "(": return closing === ")"
        case "[": return closing === "]"
        case "{": return closing === "}"

    }
    return false;
}
