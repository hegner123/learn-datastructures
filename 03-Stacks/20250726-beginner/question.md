# Balanced Parentheses Checker - Beginner Level

## Problem Description

Implement a function that determines if parentheses, square brackets, and curly braces are properly balanced in a given string using a stack data structure. This is a classic application of stacks that demonstrates the LIFO (Last In, First Out) principle.

A string is considered balanced if:
1. Every opening bracket has a corresponding closing bracket
2. Brackets are closed in the correct order
3. No closing bracket appears without a matching opening bracket

## Requirements

Implement an `IsBalanced` function (Go) or `isBalanced` function (TypeScript) that:

1. **Input**: Takes a string containing only parentheses `()`, square brackets `[]`, and curly braces `{}`
2. **Output**: Returns `true` if the brackets are balanced, `false` otherwise
3. **Method**: Must use a stack data structure to solve the problem
4. **Edge Cases**: Handle empty strings (should return `true`)

## Example Usage

### Go Example
```go
fmt.Println(IsBalanced("()"))      // Output: true
fmt.Println(IsBalanced("([{}])"))  // Output: true
fmt.Println(IsBalanced("([)]"))    // Output: false
fmt.Println(IsBalanced("((())"))   // Output: false
fmt.Println(IsBalanced(""))        // Output: true
```

### TypeScript Example
```typescript
console.log(isBalanced("()"));      // Output: true
console.log(isBalanced("([{}])"));  // Output: true
console.log(isBalanced("([)]"));    // Output: false
console.log(isBalanced("((())"));   // Output: false
console.log(isBalanced(""));        // Output: true
```

## Algorithm Approach

Use a stack-based approach:

1. **Initialize** an empty stack
2. **Iterate** through each character in the string:
   - If it's an opening bracket `(`, `[`, or `{`, push it onto the stack
   - If it's a closing bracket `)`, `]`, or `}`:
     - Check if the stack is empty (unmatched closing bracket)
     - Pop the top element and verify it matches the current closing bracket
     - If they don't match, return false
3. **After processing** all characters, the stack should be empty for a balanced string

## Constraints

- The input string contains only parentheses `()`, square brackets `[]`, and curly braces `{}`
- No other characters will be present in the input
- The string length can be from 0 to 10,000 characters
- Time complexity should be O(n) where n is the length of the string
- Space complexity should be O(n) in the worst case (all opening brackets)

## Test Cases

Your implementation will be tested against these 10 test cases:

### Test Case 1: Empty String
- **Input**: `""`
- **Expected**: `true`
- **Reason**: Empty string is considered balanced

### Test Case 2: Single Pair of Parentheses
- **Input**: `"()"`
- **Expected**: `true`
- **Reason**: One properly matched pair

### Test Case 3: Single Pair of Brackets
- **Input**: `"[]"`
- **Expected**: `true`
- **Reason**: One properly matched pair of square brackets

### Test Case 4: Single Pair of Braces
- **Input**: `"{}"`
- **Expected**: `true`
- **Reason**: One properly matched pair of curly braces

### Test Case 5: Nested Brackets of Same Type
- **Input**: `"((()))"`
- **Expected**: `true`
- **Reason**: Properly nested parentheses

### Test Case 6: Mixed Balanced Brackets
- **Input**: `"([{}])"`
- **Expected**: `true`
- **Reason**: Different types of brackets properly nested

### Test Case 7: Unmatched Opening Bracket
- **Input**: `"(()"`
- **Expected**: `false`
- **Reason**: Missing closing parenthesis

### Test Case 8: Unmatched Closing Bracket
- **Input**: `"())"`
- **Expected**: `false`
- **Reason**: Extra closing parenthesis

### Test Case 9: Wrong Order of Brackets
- **Input**: `"([)]"`
- **Expected**: `false`
- **Reason**: Brackets are not properly nested (square bracket closes before parenthesis)

### Test Case 10: Complex Nested Structure
- **Input**: `"{[(){}][(())]}"`
- **Expected**: `true`
- **Reason**: Complex but properly balanced structure

## Implementation Notes

### Stack Operations Needed
- **Push**: Add opening brackets to the stack
- **Pop**: Remove and return the top element when encountering closing brackets
- **IsEmpty**: Check if stack is empty before popping
- **Size/Length**: Optional, for optimization

### Bracket Matching Logic
- `(` matches with `)`
- `[` matches with `]`
- `{` matches with `}`

### Helper Functions (Optional)
You may implement helper functions like:
- `isOpeningBracket(char)`: Check if character is `(`, `[`, or `{`
- `isClosingBracket(char)`: Check if character is `)`, `]`, or `}`
- `isMatchingPair(opening, closing)`: Verify if two brackets form a valid pair

## Getting Started

1. **Go**: Navigate to the `go/` directory and run `make test` to see the failing tests
2. **TypeScript**: Navigate to the `ts/` directory, run `npm install` (if not done), then `npm test` to see the failing tests

Complete the TODO sections in the respective `solution.go` and `solution.ts` files to make all tests pass.

## Learning Objectives

This problem helps you understand:
- **Stack Data Structure**: LIFO principle in action
- **Pattern Matching**: Using stacks for bracket/parentheses matching
- **Algorithm Design**: Breaking down a problem into stack operations
- **Edge Case Handling**: Empty strings and unmatched brackets
- **Real-world Application**: Similar logic is used in compilers and code editors