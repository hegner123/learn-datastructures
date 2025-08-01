# Stacks

## What is a Stack?

A stack is a linear data structure that follows the Last In, First Out (LIFO) principle. Elements are added and removed from the same end, called the "top" of the stack. Think of it like a stack of plates - you can only add or remove plates from the top.

## Key Properties

- **LIFO (Last In, First Out)**: The last element added is the first one to be removed
- **Single access point**: Elements can only be added or removed from the top
- **Dynamic size**: Can grow and shrink during runtime
- **No random access**: Cannot access elements in the middle directly

## Core Operations

- **Push**: Add an element to the top of the stack
- **Pop**: Remove and return the top element
- **Peek/Top**: View the top element without removing it
- **IsEmpty**: Check if the stack is empty
- **Size**: Get the number of elements in the stack

## Time Complexity

- **Push**: O(1)
- **Pop**: O(1)
- **Peek**: O(1)
- **Search**: O(n)

## Go Implementation

### Array-based Stack

```go
package main

import (
    "errors"
    "fmt"
)

// Stack represents a stack data structure
type Stack struct {
    items []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
    return &Stack{
        items: make([]int, 0),
    }
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *Stack) Pop() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }
    
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, nil
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }
    
    return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
    return len(s.items)
}

// Display prints all elements in the stack
func (s *Stack) Display() {
    if s.IsEmpty() {
        fmt.Println("Stack is empty")
        return
    }
    
    fmt.Println("Stack (top to bottom):")
    for i := len(s.items) - 1; i >= 0; i-- {
        fmt.Printf("| %d |\n", s.items[i])
    }
    fmt.Println("-----")
}

func main() {
    stack := NewStack()
    
    // Push elements
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)
    stack.Push(40)
    
    fmt.Println("After pushing 10, 20, 30, 40:")
    stack.Display()
    
    // Peek at top element
    if top, err := stack.Peek(); err == nil {
        fmt.Printf("Top element: %d\n", top)
    }
    
    // Pop elements
    if item, err := stack.Pop(); err == nil {
        fmt.Printf("Popped: %d\n", item)
    }
    
    if item, err := stack.Pop(); err == nil {
        fmt.Printf("Popped: %d\n", item)
    }
    
    fmt.Println("After popping twice:")
    stack.Display()
    
    fmt.Printf("Stack size: %d\n", stack.Size())
    fmt.Printf("Is empty: %t\n", stack.IsEmpty())
}
```

### Generic Stack in Go

```go
package main

import (
    "errors"
    "fmt"
)

// GenericStack represents a generic stack data structure
type GenericStack[T any] struct {
    items []T
}

// NewGenericStack creates a new empty generic stack
func NewGenericStack[T any]() *GenericStack[T] {
    return &GenericStack[T]{
        items: make([]T, 0),
    }
}

// Push adds an element to the top of the stack
func (s *GenericStack[T]) Push(item T) {
    s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *GenericStack[T]) Pop() (T, error) {
    var zero T
    if s.IsEmpty() {
        return zero, errors.New("stack is empty")
    }
    
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, nil
}

// Peek returns the top element without removing it
func (s *GenericStack[T]) Peek() (T, error) {
    var zero T
    if s.IsEmpty() {
        return zero, errors.New("stack is empty")
    }
    
    return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *GenericStack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *GenericStack[T]) Size() int {
    return len(s.items)
}

func main() {
    // String stack
    stringStack := NewGenericStack[string]()
    stringStack.Push("hello")
    stringStack.Push("world")
    stringStack.Push("!")
    
    fmt.Println("String stack operations:")
    for !stringStack.IsEmpty() {
        if item, err := stringStack.Pop(); err == nil {
            fmt.Printf("Popped: %s\n", item)
        }
    }
    
    // Float stack
    floatStack := NewGenericStack[float64]()
    floatStack.Push(3.14)
    floatStack.Push(2.71)
    floatStack.Push(1.41)
    
    fmt.Println("\nFloat stack operations:")
    for !floatStack.IsEmpty() {
        if item, err := floatStack.Pop(); err == nil {
            fmt.Printf("Popped: %.2f\n", item)
        }
    }
}
```

### Stack Applications in Go

```go
package main

import (
    "fmt"
    "strings"
)

// Check if parentheses are balanced
func isBalanced(expression string) bool {
    stack := NewStack()
    
    for _, char := range expression {
        switch char {
        case '(', '[', '{':
            stack.Push(int(char))
        case ')', ']', '}':
            if stack.IsEmpty() {
                return false
            }
            
            top, _ := stack.Pop()
            if !isMatchingPair(rune(top), char) {
                return false
            }
        }
    }
    
    return stack.IsEmpty()
}

func isMatchingPair(open, close rune) bool {
    return (open == '(' && close == ')') ||
           (open == '[' && close == ']') ||
           (open == '{' && close == '}')
}

// Reverse a string using stack
func reverseString(s string) string {
    stack := NewGenericStack[rune]()
    
    // Push all characters onto stack
    for _, char := range s {
        stack.Push(char)
    }
    
    // Pop all characters to create reversed string
    var result strings.Builder
    for !stack.IsEmpty() {
        if char, err := stack.Pop(); err == nil {
            result.WriteRune(char)
        }
    }
    
    return result.String()
}

// Evaluate postfix expression
func evaluatePostfix(expression string) int {
    stack := NewStack()
    tokens := strings.Fields(expression)
    
    for _, token := range tokens {
        switch token {
        case "+", "-", "*", "/":
            if stack.Size() < 2 {
                fmt.Println("Invalid expression")
                return 0
            }
            
            operand2, _ := stack.Pop()
            operand1, _ := stack.Pop()
            
            var result int
            switch token {
            case "+":
                result = operand1 + operand2
            case "-":
                result = operand1 - operand2
            case "*":
                result = operand1 * operand2
            case "/":
                if operand2 != 0 {
                    result = operand1 / operand2
                }
            }
            
            stack.Push(result)
        default:
            // Convert string to number and push
            var num int
            fmt.Sscanf(token, "%d", &num)
            stack.Push(num)
        }
    }
    
    if result, err := stack.Pop(); err == nil {
        return result
    }
    return 0
}

func main() {
    // Test balanced parentheses
    expressions := []string{
        "(())",
        "([{}])",
        "(()",
        "([)]",
        "{[()]}",
    }
    
    fmt.Println("Balanced parentheses check:")
    for _, expr := range expressions {
        fmt.Printf("%s: %t\n", expr, isBalanced(expr))
    }
    
    // Test string reversal
    fmt.Println("\nString reversal:")
    original := "Hello, World!"
    reversed := reverseString(original)
    fmt.Printf("Original: %s\n", original)
    fmt.Printf("Reversed: %s\n", reversed)
    
    // Test postfix evaluation
    fmt.Println("\nPostfix expression evaluation:")
    postfix := "3 4 + 2 * 7 /"
    result := evaluatePostfix(postfix)
    fmt.Printf("Expression: %s\n", postfix)
    fmt.Printf("Result: %d\n", result)
}
```

## TypeScript Implementation

### Basic Stack

```typescript
class Stack<T> {
    private items: T[];
    
    constructor() {
        this.items = [];
    }
    
    // Add element to top of stack
    public push(item: T): void {
        this.items.push(item);
    }
    
    // Remove and return top element
    public pop(): T {
        if (this.isEmpty()) {
            throw new Error("Stack is empty");
        }
        return this.items.pop()!;
    }
    
    // View top element without removing
    public peek(): T {
        if (this.isEmpty()) {
            throw new Error("Stack is empty");
        }
        return this.items[this.items.length - 1];
    }
    
    // Check if stack is empty
    public isEmpty(): boolean {
        return this.items.length === 0;
    }
    
    // Get stack size
    public size(): number {
        return this.items.length;
    }
    
    // Display stack contents
    public display(): void {
        if (this.isEmpty()) {
            console.log("Stack is empty");
            return;
        }
        
        console.log("Stack (top to bottom):");
        for (let i = this.items.length - 1; i >= 0; i--) {
            console.log(`| ${this.items[i]} |`);
        }
        console.log("-----");
    }
    
    // Clear all elements
    public clear(): void {
        this.items = [];
    }
    
    // Convert to array
    public toArray(): T[] {
        return [...this.items];
    }
}

// Example usage
const stack = new Stack<number>();

// Push elements
stack.push(10);
stack.push(20);
stack.push(30);
stack.push(40);

console.log("After pushing 10, 20, 30, 40:");
stack.display();

// Peek at top element
console.log("Top element:", stack.peek());

// Pop elements
console.log("Popped:", stack.pop());
console.log("Popped:", stack.pop());

console.log("After popping twice:");
stack.display();

console.log("Stack size:", stack.size());
console.log("Is empty:", stack.isEmpty());
```

### Stack Applications in TypeScript

```typescript
// Check if parentheses are balanced
function isBalanced(expression: string): boolean {
    const stack = new Stack<string>();
    const pairs: { [key: string]: string } = {
        ')': '(',
        ']': '[',
        '}': '{'
    };
    
    for (const char of expression) {
        if (char === '(' || char === '[' || char === '{') {
            stack.push(char);
        } else if (char === ')' || char === ']' || char === '}') {
            if (stack.isEmpty() || stack.pop() !== pairs[char]) {
                return false;
            }
        }
    }
    
    return stack.isEmpty();
}

// Reverse a string using stack
function reverseString(str: string): string {
    const stack = new Stack<string>();
    
    // Push all characters onto stack
    for (const char of str) {
        stack.push(char);
    }
    
    // Pop all characters to create reversed string
    let reversed = '';
    while (!stack.isEmpty()) {
        reversed += stack.pop();
    }
    
    return reversed;
}

// Convert infix to postfix notation
function infixToPostfix(expression: string): string {
    const stack = new Stack<string>();
    const precedence: { [key: string]: number } = {
        '+': 1,
        '-': 1,
        '*': 2,
        '/': 2,
        '^': 3
    };
    
    const isOperator = (char: string): boolean => ['+', '-', '*', '/', '^'].includes(char);
    const isOperand = (char: string): boolean => /[a-zA-Z0-9]/.test(char);
    
    let postfix = '';
    
    for (const char of expression) {
        if (isOperand(char)) {
            postfix += char;
        } else if (char === '(') {
            stack.push(char);
        } else if (char === ')') {
            while (!stack.isEmpty() && stack.peek() !== '(') {
                postfix += stack.pop();
            }
            stack.pop(); // Remove the '('
        } else if (isOperator(char)) {
            while (!stack.isEmpty() && 
                   stack.peek() !== '(' && 
                   precedence[stack.peek()] >= precedence[char]) {
                postfix += stack.pop();
            }
            stack.push(char);
        }
    }
    
    while (!stack.isEmpty()) {
        postfix += stack.pop();
    }
    
    return postfix;
}

// Evaluate postfix expression
function evaluatePostfix(expression: string): number {
    const stack = new Stack<number>();
    const tokens: string[] = expression.split(' ').filter(token => token.trim() !== '');
    
    for (const token of tokens) {
        if (['+', '-', '*', '/'].includes(token)) {
            if (stack.size() < 2) {
                throw new Error("Invalid expression");
            }
            
            const operand2: number = stack.pop();
            const operand1: number = stack.pop();
            
            let result: number;
            switch (token) {
                case '+':
                    result = operand1 + operand2;
                    break;
                case '-':
                    result = operand1 - operand2;
                    break;
                case '*':
                    result = operand1 * operand2;
                    break;
                case '/':
                    if (operand2 === 0) {
                        throw new Error("Division by zero");
                    }
                    result = operand1 / operand2;
                    break;
                default:
                    throw new Error(`Unknown operator: ${token}`);
            }
            
            stack.push(result);
        } else {
            // Convert to number and push
            const num: number = parseFloat(token);
            if (!isNaN(num)) {
                stack.push(num);
            }
        }
    }
    
    if (stack.size() !== 1) {
        throw new Error("Invalid expression");
    }
    
    return stack.pop();
}

// Browser history simulation
class BrowserHistory {
    private backStack: Stack<string>;
    private forwardStack: Stack<string>;
    private currentPage: string | null;
    
    constructor() {
        this.backStack = new Stack<string>();
        this.forwardStack = new Stack<string>();
        this.currentPage = null;
    }
    
    public visit(page: string): void {
        if (this.currentPage) {
            this.backStack.push(this.currentPage);
        }
        this.currentPage = page;
        this.forwardStack.clear(); // Clear forward history
        console.log(`Visited: ${page}`);
    }
    
    public back(): void {
        if (this.backStack.isEmpty()) {
            console.log("No previous page");
            return;
        }
        
        if (this.currentPage) {
            this.forwardStack.push(this.currentPage);
        }
        this.currentPage = this.backStack.pop();
        console.log(`Back to: ${this.currentPage}`);
    }
    
    public forward(): void {
        if (this.forwardStack.isEmpty()) {
            console.log("No next page");
            return;
        }
        
        if (this.currentPage) {
            this.backStack.push(this.currentPage);
        }
        this.currentPage = this.forwardStack.pop();
        console.log(`Forward to: ${this.currentPage}`);
    }
    
    public getCurrentPage(): string | null {
        return this.currentPage;
    }
}

// Example usage
console.log("Balanced parentheses check:");
const expressions: string[] = ["(())", "([{}])", "(()", "([)]", "{[()]}"];
expressions.forEach(expr => {
    console.log(`${expr}: ${isBalanced(expr)}`);
});

console.log("\nString reversal:");
const original: string = "Hello, World!";
const reversed: string = reverseString(original);
console.log(`Original: ${original}`);
console.log(`Reversed: ${reversed}`);

console.log("\nInfix to Postfix conversion:");
const infix: string = "A+B*C-D";
const postfix: string = infixToPostfix(infix);
console.log(`Infix: ${infix}`);
console.log(`Postfix: ${postfix}`);

console.log("\nPostfix evaluation:");
const postfixExpr: string = "3 4 + 2 * 7 /";
const result: number = evaluatePostfix(postfixExpr);
console.log(`Expression: ${postfixExpr}`);
console.log(`Result: ${result}`);

console.log("\nBrowser history simulation:");
const browser = new BrowserHistory();
browser.visit("google.com");
browser.visit("stackoverflow.com");
browser.visit("github.com");
browser.back();
browser.back();
browser.forward();
browser.visit("reddit.com");
browser.back();
```

## Common Use Cases

- **Function calls**: Managing function call stack in programming languages
- **Expression evaluation**: Converting infix to postfix notation
- **Undo operations**: Implementing undo functionality in applications
- **Browser history**: Back/forward navigation
- **Balanced parentheses**: Checking syntax in compilers
- **Recursion**: Managing recursive function calls
- **Backtracking algorithms**: DFS, maze solving
- **Memory management**: Stack-based memory allocation

## Advantages

- Simple and intuitive LIFO behavior
- Efficient O(1) push and pop operations
- Memory efficient for temporary storage
- Natural fit for recursive algorithms
- Easy to implement and understand

## Disadvantages

- No random access to elements
- Limited access (only top element)
- Can lead to stack overflow with deep recursion
- Not suitable for searching or sorting operations