# Reference Implementation Problem

## Problem Description

Implement a simple array-based calculator that can perform basic operations (add, subtract, multiply, divide) on integers.

## Requirements

- Create a `Calculator` class/struct that can store numbers and perform operations
- Implement methods for: `add`, `subtract`, `multiply`, `divide`
- Handle division by zero appropriately
- Return results as integers (truncate division results)

## Example Usage

```
calc = new Calculator()
calc.add(5, 3) → 8
calc.subtract(10, 4) → 6
calc.multiply(7, 2) → 14
calc.divide(15, 3) → 5
calc.divide(10, 0) → error or special handling
```

## Constraints

- All inputs will be integers
- Division results should be truncated to integers
- Handle edge cases appropriately

## Test Cases

1. **Basic Addition**: add(5, 3) → 8
2. **Basic Subtraction**: subtract(10, 4) → 6
3. **Basic Multiplication**: multiply(7, 2) → 14
4. **Basic Division**: divide(15, 3) → 5
5. **Division with Remainder**: divide(10, 3) → 3
6. **Division by Zero**: divide(10, 0) → error/exception
7. **Negative Numbers**: add(-5, 3) → -2
8. **Zero Operations**: multiply(0, 5) → 0
9. **Large Numbers**: add(1000, 2000) → 3000
10. **Subtraction Result Negative**: subtract(3, 5) → -2