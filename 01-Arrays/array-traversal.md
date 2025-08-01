# Array Traversal Methods

## What is Array Traversal?

Array traversal refers to the process of visiting each element in an array in a systematic way. Different traversal methods provide various ways to access, process, or iterate through array elements, each with specific use cases and performance characteristics.

## Types of Array Traversal

### Basic Traversal Methods
1. Index-Based Traversal (Forward)
2. Index-Based Traversal (Backward)
3. Range-Based Traversal
4. Iterator-Based Traversal
5. Functional Traversal Methods
6. Multi-dimensional Array Traversal
7. Conditional Traversal
8. Partial Traversal

### Advanced Traversal Methods
9. Two-Pointer Traversal
10. Recursive Traversal
11. Sliding Window Traversal
12. Binary Search Traversal
13. Divide and Conquer Traversal
14. Backtracking Traversal

## Go Code Examples

### 1. Index-Based Forward Traversal

```go
package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    // Basic forward traversal
    fmt.Println("Forward traversal:")
    for i := 0; i < len(arr); i++ {
        fmt.Printf("Index %d: %d\n", i, arr[i])
    }
    
    // Processing elements during traversal
    sum := 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    fmt.Printf("Sum: %d\n", sum)
}
```

### 2. Index-Based Backward Traversal

```go
package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    // Backward traversal
    fmt.Println("Backward traversal:")
    for i := len(arr) - 1; i >= 0; i-- {
        fmt.Printf("Index %d: %d\n", i, arr[i])
    }
    
    // Reversing array elements
    reversed := [5]int{}
    for i := len(arr) - 1; i >= 0; i-- {
        reversed[len(arr)-1-i] = arr[i]
    }
    fmt.Printf("Reversed: %v\n", reversed)
}
```

### 3. Range-Based Traversal

```go
package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    // Range with both index and value
    fmt.Println("Range with index and value:")
    for index, value := range arr {
        fmt.Printf("Index %d: %d\n", index, value)
    }
    
    // Range with only values (ignoring index)
    fmt.Println("Range with only values:")
    for _, value := range arr {
        fmt.Printf("Value: %d\n", value)
    }
    
    // Range with only indices (ignoring values)
    fmt.Println("Range with only indices:")
    for index := range arr {
        fmt.Printf("Index: %d\n", index)
    }
}
```

### 4. Slice-Based Traversal

```go
package main

import "fmt"

func main() {
    slice := []int{10, 20, 30, 40, 50, 60, 70}
    
    // Traversing entire slice
    fmt.Println("Full slice traversal:")
    for i, v := range slice {
        fmt.Printf("Index %d: %d\n", i, v)
    }
    
    // Traversing slice segments
    fmt.Println("First 3 elements:")
    for i, v := range slice[:3] {
        fmt.Printf("Index %d: %d\n", i, v)
    }
    
    fmt.Println("Last 3 elements:")
    for i, v := range slice[len(slice)-3:] {
        fmt.Printf("Index %d: %d\n", i, v)
    }
    
    fmt.Println("Middle elements (index 2-4):")
    for i, v := range slice[2:5] {
        fmt.Printf("Index %d: %d\n", i+2, v)
    }
}
```

### 5. Functional-Style Traversal

```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Map operation (transform each element)
    squares := make([]int, len(arr))
    for i, v := range arr {
        squares[i] = v * v
    }
    fmt.Printf("Squares: %v\n", squares)
    
    // Filter operation (select elements based on condition)
    var evens []int
    for _, v := range arr {
        if v%2 == 0 {
            evens = append(evens, v)
        }
    }
    fmt.Printf("Even numbers: %v\n", evens)
    
    // Reduce operation (accumulate values)
    sum := 0
    for _, v := range arr {
        sum += v
    }
    fmt.Printf("Sum: %d\n", sum)
    
    // Find operation (locate first match)
    target := 7
    found := false
    foundIndex := -1
    for i, v := range arr {
        if v == target {
            found = true
            foundIndex = i
            break
        }
    }
    if found {
        fmt.Printf("Found %d at index %d\n", target, foundIndex)
    }
}
```

### 6. Multi-dimensional Array Traversal

```go
package main

import "fmt"

func main() {
    // 2D array traversal
    matrix := [3][4]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    
    // Row-major traversal (row by row)
    fmt.Println("Row-major traversal:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%2d ", matrix[i][j])
        }
        fmt.Println()
    }
    
    // Column-major traversal (column by column)
    fmt.Println("Column-major traversal:")
    for j := 0; j < len(matrix[0]); j++ {
        for i := 0; i < len(matrix); i++ {
            fmt.Printf("%2d ", matrix[i][j])
        }
        fmt.Println()
    }
    
    // Range-based 2D traversal
    fmt.Println("Range-based 2D traversal:")
    for i, row := range matrix {
        for j, value := range row {
            fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
        }
    }
    
    // Diagonal traversal
    fmt.Println("Main diagonal:")
    for i := 0; i < len(matrix) && i < len(matrix[0]); i++ {
        fmt.Printf("%d ", matrix[i][i])
    }
    fmt.Println()
    
    // 3D array traversal
    cube := [2][2][3]int{
        {
            {1, 2, 3},
            {4, 5, 6},
        },
        {
            {7, 8, 9},
            {10, 11, 12},
        },
    }
    
    fmt.Println("3D array traversal:")
    for i := range cube {
        fmt.Printf("Layer %d:\n", i)
        for j := range cube[i] {
            for k := range cube[i][j] {
                fmt.Printf("%2d ", cube[i][j][k])
            }
            fmt.Println()
        }
        fmt.Println()
    }
}
```

### 7. Conditional and Controlled Traversal

```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Early termination (break)
    fmt.Println("Find first number > 5:")
    for i, v := range arr {
        if v > 5 {
            fmt.Printf("Found %d at index %d\n", v, i)
            break
        }
    }
    
    // Skip elements (continue)
    fmt.Println("Print only odd numbers:")
    for _, v := range arr {
        if v%2 == 0 {
            continue
        }
        fmt.Printf("%d ", v)
    }
    fmt.Println()
    
    // Step traversal (every nth element)
    fmt.Println("Every 2nd element:")
    for i := 0; i < len(arr); i += 2 {
        fmt.Printf("Index %d: %d\n", i, arr[i])
    }
    
    // Traversal with bounds checking
    fmt.Println("Safe traversal with bounds checking:")
    indices := []int{0, 2, 5, 15, -1} // Some invalid indices
    for _, idx := range indices {
        if idx >= 0 && idx < len(arr) {
            fmt.Printf("arr[%d] = %d\n", idx, arr[idx])
        } else {
            fmt.Printf("Index %d is out of bounds\n", idx)
        }
    }
}
```

### 8. Parallel and Concurrent Traversal

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Concurrent processing with goroutines
    var wg sync.WaitGroup
    results := make([]int, len(arr))
    
    fmt.Println("Concurrent square calculation:")
    for i, v := range arr {
        wg.Add(1)
        go func(index, value int) {
            defer wg.Done()
            results[index] = value * value
            fmt.Printf("Goroutine %d: %d^2 = %d\n", index, value, value*value)
        }(i, v)
    }
    
    wg.Wait()
    fmt.Printf("Results: %v\n", results)
    
    // Channel-based traversal
    ch := make(chan int, len(arr))
    
    // Send array elements to channel
    go func() {
        for _, v := range arr {
            ch <- v
        }
        close(ch)
    }()
    
    // Receive and process from channel
    fmt.Println("Channel-based traversal:")
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
}
```

## TypeScript Code Examples

### 1. Index-Based Forward Traversal

```typescript
// Basic forward traversal
const arr: number[] = [10, 20, 30, 40, 50];

console.log("Forward traversal:");
for (let i: number = 0; i < arr.length; i++) {
    console.log(`Index ${i}: ${arr[i]}`);
}

// Processing elements during traversal
let sum: number = 0;
for (let i: number = 0; i < arr.length; i++) {
    sum += arr[i];
}
console.log(`Sum: ${sum}`);

// Type-safe traversal with generic function
function forwardTraversal<T>(array: T[], callback: (element: T, index: number) => void): void {
    for (let i: number = 0; i < array.length; i++) {
        callback(array[i], i);
    }
}

forwardTraversal(arr, (element: number, index: number) => {
    console.log(`Element at ${index}: ${element}`);
});
```

### 2. Index-Based Backward Traversal

```typescript
const arr: number[] = [10, 20, 30, 40, 50];

// Backward traversal
console.log("Backward traversal:");
for (let i: number = arr.length - 1; i >= 0; i--) {
    console.log(`Index ${i}: ${arr[i]}`);
}

// Reversing array elements
const reversed: number[] = [];
for (let i: number = arr.length - 1; i >= 0; i--) {
    reversed.push(arr[i]);
}
console.log(`Reversed: ${reversed}`);

// Generic backward traversal function
function backwardTraversal<T>(array: T[], callback: (element: T, index: number) => void): void {
    for (let i: number = array.length - 1; i >= 0; i--) {
        callback(array[i], i);
    }
}

backwardTraversal(arr, (element: number, index: number) => {
    console.log(`Backward - Index ${index}: ${element}`);
});
```

### 3. For...of and For...in Traversal

```typescript
const arr: number[] = [10, 20, 30, 40, 50];

// For...of loop (iterates over values)
console.log("For...of traversal (values):");
for (const value of arr) {
    console.log(`Value: ${value}`);
}

// For...in loop (iterates over indices)
console.log("For...in traversal (indices):");
for (const index in arr) {
    console.log(`Index ${index}: ${arr[index]}`);
}

// For...of with entries() for both index and value
console.log("For...of with entries:");
for (const [index, value] of arr.entries()) {
    console.log(`Index ${index}: ${value}`);
}

// Type-safe for...of with interface
interface Person {
    name: string;
    age: number;
}

const people: Person[] = [
    { name: "Alice", age: 25 },
    { name: "Bob", age: 30 },
    { name: "Charlie", age: 35 }
];

console.log("Traversing object array:");
for (const person of people) {
    console.log(`${person.name} is ${person.age} years old`);
}
```

### 4. Array Method-Based Traversal

```typescript
const arr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// forEach method
console.log("forEach traversal:");
arr.forEach((value: number, index: number, array: number[]) => {
    console.log(`Index ${index}: ${value}`);
});

// map method (transform elements)
console.log("Map transformation:");
const squares: number[] = arr.map((value: number, index: number) => {
    console.log(`Squaring ${value}`);
    return value * value;
});
console.log(`Squares: ${squares}`);

// filter method (conditional traversal)
console.log("Filter traversal:");
const evens: number[] = arr.filter((value: number, index: number) => {
    console.log(`Checking ${value}`);
    return value % 2 === 0;
});
console.log(`Even numbers: ${evens}`);

// reduce method (accumulative traversal)
console.log("Reduce traversal:");
const sum: number = arr.reduce((accumulator: number, value: number, index: number) => {
    console.log(`Adding ${value} to ${accumulator}`);
    return accumulator + value;
}, 0);
console.log(`Sum: ${sum}`);

// find method (find first match)
console.log("Find traversal:");
const found: number | undefined = arr.find((value: number, index: number) => {
    console.log(`Checking ${value}`);
    return value > 7;
});
console.log(`First number > 7: ${found}`);

// findIndex method
const foundIndex: number = arr.findIndex((value: number) => value > 7);
console.log(`Index of first number > 7: ${foundIndex}`);
```

### 5. Functional-Style Advanced Traversal

```typescript
const arr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Chaining multiple array methods
console.log("Method chaining:");
const result: number = arr
    .filter((value: number) => {
        console.log(`Filtering ${value}`);
        return value % 2 === 0;
    })
    .map((value: number) => {
        console.log(`Mapping ${value}`);
        return value * 2;
    })
    .reduce((sum: number, value: number) => {
        console.log(`Reducing ${value}`);
        return sum + value;
    }, 0);
console.log(`Chained result: ${result}`);

// Custom traversal functions
function customTraversal<T, R>(
    array: T[], 
    transform: (item: T, index: number) => R,
    filter?: (item: T, index: number) => boolean
): R[] {
    const results: R[] = [];
    
    for (let i: number = 0; i < array.length; i++) {
        if (!filter || filter(array[i], i)) {
            results.push(transform(array[i], i));
        }
    }
    
    return results;
}

const customResult: string[] = customTraversal(
    arr,
    (value: number, index: number) => `Item ${index}: ${value}`,
    (value: number) => value > 5
);
console.log("Custom traversal result:", customResult);

// Async traversal with Promise.all
async function asyncTraversal<T, R>(
    array: T[],
    asyncCallback: (item: T, index: number) => Promise<R>
): Promise<R[]> {
    const promises: Promise<R>[] = array.map((item: T, index: number) => 
        asyncCallback(item, index)
    );
    return Promise.all(promises);
}

// Example usage (simulated async operation)
async function processAsync(value: number, index: number): Promise<string> {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(`Processed ${value} at index ${index}`);
        }, 100);
    });
}

asyncTraversal(arr.slice(0, 3), processAsync)
    .then((results: string[]) => {
        console.log("Async traversal results:", results);
    });
```

### 6. Multi-dimensional Array Traversal

```typescript
// 2D array traversal
const matrix: number[][] = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12]
];

// Row-major traversal
console.log("Row-major traversal:");
for (let i: number = 0; i < matrix.length; i++) {
    let row: string = "";
    for (let j: number = 0; j < matrix[i].length; j++) {
        row += matrix[i][j].toString().padStart(3);
    }
    console.log(row);
}

// Column-major traversal
console.log("Column-major traversal:");
const cols: number = matrix[0].length;
for (let j: number = 0; j < cols; j++) {
    let column: string = "";
    for (let i: number = 0; i < matrix.length; i++) {
        column += matrix[i][j].toString().padStart(3);
    }
    console.log(column);
}

// forEach-based 2D traversal
console.log("forEach-based 2D traversal:");
matrix.forEach((row: number[], rowIndex: number) => {
    row.forEach((value: number, colIndex: number) => {
        console.log(`matrix[${rowIndex}][${colIndex}] = ${value}`);
    });
});

// Flat traversal (convert 2D to 1D)
console.log("Flattened traversal:");
const flattened: number[] = matrix.flat();
flattened.forEach((value: number, index: number) => {
    console.log(`Flat[${index}]: ${value}`);
});

// Type-safe matrix operations
interface Matrix<T> {
    data: T[][];
    rows: number;
    cols: number;
}

function createMatrix<T>(data: T[][]): Matrix<T> {
    return {
        data,
        rows: data.length,
        cols: data[0]?.length || 0
    };
}

function traverseMatrix<T>(
    matrix: Matrix<T>, 
    callback: (value: T, row: number, col: number) => void
): void {
    for (let i: number = 0; i < matrix.rows; i++) {
        for (let j: number = 0; j < matrix.cols; j++) {
            callback(matrix.data[i][j], i, j);
        }
    }
}

const typedMatrix: Matrix<number> = createMatrix(matrix);
console.log("Type-safe matrix traversal:");
traverseMatrix(typedMatrix, (value: number, row: number, col: number) => {
    console.log(`matrix[${row}][${col}] = ${value}`);
});

// 3D array traversal
const cube: number[][][] = [
    [
        [1, 2, 3],
        [4, 5, 6]
    ],
    [
        [7, 8, 9],
        [10, 11, 12]
    ]
];

console.log("3D array traversal:");
cube.forEach((layer: number[][], layerIndex: number) => {
    console.log(`Layer ${layerIndex}:`);
    layer.forEach((row: number[], rowIndex: number) => {
        row.forEach((value: number, colIndex: number) => {
            console.log(`  cube[${layerIndex}][${rowIndex}][${colIndex}] = ${value}`);
        });
    });
});
```

### 7. Conditional and Controlled Traversal

```typescript
const arr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Early termination with some/every
console.log("Early termination with some:");
const hasNumberGreaterThan5: boolean = arr.some((value: number, index: number) => {
    console.log(`Checking ${value}`);
    return value > 5;
});
console.log(`Has number > 5: ${hasNumberGreaterThan5}`);

console.log("Check all with every:");
const allPositive: boolean = arr.every((value: number, index: number) => {
    console.log(`Checking ${value}`);
    return value > 0;
});
console.log(`All positive: ${allPositive}`);

// Step traversal
console.log("Every 2nd element:");
for (let i: number = 0; i < arr.length; i += 2) {
    console.log(`Index ${i}: ${arr[i]}`);
}

// Conditional traversal with filter
console.log("Conditional traversal:");
const conditions = [
    (x: number) => x % 2 === 0,
    (x: number) => x > 5,
    (x: number) => x < 8
];

conditions.forEach((condition, condIndex) => {
    console.log(`Condition ${condIndex + 1}:`);
    arr.filter(condition).forEach((value: number, index: number) => {
        console.log(`  Matched: ${value}`);
    });
});

// Safe traversal with bounds checking
console.log("Safe traversal:");
const indices: number[] = [0, 2, 5, 15, -1];
indices.forEach((idx: number) => {
    if (idx >= 0 && idx < arr.length) {
        console.log(`arr[${idx}] = ${arr[idx]}`);
    } else {
        console.log(`Index ${idx} is out of bounds`);
    }
});

// Custom iterator
class ArrayIterator<T> {
    private array: T[];
    private position: number = 0;

    constructor(array: T[]) {
        this.array = array;
    }

    hasNext(): boolean {
        return this.position < this.array.length;
    }

    next(): T | undefined {
        if (this.hasNext()) {
            return this.array[this.position++];
        }
        return undefined;
    }

    reset(): void {
        this.position = 0;
    }

    skip(count: number): void {
        this.position = Math.min(this.position + count, this.array.length);
    }
}

console.log("Custom iterator traversal:");
const iterator = new ArrayIterator(arr);
while (iterator.hasNext()) {
    const value = iterator.next();
    console.log(`Iterator value: ${value}`);
    
    // Skip next element if current is even
    if (value && value % 2 === 0) {
        iterator.skip(1);
    }
}
```

### 8. Async and Generator-Based Traversal

```typescript
const arr: number[] = [1, 2, 3, 4, 5];

// Generator function for traversal
function* arrayGenerator<T>(array: T[]): Generator<T, void, unknown> {
    for (const item of array) {
        console.log(`Yielding: ${item}`);
        yield item;
    }
}

console.log("Generator traversal:");
const gen = arrayGenerator(arr);
for (const value of gen) {
    console.log(`Received: ${value}`);
}

// Async generator
async function* asyncArrayGenerator<T>(array: T[]): AsyncGenerator<T, void, unknown> {
    for (const item of array) {
        await new Promise(resolve => setTimeout(resolve, 100));
        console.log(`Async yielding: ${item}`);
        yield item;
    }
}

console.log("Async generator traversal:");
(async () => {
    for await (const value of asyncArrayGenerator(arr)) {
        console.log(`Async received: ${value}`);
    }
})();

// Sequential async processing
async function sequentialAsyncTraversal<T, R>(
    array: T[],
    asyncCallback: (item: T, index: number) => Promise<R>
): Promise<R[]> {
    const results: R[] = [];
    
    for (let i = 0; i < array.length; i++) {
        console.log(`Processing item ${i}`);
        const result = await asyncCallback(array[i], i);
        results.push(result);
    }
    
    return results;
}

// Parallel async processing
async function parallelAsyncTraversal<T, R>(
    array: T[],
    asyncCallback: (item: T, index: number) => Promise<R>
): Promise<R[]> {
    console.log("Starting parallel processing");
    const promises = array.map((item, index) => asyncCallback(item, index));
    return Promise.all(promises);
}

// Example async processing
async function processItem(value: number, index: number): Promise<string> {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(`Processed ${value} at index ${index}`);
        }, Math.random() * 1000);
    });
}

// Sequential processing
sequentialAsyncTraversal(arr.slice(0, 3), processItem)
    .then((results: string[]) => {
        console.log("Sequential results:", results);
    });

// Parallel processing
parallelAsyncTraversal(arr.slice(0, 3), processItem)
    .then((results: string[]) => {
        console.log("Parallel results:", results);
    });

// Observable-like pattern
class ArrayObservable<T> {
    private array: T[];

    constructor(array: T[]) {
        this.array = array;
    }

    subscribe(observer: {
        next: (value: T, index: number) => void;
        complete: () => void;
        error?: (error: Error) => void;
    }): void {
        try {
            this.array.forEach((value, index) => {
                observer.next(value, index);
            });
            observer.complete();
        } catch (error) {
            if (observer.error) {
                observer.error(error as Error);
            }
        }
    }
}

console.log("Observable-like traversal:");
const observable = new ArrayObservable(arr);
observable.subscribe({
    next: (value: number, index: number) => {
        console.log(`Observable - Index ${index}: ${value}`);
    },
    complete: () => {
        console.log("Observable traversal completed");
    },
    error: (error: Error) => {
        console.error("Error during traversal:", error);
    }
});
```

## Advanced Traversal Methods

### 9. Two-Pointer Traversal

#### Go Examples

```go
package main

import "fmt"

func main() {
    // Two-pointer technique for finding pairs
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    target := 10
    
    fmt.Println("Two-pointer pair finding:")
    findPairsWithSum(arr, target)
    
    // Two arrays traversal
    arr1 := []int{1, 3, 5, 7, 9}
    arr2 := []int{2, 4, 6, 8, 10}
    fmt.Println("Merging two sorted arrays:")
    merged := mergeTwoSortedArrays(arr1, arr2)
    fmt.Printf("Merged: %v\n", merged)
    
    // Opposite direction pointers
    fmt.Println("Palindrome check with two pointers:")
    testStrings := []string{"racecar", "hello", "madam"}
    for _, s := range testStrings {
        fmt.Printf("%s is palindrome: %v\n", s, isPalindrome(s))
    }
    
    // Fast and slow pointers (Floyd's algorithm concept)
    fmt.Println("Finding middle element with fast/slow pointers:")
    findMiddleElement(arr)
}

func findPairsWithSum(arr []int, target int) {
    left, right := 0, len(arr)-1
    
    for left < right {
        sum := arr[left] + arr[right]
        if sum == target {
            fmt.Printf("Pair found: %d + %d = %d\n", arr[left], arr[right], target)
            left++
            right--
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
}

func mergeTwoSortedArrays(arr1, arr2 []int) []int {
    result := make([]int, 0, len(arr1)+len(arr2))
    i, j := 0, 0
    
    // Two-pointer traversal of both arrays
    for i < len(arr1) && j < len(arr2) {
        if arr1[i] <= arr2[j] {
            result = append(result, arr1[i])
            i++
        } else {
            result = append(result, arr2[j])
            j++
        }
    }
    
    // Add remaining elements
    for i < len(arr1) {
        result = append(result, arr1[i])
        i++
    }
    
    for j < len(arr2) {
        result = append(result, arr2[j])
        j++
    }
    
    return result
}

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func findMiddleElement(arr []int) {
    if len(arr) == 0 {
        return
    }
    
    slow, fast := 0, 0
    
    // Fast pointer moves 2 steps, slow moves 1 step
    for fast < len(arr)-1 && fast+1 < len(arr) {
        slow++
        fast += 2
    }
    
    fmt.Printf("Middle element: %d at index %d\n", arr[slow], slow)
}

// Three-pointer technique for triplet finding
func findTripletWithSum(arr []int, target int) {
    if len(arr) < 3 {
        return
    }
    
    fmt.Printf("Finding triplets with sum %d:\n", target)
    
    for i := 0; i < len(arr)-2; i++ {
        left, right := i+1, len(arr)-1
        
        for left < right {
            sum := arr[i] + arr[left] + arr[right]
            if sum == target {
                fmt.Printf("Triplet: %d + %d + %d = %d\n", arr[i], arr[left], arr[right], target)
                left++
                right--
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }
}

// Container with most water problem
func maxWaterContainer(heights []int) int {
    if len(heights) < 2 {
        return 0
    }
    
    left, right := 0, len(heights)-1
    maxArea := 0
    
    for left < right {
        width := right - left
        height := min(heights[left], heights[right])
        area := width * height
        
        if area > maxArea {
            maxArea = area
        }
        
        // Move the pointer with smaller height
        if heights[left] < heights[right] {
            left++
        } else {
            right--
        }
    }
    
    fmt.Printf("Maximum water container area: %d\n", maxArea)
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

#### TypeScript Examples

```typescript
// Two-pointer traversal techniques

// Finding pairs with target sum
function findPairsWithSum(arr: number[], target: number): [number, number][] {
    const pairs: [number, number][] = [];
    let left: number = 0;
    let right: number = arr.length - 1;
    
    console.log(`Finding pairs with sum ${target}:`);
    
    while (left < right) {
        const sum: number = arr[left] + arr[right];
        
        if (sum === target) {
            pairs.push([arr[left], arr[right]]);
            console.log(`Pair found: ${arr[left]} + ${arr[right]} = ${target}`);
            left++;
            right--;
        } else if (sum < target) {
            left++;
        } else {
            right--;
        }
    }
    
    return pairs;
}

// Merging two sorted arrays
function mergeTwoSortedArrays<T>(arr1: T[], arr2: T[]): T[] {
    const result: T[] = [];
    let i: number = 0;
    let j: number = 0;
    
    // Two-pointer traversal
    while (i < arr1.length && j < arr2.length) {
        if (arr1[i] <= arr2[j]) {
            result.push(arr1[i]);
            i++;
        } else {
            result.push(arr2[j]);
            j++;
        }
    }
    
    // Add remaining elements
    while (i < arr1.length) {
        result.push(arr1[i]);
        i++;
    }
    
    while (j < arr2.length) {
        result.push(arr2[j]);
        j++;
    }
    
    return result;
}

// Palindrome check with two pointers
function isPalindrome(s: string): boolean {
    let left: number = 0;
    let right: number = s.length - 1;
    
    while (left < right) {
        if (s[left] !== s[right]) {
            return false;
        }
        left++;
        right--;
    }
    
    return true;
}

// Fast and slow pointers for finding middle
function findMiddleElement<T>(arr: T[]): T | undefined {
    if (arr.length === 0) return undefined;
    
    let slow: number = 0;
    let fast: number = 0;
    
    while (fast < arr.length - 1 && fast + 1 < arr.length) {
        slow++;
        fast += 2;
    }
    
    console.log(`Middle element: ${arr[slow]} at index ${slow}`);
    return arr[slow];
}

// Three-pointer technique for triplets
function findTripletsWithSum(arr: number[], target: number): number[][] {
    const triplets: number[][] = [];
    
    if (arr.length < 3) return triplets;
    
    console.log(`Finding triplets with sum ${target}:`);
    
    for (let i = 0; i < arr.length - 2; i++) {
        let left: number = i + 1;
        let right: number = arr.length - 1;
        
        while (left < right) {
            const sum: number = arr[i] + arr[left] + arr[right];
            
            if (sum === target) {
                triplets.push([arr[i], arr[left], arr[right]]);
                console.log(`Triplet: ${arr[i]} + ${arr[left]} + ${arr[right]} = ${target}`);
                left++;
                right--;
            } else if (sum < target) {
                left++;
            } else {
                right--;
            }
        }
    }
    
    return triplets;
}

// Container with most water
function maxWaterContainer(heights: number[]): number {
    if (heights.length < 2) return 0;
    
    let left: number = 0;
    let right: number = heights.length - 1;
    let maxArea: number = 0;
    
    while (left < right) {
        const width: number = right - left;
        const height: number = Math.min(heights[left], heights[right]);
        const area: number = width * height;
        
        maxArea = Math.max(maxArea, area);
        
        // Move pointer with smaller height
        if (heights[left] < heights[right]) {
            left++;
        } else {
            right--;
        }
    }
    
    console.log(`Maximum water container area: ${maxArea}`);
    return maxArea;
}

// Remove duplicates from sorted array (in-place with two pointers)
function removeDuplicates(arr: number[]): number {
    if (arr.length <= 1) return arr.length;
    
    let writeIndex: number = 1;
    
    for (let readIndex: number = 1; readIndex < arr.length; readIndex++) {
        if (arr[readIndex] !== arr[readIndex - 1]) {
            arr[writeIndex] = arr[readIndex];
            writeIndex++;
        }
    }
    
    console.log(`Array after removing duplicates: ${arr.slice(0, writeIndex)}`);
    return writeIndex;
}

// Dutch National Flag problem (three pointers)
function dutchNationalFlag(arr: number[]): void {
    let low: number = 0;
    let mid: number = 0;
    let high: number = arr.length - 1;
    
    console.log(`Original array: ${arr}`);
    
    while (mid <= high) {
        switch (arr[mid]) {
            case 0:
                [arr[low], arr[mid]] = [arr[mid], arr[low]];
                low++;
                mid++;
                break;
            case 1:
                mid++;
                break;
            case 2:
                [arr[mid], arr[high]] = [arr[high], arr[mid]];
                high--;
                break;
        }
    }
    
    console.log(`Sorted array (0s, 1s, 2s): ${arr}`);
}

// Example usage
const sortedArr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9];
findPairsWithSum(sortedArr, 10);

const arr1: number[] = [1, 3, 5, 7];
const arr2: number[] = [2, 4, 6, 8];
const merged = mergeTwoSortedArrays(arr1, arr2);
console.log("Merged arrays:", merged);

console.log("racecar is palindrome:", isPalindrome("racecar"));
console.log("hello is palindrome:", isPalindrome("hello"));

findMiddleElement(sortedArr);
findTripletsWithSum([1, 2, 3, 4, 5, 6], 10);

const heights: number[] = [1, 8, 6, 2, 5, 4, 8, 3, 7];
maxWaterContainer(heights);

const duplicateArr: number[] = [1, 1, 2, 2, 3, 4, 4, 5];
removeDuplicates(duplicateArr);

const flagArr: number[] = [2, 0, 1, 2, 1, 0, 1, 2, 0];
dutchNationalFlag(flagArr);
```

### 10. Recursive Traversal

#### Go Examples

```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    
    fmt.Println("Recursive traversal examples:")
    
    // Forward recursive traversal
    fmt.Println("Forward recursive traversal:")
    traverseForward(arr, 0)
    
    // Backward recursive traversal
    fmt.Println("Backward recursive traversal:")
    traverseBackward(arr, len(arr)-1)
    
    // Recursive sum
    sum := recursiveSum(arr, 0)
    fmt.Printf("Recursive sum: %d\n", sum)
    
    // Binary search (recursive)
    sortedArr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    target := 7
    index := binarySearchRecursive(sortedArr, target, 0, len(sortedArr)-1)
    fmt.Printf("Binary search for %d: index %d\n", target, index)
    
    // Recursive array reversal
    fmt.Printf("Original array: %v\n", arr)
    reverseRecursive(arr, 0, len(arr)-1)
    fmt.Printf("Reversed array: %v\n", arr)
    
    // Recursive palindrome check
    palindromes := [][]int{{1, 2, 3, 2, 1}, {1, 2, 3, 4, 5}}
    for _, p := range palindromes {
        fmt.Printf("Array %v is palindrome: %v\n", p, isPalindromeRecursive(p, 0, len(p)-1))
    }
    
    // Recursive permutation generation
    fmt.Println("All permutations:")
    generatePermutations([]int{1, 2, 3}, 0)
}

func traverseForward(arr []int, index int) {
    if index >= len(arr) {
        return
    }
    
    fmt.Printf("Index %d: %d\n", index, arr[index])
    traverseForward(arr, index+1)
}

func traverseBackward(arr []int, index int) {
    if index < 0 {
        return
    }
    
    fmt.Printf("Index %d: %d\n", index, arr[index])
    traverseBackward(arr, index-1)
}

func recursiveSum(arr []int, index int) int {
    if index >= len(arr) {
        return 0
    }
    
    return arr[index] + recursiveSum(arr, index+1)
}

func binarySearchRecursive(arr []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    mid := left + (right-left)/2
    
    if arr[mid] == target {
        return mid
    } else if arr[mid] > target {
        return binarySearchRecursive(arr, target, left, mid-1)
    } else {
        return binarySearchRecursive(arr, target, mid+1, right)
    }
}

func reverseRecursive(arr []int, left, right int) {
    if left >= right {
        return
    }
    
    arr[left], arr[right] = arr[right], arr[left]
    reverseRecursive(arr, left+1, right-1)
}

func isPalindromeRecursive(arr []int, left, right int) bool {
    if left >= right {
        return true
    }
    
    if arr[left] != arr[right] {
        return false
    }
    
    return isPalindromeRecursive(arr, left+1, right-1)
}

func generatePermutations(arr []int, start int) {
    if start == len(arr) {
        fmt.Printf("Permutation: %v\n", arr)
        return
    }
    
    for i := start; i < len(arr); i++ {
        arr[start], arr[i] = arr[i], arr[start]
        generatePermutations(arr, start+1)
        arr[start], arr[i] = arr[i], arr[start] // backtrack
    }
}

// Recursive merge sort
func mergeSort(arr []int, left, right int) {
    if left >= right {
        return
    }
    
    mid := left + (right-left)/2
    
    mergeSort(arr, left, mid)
    mergeSort(arr, mid+1, right)
    merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
    leftArr := make([]int, mid-left+1)
    rightArr := make([]int, right-mid)
    
    copy(leftArr, arr[left:mid+1])
    copy(rightArr, arr[mid+1:right+1])
    
    i, j, k := 0, 0, left
    
    for i < len(leftArr) && j < len(rightArr) {
        if leftArr[i] <= rightArr[j] {
            arr[k] = leftArr[i]
            i++
        } else {
            arr[k] = rightArr[j]
            j++
        }
        k++
    }
    
    for i < len(leftArr) {
        arr[k] = leftArr[i]
        i++
        k++
    }
    
    for j < len(rightArr) {
        arr[k] = rightArr[j]
        j++
        k++
    }
}

// Recursive quick sort
func quickSort(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)
        
        quickSort(arr, low, pi-1)
        quickSort(arr, pi+1, high)
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}
```

#### TypeScript Examples

```typescript
// Recursive traversal methods

// Forward recursive traversal
function traverseForward<T>(arr: T[], index: number = 0): void {
    if (index >= arr.length) {
        return;
    }
    
    console.log(`Index ${index}: ${arr[index]}`);
    traverseForward(arr, index + 1);
}

// Backward recursive traversal
function traverseBackward<T>(arr: T[], index: number = arr.length - 1): void {
    if (index < 0) {
        return;
    }
    
    console.log(`Index ${index}: ${arr[index]}`);
    traverseBackward(arr, index - 1);
}

// Recursive sum calculation
function recursiveSum(arr: number[], index: number = 0): number {
    if (index >= arr.length) {
        return 0;
    }
    
    return arr[index] + recursiveSum(arr, index + 1);
}

// Recursive binary search
function binarySearchRecursive(
    arr: number[], 
    target: number, 
    left: number = 0, 
    right: number = arr.length - 1
): number {
    if (left > right) {
        return -1;
    }
    
    const mid: number = Math.floor((left + right) / 2);
    
    if (arr[mid] === target) {
        return mid;
    } else if (arr[mid] > target) {
        return binarySearchRecursive(arr, target, left, mid - 1);
    } else {
        return binarySearchRecursive(arr, target, mid + 1, right);
    }
}

// Recursive array reversal
function reverseRecursive<T>(arr: T[], left: number = 0, right: number = arr.length - 1): void {
    if (left >= right) {
        return;
    }
    
    [arr[left], arr[right]] = [arr[right], arr[left]];
    reverseRecursive(arr, left + 1, right - 1);
}

// Recursive palindrome check
function isPalindromeRecursive<T>(arr: T[], left: number = 0, right: number = arr.length - 1): boolean {
    if (left >= right) {
        return true;
    }
    
    if (arr[left] !== arr[right]) {
        return false;
    }
    
    return isPalindromeRecursive(arr, left + 1, right - 1);
}

// Recursive permutation generation
function generatePermutations<T>(arr: T[], start: number = 0, result: T[][] = []): T[][] {
    if (start === arr.length) {
        result.push([...arr]); // Create a copy
        return result;
    }
    
    for (let i = start; i < arr.length; i++) {
        [arr[start], arr[i]] = [arr[i], arr[start]]; // swap
        generatePermutations(arr, start + 1, result);
        [arr[start], arr[i]] = [arr[i], arr[start]]; // backtrack
    }
    
    return result;
}

// Recursive combination generation
function generateCombinations<T>(
    arr: T[], 
    k: number, 
    start: number = 0, 
    current: T[] = [], 
    result: T[][] = []
): T[][] {
    if (current.length === k) {
        result.push([...current]);
        return result;
    }
    
    for (let i = start; i < arr.length; i++) {
        current.push(arr[i]);
        generateCombinations(arr, k, i + 1, current, result);
        current.pop(); // backtrack
    }
    
    return result;
}

// Recursive merge sort
function mergeSort(arr: number[]): number[] {
    if (arr.length <= 1) {
        return arr;
    }
    
    const mid: number = Math.floor(arr.length / 2);
    const left: number[] = mergeSort(arr.slice(0, mid));
    const right: number[] = mergeSort(arr.slice(mid));
    
    return merge(left, right);
}

function merge(left: number[], right: number[]): number[] {
    const result: number[] = [];
    let i: number = 0;
    let j: number = 0;
    
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            result.push(left[i]);
            i++;
        } else {
            result.push(right[j]);
            j++;
        }
    }
    
    return result.concat(left.slice(i)).concat(right.slice(j));
}

// Recursive quick sort
function quickSort(arr: number[], low: number = 0, high: number = arr.length - 1): number[] {
    const sortedArr = [...arr]; // Create copy to avoid mutation
    
    function quickSortHelper(arr: number[], low: number, high: number): void {
        if (low < high) {
            const pi: number = partition(arr, low, high);
            
            quickSortHelper(arr, low, pi - 1);
            quickSortHelper(arr, pi + 1, high);
        }
    }
    
    quickSortHelper(sortedArr, low, high);
    return sortedArr;
}

function partition(arr: number[], low: number, high: number): number {
    const pivot: number = arr[high];
    let i: number = low - 1;
    
    for (let j = low; j < high; j++) {
        if (arr[j] < pivot) {
            i++;
            [arr[i], arr[j]] = [arr[j], arr[i]];
        }
    }
    
    [arr[i + 1], arr[high]] = [arr[high], arr[i + 1]];
    return i + 1;
}

// Recursive tree traversal simulation with arrays
interface TreeNode {
    value: number;
    left?: TreeNode;
    right?: TreeNode;
}

function arrayToTree(arr: (number | null)[], index: number = 0): TreeNode | null {
    if (index >= arr.length || arr[index] === null) {
        return null;
    }
    
    const node: TreeNode = { value: arr[index] as number };
    node.left = arrayToTree(arr, 2 * index + 1);
    node.right = arrayToTree(arr, 2 * index + 2);
    
    return node;
}

function preorderTraversal(node: TreeNode | null, result: number[] = []): number[] {
    if (node === null) {
        return result;
    }
    
    result.push(node.value);
    preorderTraversal(node.left, result);
    preorderTraversal(node.right, result);
    
    return result;
}

function inorderTraversal(node: TreeNode | null, result: number[] = []): number[] {
    if (node === null) {
        return result;
    }
    
    inorderTraversal(node.left, result);
    result.push(node.value);
    inorderTraversal(node.right, result);
    
    return result;
}

function postorderTraversal(node: TreeNode | null, result: number[] = []): number[] {
    if (node === null) {
        return result;
    }
    
    postorderTraversal(node.left, result);
    postorderTraversal(node.right, result);
    result.push(node.value);
    
    return result;
}

// Example usage
const arr: number[] = [1, 2, 3, 4, 5];

console.log("Forward recursive traversal:");
traverseForward(arr);

console.log("Backward recursive traversal:");
traverseBackward(arr);

console.log("Recursive sum:", recursiveSum(arr));

const sortedArr: number[] = [1, 3, 5, 7, 9, 11, 13, 15];
console.log("Binary search for 7:", binarySearchRecursive(sortedArr, 7));

const testArr: number[] = [1, 2, 3, 4, 5];
console.log("Original array:", testArr);
reverseRecursive(testArr);
console.log("Reversed array:", testArr);

const palindromes: number[][] = [[1, 2, 3, 2, 1], [1, 2, 3, 4, 5]];
palindromes.forEach(p => {
    console.log(`Array ${p} is palindrome:`, isPalindromeRecursive(p));
});

console.log("Permutations of [1, 2, 3]:");
const permutations = generatePermutations([1, 2, 3]);
permutations.forEach(p => console.log(p));

console.log("Combinations of [1, 2, 3, 4] choose 2:");
const combinations = generateCombinations([1, 2, 3, 4], 2);
combinations.forEach(c => console.log(c));

const unsortedArr: number[] = [64, 34, 25, 12, 22, 11, 90];
console.log("Original array:", unsortedArr);
console.log("Merge sorted:", mergeSort(unsortedArr));
console.log("Quick sorted:", quickSort(unsortedArr));

// Tree traversal example
const treeArray: (number | null)[] = [1, 2, 3, 4, 5, null, 6];
const tree = arrayToTree(treeArray);
console.log("Preorder traversal:", preorderTraversal(tree));
console.log("Inorder traversal:", inorderTraversal(tree));
console.log("Postorder traversal:", postorderTraversal(tree));
```

### 11. Sliding Window Traversal

#### Go Examples

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    fmt.Println("Sliding window traversal examples:")
    
    // Fixed-size sliding window
    windowSize := 3
    fmt.Printf("Fixed window size %d - max values:\n", windowSize)
    maxInWindow(arr, windowSize)
    
    // Variable-size sliding window
    targetSum := 15
    fmt.Printf("Subarray with sum %d:\n", targetSum)
    findSubarrayWithSum(arr, targetSum)
    
    // Longest substring without repeating characters (simulated with array)
    str := "abcabcbb"
    fmt.Printf("Longest substring without repeating chars in '%s': %d\n", 
               str, longestSubstringWithoutRepeating(str))
    
    // Minimum window substring
    s := "ADOBECODEBANC"
    t := "ABC"
    fmt.Printf("Minimum window substring of '%s' containing '%s': '%s'\n", 
               s, t, minWindowSubstring(s, t))
}

func maxInWindow(arr []int, k int) {
    if len(arr) < k {
        return
    }
    
    // Sliding window maximum
    for i := 0; i <= len(arr)-k; i++ {
        max := arr[i]
        for j := i; j < i+k; j++ {
            if arr[j] > max {
                max = arr[j]
            }
        }
        fmt.Printf("Window [%d:%d] = %v, Max = %d\n", 
                   i, i+k-1, arr[i:i+k], max)
    }
}

func findSubarrayWithSum(arr []int, targetSum int) bool {
    left, right := 0, 0
    currentSum := 0
    
    for right < len(arr) {
        currentSum += arr[right]
        
        // Shrink window if sum exceeds target
        for currentSum > targetSum && left <= right {
            currentSum -= arr[left]
            left++
        }
        
        // Check if we found the target sum
        if currentSum == targetSum {
            fmt.Printf("Subarray found: %v (indices %d to %d)\n", 
                       arr[left:right+1], left, right)
            return true
        }
        
        right++
    }
    
    fmt.Println("No subarray found with the target sum")
    return false
}

func longestSubstringWithoutRepeating(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    charMap := make(map[byte]int)
    left, maxLength := 0, 0
    
    for right := 0; right < len(s); right++ {
        if lastIndex, exists := charMap[s[right]]; exists && lastIndex >= left {
            left = lastIndex + 1
        }
        
        charMap[s[right]] = right
        currentLength := right - left + 1
        
        if currentLength > maxLength {
            maxLength = currentLength
            fmt.Printf("Current longest substring: '%s' (length %d)\n", 
                       s[left:right+1], maxLength)
        }
    }
    
    return maxLength
}

func minWindowSubstring(s, t string) string {
    if len(s) < len(t) {
        return ""
    }
    
    // Count characters in t
    tCount := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        tCount[t[i]]++
    }
    
    required := len(tCount)
    formed := 0
    windowCounts := make(map[byte]int)
    
    left, right := 0, 0
    minLen := math.MaxInt32
    minLeft := 0
    
    for right < len(s) {
        char := s[right]
        windowCounts[char]++
        
        if count, exists := tCount[char]; exists && windowCounts[char] == count {
            formed++
        }
        
        // Contract window when possible
        for left <= right && formed == required {
            if right-left+1 < minLen {
                minLen = right - left + 1
                minLeft = left
            }
            
            char = s[left]
            windowCounts[char]--
            if count, exists := tCount[char]; exists && windowCounts[char] < count {
                formed--
            }
            left++
        }
        
        right++
    }
    
    if minLen == math.MaxInt32 {
        return ""
    }
    
    return s[minLeft : minLeft+minLen]
}

// Sliding window with deque for maximum in each window
func maxSlidingWindow(arr []int, k int) []int {
    if len(arr) == 0 || k == 0 {
        return []int{}
    }
    
    result := make([]int, 0, len(arr)-k+1)
    deque := make([]int, 0) // indices
    
    for i := 0; i < len(arr); i++ {
        // Remove indices outside current window
        for len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]
        }
        
        // Remove indices with smaller values from back
        for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
            deque = deque[:len(deque)-1]
        }
        
        deque = append(deque, i)
        
        // Add to result when window is complete
        if i >= k-1 {
            result = append(result, arr[deque[0]])
        }
    }
    
    return result
}
```

#### TypeScript Examples

```typescript
// Sliding window traversal patterns

// Fixed-size sliding window - maximum in each window
function maxInSlidingWindow(arr: number[], k: number): number[] {
    if (arr.length < k) return [];
    
    const result: number[] = [];
    
    console.log(`Fixed window size ${k} - finding maximums:`);
    
    for (let i = 0; i <= arr.length - k; i++) {
        const window = arr.slice(i, i + k);
        const max = Math.max(...window);
        
        console.log(`Window [${i}:${i + k - 1}] = [${window.join(', ')}], Max = ${max}`);
        result.push(max);
    }
    
    return result;
}

// Variable-size sliding window - subarray with target sum
function findSubarrayWithSum(arr: number[], targetSum: number): number[] | null {
    let left: number = 0;
    let currentSum: number = 0;
    
    for (let right = 0; right < arr.length; right++) {
        currentSum += arr[right];
        
        // Shrink window if sum exceeds target
        while (currentSum > targetSum && left <= right) {
            currentSum -= arr[left];
            left++;
        }
        
        // Check if we found the target sum
        if (currentSum === targetSum) {
            const subarray = arr.slice(left, right + 1);
            console.log(`Subarray found: [${subarray.join(', ')}] (indices ${left} to ${right})`);
            return subarray;
        }
    }
    
    console.log("No subarray found with the target sum");
    return null;
}

// Longest substring without repeating characters
function longestSubstringWithoutRepeating(s: string): number {
    if (s.length === 0) return 0;
    
    const charMap = new Map<string, number>();
    let left: number = 0;
    let maxLength: number = 0;
    let maxSubstring: string = "";
    
    for (let right = 0; right < s.length; right++) {
        const char = s[right];
        
        if (charMap.has(char) && charMap.get(char)! >= left) {
            left = charMap.get(char)! + 1;
        }
        
        charMap.set(char, right);
        const currentLength = right - left + 1;
        
        if (currentLength > maxLength) {
            maxLength = currentLength;
            maxSubstring = s.substring(left, right + 1);
            console.log(`Current longest substring: '${maxSubstring}' (length ${maxLength})`);
        }
    }
    
    return maxLength;
}

// Minimum window substring
function minWindowSubstring(s: string, t: string): string {
    if (s.length < t.length) return "";
    
    // Count characters in t
    const tCount = new Map<string, number>();
    for (const char of t) {
        tCount.set(char, (tCount.get(char) || 0) + 1);
    }
    
    const required = tCount.size;
    let formed = 0;
    const windowCounts = new Map<string, number>();
    
    let left = 0;
    let minLen = Infinity;
    let minLeft = 0;
    
    for (let right = 0; right < s.length; right++) {
        const char = s[right];
        windowCounts.set(char, (windowCounts.get(char) || 0) + 1);
        
        if (tCount.has(char) && windowCounts.get(char) === tCount.get(char)) {
            formed++;
        }
        
        // Contract window when possible
        while (left <= right && formed === required) {
            if (right - left + 1 < minLen) {
                minLen = right - left + 1;
                minLeft = left;
            }
            
            const leftChar = s[left];
            windowCounts.set(leftChar, windowCounts.get(leftChar)! - 1);
            
            if (tCount.has(leftChar) && windowCounts.get(leftChar)! < tCount.get(leftChar)!) {
                formed--;
            }
            
            left++;
        }
    }
    
    return minLen === Infinity ? "" : s.substring(minLeft, minLeft + minLen);
}

// Sliding window with deque for maximum efficiency
class Deque<T> {
    private items: T[] = [];
    
    pushFront(item: T): void {
        this.items.unshift(item);
    }
    
    pushBack(item: T): void {
        this.items.push(item);
    }
    
    popFront(): T | undefined {
        return this.items.shift();
    }
    
    popBack(): T | undefined {
        return this.items.pop();
    }
    
    front(): T | undefined {
        return this.items[0];
    }
    
    back(): T | undefined {
        return this.items[this.items.length - 1];
    }
    
    isEmpty(): boolean {
        return this.items.length === 0;
    }
    
    size(): number {
        return this.items.length;
    }
}

function maxSlidingWindowOptimized(arr: number[], k: number): number[] {
    if (arr.length === 0 || k === 0) return [];
    
    const result: number[] = [];
    const deque = new Deque<number>(); // Store indices
    
    for (let i = 0; i < arr.length; i++) {
        // Remove indices outside current window
        while (!deque.isEmpty() && deque.front()! <= i - k) {
            deque.popFront();
        }
        
        // Remove indices with smaller values from back
        while (!deque.isEmpty() && arr[deque.back()!] <= arr[i]) {
            deque.popBack();
        }
        
        deque.pushBack(i);
        
        // Add to result when window is complete
        if (i >= k - 1) {
            result.push(arr[deque.front()!]);
        }
    }
    
    return result;
}

// Longest subarray with at most K distinct elements
function longestSubarrayWithKDistinct(arr: number[], k: number): number[] {
    if (k === 0) return [];
    
    const count = new Map<number, number>();
    let left = 0;
    let maxLength = 0;
    let maxStart = 0;
    
    for (let right = 0; right < arr.length; right++) {
        count.set(arr[right], (count.get(arr[right]) || 0) + 1);
        
        // Shrink window if we have more than k distinct elements
        while (count.size > k) {
            const leftElement = arr[left];
            count.set(leftElement, count.get(leftElement)! - 1);
            
            if (count.get(leftElement) === 0) {
                count.delete(leftElement);
            }
            
            left++;
        }
        
        // Update maximum length
        if (right - left + 1 > maxLength) {
            maxLength = right - left + 1;
            maxStart = left;
        }
    }
    
    const result = arr.slice(maxStart, maxStart + maxLength);
    console.log(`Longest subarray with at most ${k} distinct elements: [${result.join(', ')}]`);
    return result;
}

// Sliding window average
function slidingWindowAverage(arr: number[], k: number): number[] {
    if (arr.length < k) return [];
    
    const averages: number[] = [];
    let windowSum = 0;
    
    // Calculate sum of first window
    for (let i = 0; i < k; i++) {
        windowSum += arr[i];
    }
    averages.push(windowSum / k);
    
    // Slide the window
    for (let i = k; i < arr.length; i++) {
        windowSum = windowSum - arr[i - k] + arr[i];
        averages.push(windowSum / k);
    }
    
    console.log(`Sliding window averages (window size ${k}):`, averages);
    return averages;
}

// Example usage
const arr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9];

console.log("Sliding window examples:");
maxInSlidingWindow(arr, 3);

findSubarrayWithSum(arr, 15);

console.log("Longest substring without repeating:", 
    longestSubstringWithoutRepeating("abcabcbb"));

console.log("Minimum window substring:", 
    minWindowSubstring("ADOBECODEBANC", "ABC"));

console.log("Optimized max sliding window:", 
    maxSlidingWindowOptimized([1, 3, -1, -3, 5, 3, 6, 7], 3));

longestSubarrayWithKDistinct([1, 2, 1, 2, 3], 2);

slidingWindowAverage([1, 3, 2, 6, -1, 4, 1, 8, 2], 5);
```

### 12. Binary Search Traversal

#### Go Examples

```go
package main

import "fmt"

func main() {
    sortedArr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    
    fmt.Println("Binary search traversal examples:")
    
    // Standard binary search
    target := 7
    index := binarySearch(sortedArr, target)
    fmt.Printf("Binary search for %d: index %d\n", target, index)
    
    // Find first occurrence
    duplicateArr := []int{1, 2, 2, 2, 3, 4, 5, 5, 5, 6}
    target = 2
    firstIndex := findFirstOccurrence(duplicateArr, target)
    fmt.Printf("First occurrence of %d: index %d\n", target, firstIndex)
    
    // Find last occurrence
    lastIndex := findLastOccurrence(duplicateArr, target)
    fmt.Printf("Last occurrence of %d: index %d\n", target, lastIndex)
    
    // Search in rotated sorted array
    rotatedArr := []int{4, 5, 6, 7, 0, 1, 2}
    target = 0
    rotatedIndex := searchRotatedArray(rotatedArr, target)
    fmt.Printf("Search %d in rotated array: index %d\n", target, rotatedIndex)
    
    // Find peak element
    peakArr := []int{1, 2, 3, 1}
    peakIndex := findPeakElement(peakArr)
    fmt.Printf("Peak element in %v: index %d (value %d)\n", peakArr, peakIndex, peakArr[peakIndex])
    
    // Square root using binary search
    num := 25
    sqrt := mySqrt(num)
    fmt.Printf("Square root of %d: %d\n", num, sqrt)
}

func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

func findFirstOccurrence(arr []int, target int) int {
    left, right := 0, len(arr)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            result = mid
            right = mid - 1 // Continue searching in left half
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}

func findLastOccurrence(arr []int, target int) int {
    left, right := 0, len(arr)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            result = mid
            left = mid + 1 // Continue searching in right half
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}

func searchRotatedArray(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            return mid
        }
        
        // Check if left half is sorted
        if arr[left] <= arr[mid] {
            if target >= arr[left] && target < arr[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // Right half is sorted
            if target > arr[mid] && target <= arr[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}

func findPeakElement(arr []int) int {
    if len(arr) == 1 {
        return 0
    }
    
    left, right := 0, len(arr)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] > arr[mid+1] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}

func mySqrt(x int) int {
    if x < 2 {
        return x
    }
    
    left, right := 2, x/2
    
    for left <= right {
        mid := left + (right-left)/2
        num := mid * mid
        
        if num == x {
            return mid
        } else if num < x {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return right
}

// Search for insertion position
func searchInsert(arr []int, target int) int {
    left, right := 0, len(arr)
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}

// Find minimum in rotated sorted array
func findMin(arr []int) int {
    left, right := 0, len(arr)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] > arr[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return arr[left]
}

// Two sum in sorted array using binary search
func twoSumBinarySearch(arr []int, target int) []int {
    for i := 0; i < len(arr)-1; i++ {
        complement := target - arr[i]
        j := binarySearch(arr[i+1:], complement)
        if j != -1 {
            return []int{i, i + j + 1}
        }
    }
    return []int{-1, -1}
}
```

#### TypeScript Examples

```typescript
// Binary search traversal methods

// Standard binary search
function binarySearch(arr: number[], target: number): number {
    let left: number = 0;
    let right: number = arr.length - 1;
    
    while (left <= right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] === target) {
            return mid;
        } else if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return -1;
}

// Find first occurrence of target
function findFirstOccurrence(arr: number[], target: number): number {
    let left: number = 0;
    let right: number = arr.length - 1;
    let result: number = -1;
    
    while (left <= right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] === target) {
            result = mid;
            right = mid - 1; // Continue searching in left half
        } else if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return result;
}

// Find last occurrence of target
function findLastOccurrence(arr: number[], target: number): number {
    let left: number = 0;
    let right: number = arr.length - 1;
    let result: number = -1;
    
    while (left <= right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] === target) {
            result = mid;
            left = mid + 1; // Continue searching in right half
        } else if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return result;
}

// Search in rotated sorted array
function searchRotatedArray(arr: number[], target: number): number {
    let left: number = 0;
    let right: number = arr.length - 1;
    
    while (left <= right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] === target) {
            return mid;
        }
        
        // Check if left half is sorted
        if (arr[left] <= arr[mid]) {
            if (target >= arr[left] && target < arr[mid]) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        } else { // Right half is sorted
            if (target > arr[mid] && target <= arr[right]) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
    }
    
    return -1;
}

// Find peak element
function findPeakElement(arr: number[]): number {
    if (arr.length === 1) return 0;
    
    let left: number = 0;
    let right: number = arr.length - 1;
    
    while (left < right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] > arr[mid + 1]) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    
    return left;
}

// Square root using binary search
function mySqrt(x: number): number {
    if (x < 2) return x;
    
    let left: number = 2;
    let right: number = Math.floor(x / 2);
    
    while (left <= right) {
        const mid: number = Math.floor((left + right) / 2);
        const square: number = mid * mid;
        
        if (square === x) {
            return mid;
        } else if (square < x) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return right;
}

// Search for insertion position
function searchInsert(arr: number[], target: number): number {
    let left: number = 0;
    let right: number = arr.length;
    
    while (left < right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid;
        }
    }
    
    return left;
}

// Find minimum in rotated sorted array
function findMin(arr: number[]): number {
    let left: number = 0;
    let right: number = arr.length - 1;
    
    while (left < right) {
        const mid: number = Math.floor((left + right) / 2);
        
        if (arr[mid] > arr[right]) {
            left = mid + 1;
        } else {
            right = mid;
        }
    }
    
    return arr[left];
}

// Find range of target (first and last occurrence)
function searchRange(arr: number[], target: number): [number, number] {
    const first = findFirstOccurrence(arr, target);
    if (first === -1) {
        return [-1, -1];
    }
    
    const last = findLastOccurrence(arr, target);
    return [first, last];
}

// Median of two sorted arrays using binary search
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    if (nums1.length > nums2.length) {
        [nums1, nums2] = [nums2, nums1];
    }
    
    const m = nums1.length;
    const n = nums2.length;
    let left = 0;
    let right = m;
    
    while (left <= right) {
        const partitionX = Math.floor((left + right) / 2);
        const partitionY = Math.floor((m + n + 1) / 2) - partitionX;
        
        const maxLeftX = partitionX === 0 ? -Infinity : nums1[partitionX - 1];
        const minRightX = partitionX === m ? Infinity : nums1[partitionX];
        
        const maxLeftY = partitionY === 0 ? -Infinity : nums2[partitionY - 1];
        const minRightY = partitionY === n ? Infinity : nums2[partitionY];
        
        if (maxLeftX <= minRightY && maxLeftY <= minRightX) {
            if ((m + n) % 2 === 0) {
                return (Math.max(maxLeftX, maxLeftY) + Math.min(minRightX, minRightY)) / 2;
            } else {
                return Math.max(maxLeftX, maxLeftY);
            }
        } else if (maxLeftX > minRightY) {
            right = partitionX - 1;
        } else {
            left = partitionX + 1;
        }
    }
    
    throw new Error("Input arrays are not sorted");
}

// Binary search on answer (find minimum capacity)
function shipWithinDays(weights: number[], days: number): number {
    let left = Math.max(...weights);
    let right = weights.reduce((sum, w) => sum + w, 0);
    
    function canShip(capacity: number): boolean {
        let daysNeeded = 1;
        let currentWeight = 0;
        
        for (const weight of weights) {
            if (currentWeight + weight > capacity) {
                daysNeeded++;
                currentWeight = weight;
            } else {
                currentWeight += weight;
            }
        }
        
        return daysNeeded <= days;
    }
    
    while (left < right) {
        const mid = Math.floor((left + right) / 2);
        
        if (canShip(mid)) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    
    return left;
}

// Kth smallest element in sorted matrix using binary search
function kthSmallest(matrix: number[][], k: number): number {
    const n = matrix.length;
    let left = matrix[0][0];
    let right = matrix[n - 1][n - 1];
    
    function countSmallerEqual(target: number): number {
        let count = 0;
        let row = n - 1;
        let col = 0;
        
        while (row >= 0 && col < n) {
            if (matrix[row][col] <= target) {
                count += row + 1;
                col++;
            } else {
                row--;
            }
        }
        
        return count;
    }
    
    while (left < right) {
        const mid = Math.floor((left + right) / 2);
        
        if (countSmallerEqual(mid) < k) {
            left = mid + 1;
        } else {
            right = mid;
        }
    }
    
    return left;
}

// Example usage
const sortedArr: number[] = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19];

console.log("Binary search examples:");
console.log("Binary search for 7:", binarySearch(sortedArr, 7));

const duplicateArr: number[] = [1, 2, 2, 2, 3, 4, 5, 5, 5, 6];
console.log("First occurrence of 2:", findFirstOccurrence(duplicateArr, 2));
console.log("Last occurrence of 2:", findLastOccurrence(duplicateArr, 2));
console.log("Range of 2:", searchRange(duplicateArr, 2));

const rotatedArr: number[] = [4, 5, 6, 7, 0, 1, 2];
console.log("Search 0 in rotated array:", searchRotatedArray(rotatedArr, 0));

const peakArr: number[] = [1, 2, 3, 1];
const peakIndex = findPeakElement(peakArr);
console.log(`Peak element: index ${peakIndex}, value ${peakArr[peakIndex]}`);

console.log("Square root of 25:", mySqrt(25));
console.log("Insert position for 5 in [1,3,5,6]:", searchInsert([1, 3, 5, 6], 5));

const rotatedMin: number[] = [3, 4, 5, 1, 2];
console.log("Minimum in rotated array:", findMin(rotatedMin));

console.log("Median of [1,3] and [2]:", findMedianSortedArrays([1, 3], [2]));

const weights: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
console.log("Ship within 5 days:", shipWithinDays(weights, 5));

const matrix: number[][] = [[1, 5, 9], [10, 11, 13], [12, 13, 15]];
console.log("3rd smallest in matrix:", kthSmallest(matrix, 3));
```

### 13. Divide and Conquer Traversal

#### Go Examples

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    arr := []int{3, 5, 1, 6, 2, 0, 8, 7, 9, 4}
    
    fmt.Println("Divide and conquer traversal examples:")
    
    // Merge sort
    fmt.Printf("Original array: %v\n", arr)
    sortedMerge := make([]int, len(arr))
    copy(sortedMerge, arr)
    mergeSort(sortedMerge, 0, len(sortedMerge)-1)
    fmt.Printf("Merge sorted: %v\n", sortedMerge)
    
    // Quick sort
    sortedQuick := make([]int, len(arr))
    copy(sortedQuick, arr)
    quickSort(sortedQuick, 0, len(sortedQuick)-1)
    fmt.Printf("Quick sorted: %v\n", sortedQuick)
    
    // Maximum subarray sum (Kadane's algorithm with divide and conquer)
    subarrayArr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    maxSum := maxSubarrayDivideConquer(subarrayArr, 0, len(subarrayArr)-1)
    fmt.Printf("Maximum subarray sum in %v: %d\n", subarrayArr, maxSum)
    
    // Count inversions
    inversionArr := []int{8, 4, 2, 1}
    inversions := countInversions(inversionArr, 0, len(inversionArr)-1)
    fmt.Printf("Inversions in %v: %d\n", inversionArr, inversions)
    
    // Matrix multiplication (Strassen's algorithm concept)
    matrixA := [][]int{{1, 2}, {3, 4}}
    matrixB := [][]int{{5, 6}, {7, 8}}
    result := matrixMultiply(matrixA, matrixB)
    fmt.Printf("Matrix multiplication result: %v\n", result)
    
    // Find majority element
    majorityArr := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
    majority := findMajorityElement(majorityArr, 0, len(majorityArr)-1)
    fmt.Printf("Majority element in %v: %d\n", majorityArr, majority)
}

func mergeSort(arr []int, left, right int) {
    if left >= right {
        return
    }
    
    mid := left + (right-left)/2
    
    // Divide
    mergeSort(arr, left, mid)
    mergeSort(arr, mid+1, right)
    
    // Conquer
    merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
    leftSize := mid - left + 1
    rightSize := right - mid
    
    leftArr := make([]int, leftSize)
    rightArr := make([]int, rightSize)
    
    // Copy data to temporary arrays
    copy(leftArr, arr[left:mid+1])
    copy(rightArr, arr[mid+1:right+1])
    
    // Merge the temporary arrays back
    i, j, k := 0, 0, left
    
    for i < leftSize && j < rightSize {
        if leftArr[i] <= rightArr[j] {
            arr[k] = leftArr[i]
            i++
        } else {
            arr[k] = rightArr[j]
            j++
        }
        k++
    }
    
    // Copy remaining elements
    for i < leftSize {
        arr[k] = leftArr[i]
        i++
        k++
    }
    
    for j < rightSize {
        arr[k] = rightArr[j]
        j++
        k++
    }
}

func quickSort(arr []int, low, high int) {
    if low < high {
        // Partition and get pivot index
        pi := partition(arr, low, high)
        
        // Recursively sort elements before and after partition
        quickSort(arr, low, pi-1)
        quickSort(arr, pi+1, high)
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[high] // Choose last element as pivot
    i := low - 1       // Index of smaller element
    
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}

func maxSubarrayDivideConquer(arr []int, left, right int) int {
    if left == right {
        return arr[left]
    }
    
    mid := left + (right-left)/2
    
    // Find maximum sum in left half
    leftSum := maxSubarrayDivideConquer(arr, left, mid)
    
    // Find maximum sum in right half
    rightSum := maxSubarrayDivideConquer(arr, mid+1, right)
    
    // Find maximum sum crossing the middle
    crossSum := maxCrossingSum(arr, left, mid, right)
    
    // Return maximum of the three
    return max(max(leftSum, rightSum), crossSum)
}

func maxCrossingSum(arr []int, left, mid, right int) int {
    // Find maximum sum for left side
    leftSum := math.MinInt32
    sum := 0
    for i := mid; i >= left; i-- {
        sum += arr[i]
        if sum > leftSum {
            leftSum = sum
        }
    }
    
    // Find maximum sum for right side
    rightSum := math.MinInt32
    sum = 0
    for i := mid + 1; i <= right; i++ {
        sum += arr[i]
        if sum > rightSum {
            rightSum = sum
        }
    }
    
    return leftSum + rightSum
}

func countInversions(arr []int, left, right int) int {
    if left >= right {
        return 0
    }
    
    mid := left + (right-left)/2
    
    // Count inversions in left half
    leftInversions := countInversions(arr, left, mid)
    
    // Count inversions in right half
    rightInversions := countInversions(arr, mid+1, right)
    
    // Count inversions between left and right halves
    crossInversions := mergeAndCount(arr, left, mid, right)
    
    return leftInversions + rightInversions + crossInversions
}

func mergeAndCount(arr []int, left, mid, right int) int {
    leftArr := make([]int, mid-left+1)
    rightArr := make([]int, right-mid)
    
    copy(leftArr, arr[left:mid+1])
    copy(rightArr, arr[mid+1:right+1])
    
    i, j, k := 0, 0, left
    invCount := 0
    
    for i < len(leftArr) && j < len(rightArr) {
        if leftArr[i] <= rightArr[j] {
            arr[k] = leftArr[i]
            i++
        } else {
            arr[k] = rightArr[j]
            invCount += len(leftArr) - i // All remaining elements in left are greater
            j++
        }
        k++
    }
    
    // Copy remaining elements
    for i < len(leftArr) {
        arr[k] = leftArr[i]
        i++
        k++
    }
    
    for j < len(rightArr) {
        arr[k] = rightArr[j]
        j++
        k++
    }
    
    return invCount
}

func matrixMultiply(A, B [][]int) [][]int {
    n := len(A)
    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, n)
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                result[i][j] += A[i][k] * B[k][j]
            }
        }
    }
    
    return result
}

func findMajorityElement(arr []int, left, right int) int {
    if left == right {
        return arr[left]
    }
    
    mid := left + (right-left)/2
    
    leftMajority := findMajorityElement(arr, left, mid)
    rightMajority := findMajorityElement(arr, mid+1, right)
    
    if leftMajority == rightMajority {
        return leftMajority
    }
    
    // Count occurrences of both candidates
    leftCount := countInRange(arr, leftMajority, left, right)
    rightCount := countInRange(arr, rightMajority, left, right)
    
    if leftCount > (right-left+1)/2 {
        return leftMajority
    } else if rightCount > (right-left+1)/2 {
        return rightMajority
    }
    
    return -1 // No majority element
}

func countInRange(arr []int, target, left, right int) int {
    count := 0
    for i := left; i <= right; i++ {
        if arr[i] == target {
            count++
        }
    }
    return count
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

#### TypeScript Examples

```typescript
// Divide and conquer traversal methods

// Merge sort implementation
function mergeSort(arr: number[]): number[] {
    if (arr.length <= 1) {
        return arr;
    }
    
    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid));
    const right = mergeSort(arr.slice(mid));
    
    return merge(left, right);
}

function merge(left: number[], right: number[]): number[] {
    const result: number[] = [];
    let i = 0;
    let j = 0;
    
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            result.push(left[i]);
            i++;
        } else {
            result.push(right[j]);
            j++;
        }
    }
    
    return result.concat(left.slice(i)).concat(right.slice(j));
}

// Quick sort implementation
function quickSort(arr: number[]): number[] {
    if (arr.length <= 1) {
        return arr;
    }
    
    const pivot = arr[Math.floor(arr.length / 2)];
    const left = arr.filter(x => x < pivot);
    const middle = arr.filter(x => x === pivot);
    const right = arr.filter(x => x > pivot);
    
    return [...quickSort(left), ...middle, ...quickSort(right)];
}

// Maximum subarray sum using divide and conquer
function maxSubarrayDivideConquer(arr: number[], left: number = 0, right: number = arr.length - 1): number {
    if (left === right) {
        return arr[left];
    }
    
    const mid = Math.floor((left + right) / 2);
    
    // Find maximum sum in left half
    const leftSum = maxSubarrayDivideConquer(arr, left, mid);
    
    // Find maximum sum in right half
    const rightSum = maxSubarrayDivideConquer(arr, mid + 1, right);
    
    // Find maximum sum crossing the middle
    const crossSum = maxCrossingSum(arr, left, mid, right);
    
    return Math.max(Math.max(leftSum, rightSum), crossSum);
}

function maxCrossingSum(arr: number[], left: number, mid: number, right: number): number {
    // Find maximum sum for left side
    let leftSum = -Infinity;
    let sum = 0;
    for (let i = mid; i >= left; i--) {
        sum += arr[i];
        leftSum = Math.max(leftSum, sum);
    }
    
    // Find maximum sum for right side
    let rightSum = -Infinity;
    sum = 0;
    for (let i = mid + 1; i <= right; i++) {
        sum += arr[i];
        rightSum = Math.max(rightSum, sum);
    }
    
    return leftSum + rightSum;
}

// Count inversions using divide and conquer
function countInversions(arr: number[]): number {
    const temp = [...arr];
    return mergeSortAndCount(temp, 0, arr.length - 1);
}

function mergeSortAndCount(arr: number[], left: number, right: number): number {
    if (left >= right) {
        return 0;
    }
    
    const mid = Math.floor((left + right) / 2);
    
    const leftInversions = mergeSortAndCount(arr, left, mid);
    const rightInversions = mergeSortAndCount(arr, mid + 1, right);
    const crossInversions = mergeAndCount(arr, left, mid, right);
    
    return leftInversions + rightInversions + crossInversions;
}

function mergeAndCount(arr: number[], left: number, mid: number, right: number): number {
    const leftArr = arr.slice(left, mid + 1);
    const rightArr = arr.slice(mid + 1, right + 1);
    
    let i = 0;
    let j = 0;
    let k = left;
    let invCount = 0;
    
    while (i < leftArr.length && j < rightArr.length) {
        if (leftArr[i] <= rightArr[j]) {
            arr[k] = leftArr[i];
            i++;
        } else {
            arr[k] = rightArr[j];
            invCount += leftArr.length - i; // All remaining elements in left are greater
            j++;
        }
        k++;
    }
    
    while (i < leftArr.length) {
        arr[k] = leftArr[i];
        i++;
        k++;
    }
    
    while (j < rightArr.length) {
        arr[k] = rightArr[j];
        j++;
        k++;
    }
    
    return invCount;
}

// Matrix multiplication using divide and conquer
function matrixMultiply(A: number[][], B: number[][]): number[][] {
    const n = A.length;
    const result: number[][] = Array(n).fill(null).map(() => Array(n).fill(0));
    
    if (n === 1) {
        result[0][0] = A[0][0] * B[0][0];
        return result;
    }
    
    // Simple multiplication for small matrices
    if (n <= 2) {
        for (let i = 0; i < n; i++) {
            for (let j = 0; j < n; j++) {
                for (let k = 0; k < n; k++) {
                    result[i][j] += A[i][k] * B[k][j];
                }
            }
        }
        return result;
    }
    
    // Divide matrices into quadrants and apply Strassen's algorithm concept
    const mid = Math.floor(n / 2);
    
    const A11 = getSubMatrix(A, 0, 0, mid);
    const A12 = getSubMatrix(A, 0, mid, mid);
    const A21 = getSubMatrix(A, mid, 0, mid);
    const A22 = getSubMatrix(A, mid, mid, mid);
    
    const B11 = getSubMatrix(B, 0, 0, mid);
    const B12 = getSubMatrix(B, 0, mid, mid);
    const B21 = getSubMatrix(B, mid, 0, mid);
    const B22 = getSubMatrix(B, mid, mid, mid);
    
    // Recursive multiplication
    const C11 = addMatrices(matrixMultiply(A11, B11), matrixMultiply(A12, B21));
    const C12 = addMatrices(matrixMultiply(A11, B12), matrixMultiply(A12, B22));
    const C21 = addMatrices(matrixMultiply(A21, B11), matrixMultiply(A22, B21));
    const C22 = addMatrices(matrixMultiply(A21, B12), matrixMultiply(A22, B22));
    
    // Combine results
    combineMatrices(result, C11, C12, C21, C22);
    
    return result;
}

function getSubMatrix(matrix: number[][], row: number, col: number, size: number): number[][] {
    const result: number[][] = Array(size).fill(null).map(() => Array(size).fill(0));
    
    for (let i = 0; i < size; i++) {
        for (let j = 0; j < size; j++) {
            result[i][j] = matrix[row + i][col + j];
        }
    }
    
    return result;
}

function addMatrices(A: number[][], B: number[][]): number[][] {
    const n = A.length;
    const result: number[][] = Array(n).fill(null).map(() => Array(n).fill(0));
    
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < n; j++) {
            result[i][j] = A[i][j] + B[i][j];
        }
    }
    
    return result;
}

function combineMatrices(result: number[][], C11: number[][], C12: number[][], C21: number[][], C22: number[][]): void {
    const mid = C11.length;
    
    for (let i = 0; i < mid; i++) {
        for (let j = 0; j < mid; j++) {
            result[i][j] = C11[i][j];
            result[i][j + mid] = C12[i][j];
            result[i + mid][j] = C21[i][j];
            result[i + mid][j + mid] = C22[i][j];
        }
    }
}

// Find majority element using divide and conquer
function findMajorityElement(arr: number[], left: number = 0, right: number = arr.length - 1): number | null {
    if (left === right) {
        return arr[left];
    }
    
    const mid = Math.floor((left + right) / 2);
    
    const leftMajority = findMajorityElement(arr, left, mid);
    const rightMajority = findMajorityElement(arr, mid + 1, right);
    
    if (leftMajority === rightMajority) {
        return leftMajority;
    }
    
    // Count occurrences of both candidates
    const leftCount = leftMajority !== null ? countInRange(arr, leftMajority, left, right) : 0;
    const rightCount = rightMajority !== null ? countInRange(arr, rightMajority, left, right) : 0;
    
    const requiredCount = Math.floor((right - left + 1) / 2) + 1;
    
    if (leftCount >= requiredCount) {
        return leftMajority;
    } else if (rightCount >= requiredCount) {
        return rightMajority;
    }
    
    return null; // No majority element
}

function countInRange(arr: number[], target: number, left: number, right: number): number {
    let count = 0;
    for (let i = left; i <= right; i++) {
        if (arr[i] === target) {
            count++;
        }
    }
    return count;
}

// Closest pair of points using divide and conquer
interface Point {
    x: number;
    y: number;
}

function closestPairDistance(points: Point[]): number {
    if (points.length <= 1) return Infinity;
    
    // Sort points by x-coordinate
    const sortedByX = [...points].sort((a, b) => a.x - b.x);
    const sortedByY = [...points].sort((a, b) => a.y - b.y);
    
    return closestPairRec(sortedByX, sortedByY);
}

function closestPairRec(Px: Point[], Py: Point[]): number {
    const n = Px.length;
    
    if (n <= 3) {
        return bruteForceClosest(Px);
    }
    
    const mid = Math.floor(n / 2);
    const midPoint = Px[mid];
    
    const Pyl = Py.filter(point => point.x <= midPoint.x);
    const Pyr = Py.filter(point => point.x > midPoint.x);
    
    const dl = closestPairRec(Px.slice(0, mid), Pyl);
    const dr = closestPairRec(Px.slice(mid), Pyr);
    
    const d = Math.min(dl, dr);
    
    // Find points close to the dividing line
    const strip = Py.filter(point => Math.abs(point.x - midPoint.x) < d);
    
    return Math.min(d, stripClosest(strip, d));
}

function bruteForceClosest(points: Point[]): number {
    let minDist = Infinity;
    
    for (let i = 0; i < points.length; i++) {
        for (let j = i + 1; j < points.length; j++) {
            const dist = distance(points[i], points[j]);
            minDist = Math.min(minDist, dist);
        }
    }
    
    return minDist;
}

function stripClosest(strip: Point[], d: number): number {
    let minDist = d;
    
    for (let i = 0; i < strip.length; i++) {
        for (let j = i + 1; j < strip.length && (strip[j].y - strip[i].y) < minDist; j++) {
            const dist = distance(strip[i], strip[j]);
            minDist = Math.min(minDist, dist);
        }
    }
    
    return minDist;
}

function distance(p1: Point, p2: Point): number {
    return Math.sqrt((p1.x - p2.x) ** 2 + (p1.y - p2.y) ** 2);
}

// Example usage
console.log("Divide and conquer examples:");

const arr: number[] = [3, 5, 1, 6, 2, 0, 8, 7, 9, 4];
console.log("Original array:", arr);
console.log("Merge sorted:", mergeSort(arr));
console.log("Quick sorted:", quickSort(arr));

const subarrayArr: number[] = [-2, 1, -3, 4, -1, 2, 1, -5, 4];
console.log("Maximum subarray sum:", maxSubarrayDivideConquer(subarrayArr));

const inversionArr: number[] = [8, 4, 2, 1];
console.log("Inversions count:", countInversions(inversionArr));

const matrixA: number[][] = [[1, 2], [3, 4]];
const matrixB: number[][] = [[5, 6], [7, 8]];
console.log("Matrix multiplication:", matrixMultiply(matrixA, matrixB));

const majorityArr: number[] = [3, 3, 4, 2, 4, 4, 2, 4, 4];
console.log("Majority element:", findMajorityElement(majorityArr));

const points: Point[] = [
    { x: 2, y: 3 },
    { x: 12, y: 30 },
    { x: 40, y: 50 },
    { x: 5, y: 1 },
    { x: 12, y: 10 },
    { x: 3, y: 4 }
];
console.log("Closest pair distance:", closestPairDistance(points));
```

### 14. Backtracking Traversal

#### Go Examples

```go
package main

import "fmt"

func main() {
    fmt.Println("Backtracking traversal examples:")
    
    // Generate all subsets
    arr := []int{1, 2, 3}
    fmt.Printf("All subsets of %v:\n", arr)
    subsets := generateSubsets(arr)
    for _, subset := range subsets {
        fmt.Printf("%v\n", subset)
    }
    
    // Generate all permutations
    fmt.Printf("\nAll permutations of %v:\n", arr)
    permutations := generatePermutations(arr)
    for _, perm := range permutations {
        fmt.Printf("%v\n", perm)
    }
    
    // Generate combinations
    fmt.Printf("\nCombinations of %v choose 2:\n", arr)
    combinations := generateCombinations(arr, 2)
    for _, comb := range combinations {
        fmt.Printf("%v\n", comb)
    }
    
    // N-Queens problem
    n := 4
    fmt.Printf("\nSolutions for %d-Queens problem:\n", n)
    queens := solveNQueens(n)
    for i, solution := range queens {
        fmt.Printf("Solution %d:\n", i+1)
        printChessboard(solution)
        fmt.Println()
    }
    
    // Sudoku solver
    sudoku := [][]int{
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9},
    }
    fmt.Println("Sudoku puzzle:")
    printSudoku(sudoku)
    
    if solveSudoku(sudoku) {
        fmt.Println("Solved sudoku:")
        printSudoku(sudoku)
    } else {
        fmt.Println("No solution exists")
    }
}

// Generate all subsets using backtracking
func generateSubsets(arr []int) [][]int {
    var result [][]int
    var current []int
    
    var backtrack func(int)
    backtrack = func(start int) {
        // Add current subset to result
        subset := make([]int, len(current))
        copy(subset, current)
        result = append(result, subset)
        
        // Try adding each remaining element
        for i := start; i < len(arr); i++ {
            current = append(current, arr[i])
            backtrack(i + 1)
            current = current[:len(current)-1] // backtrack
        }
    }
    
    backtrack(0)
    return result
}

// Generate all permutations using backtracking
func generatePermutations(arr []int) [][]int {
    var result [][]int
    used := make([]bool, len(arr))
    var current []int
    
    var backtrack func()
    backtrack = func() {
        if len(current) == len(arr) {
            perm := make([]int, len(current))
            copy(perm, current)
            result = append(result, perm)
            return
        }
        
        for i := 0; i < len(arr); i++ {
            if !used[i] {
                used[i] = true
                current = append(current, arr[i])
                backtrack()
                current = current[:len(current)-1] // backtrack
                used[i] = false
            }
        }
    }
    
    backtrack()
    return result
}

// Generate combinations using backtracking
func generateCombinations(arr []int, k int) [][]int {
    var result [][]int
    var current []int
    
    var backtrack func(int)
    backtrack = func(start int) {
        if len(current) == k {
            comb := make([]int, len(current))
            copy(comb, current)
            result = append(result, comb)
            return
        }
        
        for i := start; i < len(arr); i++ {
            current = append(current, arr[i])
            backtrack(i + 1)
            current = current[:len(current)-1] // backtrack
        }
    }
    
    backtrack(0)
    return result
}

// Solve N-Queens problem using backtracking
func solveNQueens(n int) [][][]string {
    var solutions [][][]string
    board := make([][]string, n)
    for i := range board {
        board[i] = make([]string, n)
        for j := range board[i] {
            board[i][j] = "."
        }
    }
    
    var backtrack func(int)
    backtrack = func(row int) {
        if row == n {
            // Found a solution
            solution := make([][]string, n)
            for i := range solution {
                solution[i] = make([]string, n)
                copy(solution[i], board[i])
            }
            solutions = append(solutions, solution)
            return
        }
        
        for col := 0; col < n; col++ {
            if isSafeQueen(board, row, col, n) {
                board[row][col] = "Q"
                backtrack(row + 1)
                board[row][col] = "." // backtrack
            }
        }
    }
    
    backtrack(0)
    return solutions
}

func isSafeQueen(board [][]string, row, col, n int) bool {
    // Check column
    for i := 0; i < row; i++ {
        if board[i][col] == "Q" {
            return false
        }
    }
    
    // Check diagonal (top-left to bottom-right)
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    
    // Check diagonal (top-right to bottom-left)
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    
    return true
}

func printChessboard(board [][]string) {
    for _, row := range board {
        for _, cell := range row {
            fmt.Printf("%s ", cell)
        }
        fmt.Println()
    }
}

// Solve Sudoku using backtracking
func solveSudoku(board [][]int) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                for num := 1; num <= 9; num++ {
                    if isSafeSudoku(board, i, j, num) {
                        board[i][j] = num
                        
                        if solveSudoku(board) {
                            return true
                        }
                        
                        board[i][j] = 0 // backtrack
                    }
                }
                return false
            }
        }
    }
    return true
}

func isSafeSudoku(board [][]int, row, col, num int) bool {
    // Check row
    for j := 0; j < 9; j++ {
        if board[row][j] == num {
            return false
        }
    }
    
    // Check column
    for i := 0; i < 9; i++ {
        if board[i][col] == num {
            return false
        }
    }
    
    // Check 3x3 box
    startRow, startCol := 3*(row/3), 3*(col/3)
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[startRow+i][startCol+j] == num {
                return false
            }
        }
    }
    
    return true
}

func printSudoku(board [][]int) {
    for i, row := range board {
        if i%3 == 0 && i != 0 {
            fmt.Println("------+-------+------")
        }
        for j, cell := range row {
            if j%3 == 0 && j != 0 {
                fmt.Print("| ")
            }
            if cell == 0 {
                fmt.Print(". ")
            } else {
                fmt.Printf("%d ", cell)
            }
        }
        fmt.Println()
    }
}

// Word search in 2D grid using backtracking
func wordSearch(board [][]byte, word string) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    
    rows, cols := len(board), len(board[0])
    
    var backtrack func(int, int, int) bool
    backtrack = func(row, col, index int) bool {
        if index == len(word) {
            return true
        }
        
        if row < 0 || row >= rows || col < 0 || col >= cols || 
           board[row][col] != word[index] {
            return false
        }
        
        // Mark current cell as visited
        temp := board[row][col]
        board[row][col] = '#'
        
        // Explore all 4 directions
        found := backtrack(row+1, col, index+1) ||
                backtrack(row-1, col, index+1) ||
                backtrack(row, col+1, index+1) ||
                backtrack(row, col-1, index+1)
        
        // Restore the cell
        board[row][col] = temp
        
        return found
    }
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if backtrack(i, j, 0) {
                return true
            }
        }
    }
    
    return false
}
```

#### TypeScript Examples

```typescript
// Backtracking traversal methods

// Generate all subsets using backtracking
function generateSubsets<T>(arr: T[]): T[][] {
    const result: T[][] = [];
    const current: T[] = [];
    
    function backtrack(start: number): void {
        // Add current subset to result
        result.push([...current]);
        
        // Try adding each remaining element
        for (let i = start; i < arr.length; i++) {
            current.push(arr[i]);
            backtrack(i + 1);
            current.pop(); // backtrack
        }
    }
    
    backtrack(0);
    return result;
}

// Generate all permutations using backtracking
function generatePermutations<T>(arr: T[]): T[][] {
    const result: T[][] = [];
    const used: boolean[] = new Array(arr.length).fill(false);
    const current: T[] = [];
    
    function backtrack(): void {
        if (current.length === arr.length) {
            result.push([...current]);
            return;
        }
        
        for (let i = 0; i < arr.length; i++) {
            if (!used[i]) {
                used[i] = true;
                current.push(arr[i]);
                backtrack();
                current.pop(); // backtrack
                used[i] = false;
            }
        }
    }
    
    backtrack();
    return result;
}

// Generate combinations using backtracking
function generateCombinations<T>(arr: T[], k: number): T[][] {
    const result: T[][] = [];
    const current: T[] = [];
    
    function backtrack(start: number): void {
        if (current.length === k) {
            result.push([...current]);
            return;
        }
        
        for (let i = start; i < arr.length; i++) {
            current.push(arr[i]);
            backtrack(i + 1);
            current.pop(); // backtrack
        }
    }
    
    backtrack(0);
    return result;
}

// Solve N-Queens problem using backtracking
function solveNQueens(n: number): string[][][] {
    const solutions: string[][][] = [];
    const board: string[][] = Array(n).fill(null).map(() => Array(n).fill('.'));
    
    function backtrack(row: number): void {
        if (row === n) {
            // Found a solution
            const solution = board.map(row => [...row]);
            solutions.push(solution);
            return;
        }
        
        for (let col = 0; col < n; col++) {
            if (isSafeQueen(board, row, col, n)) {
                board[row][col] = 'Q';
                backtrack(row + 1);
                board[row][col] = '.'; // backtrack
            }
        }
    }
    
    backtrack(0);
    return solutions;
}

function isSafeQueen(board: string[][], row: number, col: number, n: number): boolean {
    // Check column
    for (let i = 0; i < row; i++) {
        if (board[i][col] === 'Q') {
            return false;
        }
    }
    
    // Check diagonal (top-left to bottom-right)
    for (let i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
        if (board[i][j] === 'Q') {
            return false;
        }
    }
    
    // Check diagonal (top-right to bottom-left)
    for (let i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
        if (board[i][j] === 'Q') {
            return false;
        }
    }
    
    return true;
}

// Solve Sudoku using backtracking
function solveSudoku(board: number[][]): boolean {
    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            if (board[i][j] === 0) {
                for (let num = 1; num <= 9; num++) {
                    if (isSafeSudoku(board, i, j, num)) {
                        board[i][j] = num;
                        
                        if (solveSudoku(board)) {
                            return true;
                        }
                        
                        board[i][j] = 0; // backtrack
                    }
                }
                return false;
            }
        }
    }
    return true;
}

function isSafeSudoku(board: number[][], row: number, col: number, num: number): boolean {
    // Check row
    for (let j = 0; j < 9; j++) {
        if (board[row][j] === num) {
            return false;
        }
    }
    
    // Check column
    for (let i = 0; i < 9; i++) {
        if (board[i][col] === num) {
            return false;
        }
    }
    
    // Check 3x3 box
    const startRow = Math.floor(row / 3) * 3;
    const startCol = Math.floor(col / 3) * 3;
    
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
            if (board[startRow + i][startCol + j] === num) {
                return false;
            }
        }
    }
    
    return true;
}

// Word search in 2D grid using backtracking
function wordSearch(board: string[][], word: string): boolean {
    if (board.length === 0 || board[0].length === 0) {
        return false;
    }
    
    const rows = board.length;
    const cols = board[0].length;
    
    function backtrack(row: number, col: number, index: number): boolean {
        if (index === word.length) {
            return true;
        }
        
        if (row < 0 || row >= rows || col < 0 || col >= cols || 
            board[row][col] !== word[index]) {
            return false;
        }
        
        // Mark current cell as visited
        const temp = board[row][col];
        board[row][col] = '#';
        
        // Explore all 4 directions
        const found = backtrack(row + 1, col, index + 1) ||
                     backtrack(row - 1, col, index + 1) ||
                     backtrack(row, col + 1, index + 1) ||
                     backtrack(row, col - 1, index + 1);
        
        // Restore the cell
        board[row][col] = temp;
        
        return found;
    }
    
    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            if (backtrack(i, j, 0)) {
                return true;
            }
        }
    }
    
    return false;
}

// Generate all valid parentheses combinations
function generateParentheses(n: number): string[] {
    const result: string[] = [];
    
    function backtrack(current: string, open: number, close: number): void {
        if (current.length === 2 * n) {
            result.push(current);
            return;
        }
        
        if (open < n) {
            backtrack(current + '(', open + 1, close);
        }
        
        if (close < open) {
            backtrack(current + ')', open, close + 1);
        }
    }
    
    backtrack('', 0, 0);
    return result;
}

// Letter combinations of phone number
function letterCombinations(digits: string): string[] {
    if (digits.length === 0) return [];
    
    const phoneMap: { [key: string]: string } = {
        '2': 'abc',
        '3': 'def',
        '4': 'ghi',
        '5': 'jkl',
        '6': 'mno',
        '7': 'pqrs',
        '8': 'tuv',
        '9': 'wxyz'
    };
    
    const result: string[] = [];
    
    function backtrack(index: number, current: string): void {
        if (index === digits.length) {
            result.push(current);
            return;
        }
        
        const letters = phoneMap[digits[index]];
        for (const letter of letters) {
            backtrack(index + 1, current + letter);
        }
    }
    
    backtrack(0, '');
    return result;
}

// Palindrome partitioning
function palindromePartition(s: string): string[][] {
    const result: string[][] = [];
    const current: string[] = [];
    
    function isPalindrome(str: string): boolean {
        let left = 0;
        let right = str.length - 1;
        
        while (left < right) {
            if (str[left] !== str[right]) {
                return false;
            }
            left++;
            right--;
        }
        
        return true;
    }
    
    function backtrack(start: number): void {
        if (start === s.length) {
            result.push([...current]);
            return;
        }
        
        for (let end = start; end < s.length; end++) {
            const substring = s.substring(start, end + 1);
            if (isPalindrome(substring)) {
                current.push(substring);
                backtrack(end + 1);
                current.pop(); // backtrack
            }
        }
    }
    
    backtrack(0);
    return result;
}

// Rat in a maze problem
function ratInMaze(maze: number[][]): number[][] | null {
    const n = maze.length;
    const solution: number[][] = Array(n).fill(null).map(() => Array(n).fill(0));
    
    function isSafe(x: number, y: number): boolean {
        return x >= 0 && x < n && y >= 0 && y < n && maze[x][y] === 1;
    }
    
    function backtrack(x: number, y: number): boolean {
        if (x === n - 1 && y === n - 1 && maze[x][y] === 1) {
            solution[x][y] = 1;
            return true;
        }
        
        if (isSafe(x, y)) {
            solution[x][y] = 1;
            
            // Move right
            if (backtrack(x, y + 1)) {
                return true;
            }
            
            // Move down
            if (backtrack(x + 1, y)) {
                return true;
            }
            
            // Backtrack
            solution[x][y] = 0;
            return false;
        }
        
        return false;
    }
    
    if (backtrack(0, 0)) {
        return solution;
    }
    
    return null;
}

// Example usage
console.log("Backtracking examples:");

const arr: number[] = [1, 2, 3];
console.log("All subsets:", generateSubsets(arr));
console.log("All permutations:", generatePermutations(arr));
console.log("Combinations (choose 2):", generateCombinations(arr, 2));

const nQueensSolutions = solveNQueens(4);
console.log(`4-Queens solutions (${nQueensSolutions.length} total):`);
nQueensSolutions.forEach((solution, index) => {
    console.log(`Solution ${index + 1}:`);
    solution.forEach(row => console.log(row.join(' ')));
    console.log();
});

const sudoku: number[][] = [
    [5, 3, 0, 0, 7, 0, 0, 0, 0],
    [6, 0, 0, 1, 9, 5, 0, 0, 0],
    [0, 9, 8, 0, 0, 0, 0, 6, 0],
    [8, 0, 0, 0, 6, 0, 0, 0, 3],
    [4, 0, 0, 8, 0, 3, 0, 0, 1],
    [7, 0, 0, 0, 2, 0, 0, 0, 6],
    [0, 6, 0, 0, 0, 0, 2, 8, 0],
    [0, 0, 0, 4, 1, 9, 0, 0, 5],
    [0, 0, 0, 0, 8, 0, 0, 7, 9]
];

if (solveSudoku(sudoku)) {
    console.log("Solved Sudoku:");
    sudoku.forEach(row => console.log(row.join(' ')));
}

const wordGrid: string[][] = [
    ['A', 'B', 'C', 'E'],
    ['S', 'F', 'C', 'S'],
    ['A', 'D', 'E', 'E']
];

console.log("Word search for 'ABCCED':", wordSearch(wordGrid, 'ABCCED'));
console.log("Word search for 'SEE':", wordSearch(wordGrid, 'SEE'));

console.log("Generate parentheses (n=3):", generateParentheses(3));
console.log("Letter combinations for '23':", letterCombinations('23'));

console.log("Palindrome partitions for 'aab':", palindromePartition('aab'));

const maze: number[][] = [
    [1, 0, 0, 0],
    [1, 1, 0, 1],
    [0, 1, 0, 0],
    [1, 1, 1, 1]
];

const mazeSolution = ratInMaze(maze);
if (mazeSolution) {
    console.log("Rat in maze solution:");
    mazeSolution.forEach(row => console.log(row.join(' ')));
} else {
    console.log("No solution for rat in maze");
}
```

## Performance Considerations

### Time Complexity

| Traversal Method | Time Complexity | Space Complexity | Use Case |
|------------------|-----------------|------------------|----------|
| Index-based (for) | O(n) | O(1) | Simple iteration, direct access needed |
| Range-based (Go) | O(n) | O(1) | Clean syntax, don't need index |
| forEach (TS) | O(n) | O(1) | Functional style, side effects |
| map/filter/reduce | O(n) | O(n) | Transformation, functional programming |
| Two-pointer | O(n) | O(1) | Sorted arrays, pair finding |
| Recursive | O(n) | O(n) | Tree-like problems, divide & conquer |
| Sliding window | O(n) | O(1) | Subarray problems, optimization |
| Binary search | O(log n) | O(1) | Sorted arrays, search operations |
| Divide & conquer | O(n log n) | O(log n) | Sorting, optimization problems |
| Backtracking | O(2^n) or worse | O(n) | Combinatorial problems, puzzles |
| Generator | O(n) | O(1) | Lazy evaluation, memory efficiency |

### Advanced Complexity Analysis

#### Two-Pointer Techniques
- **Same direction**: O(n) time, O(1) space - for problems like removing duplicates
- **Opposite direction**: O(n) time, O(1) space - for palindrome checks, pair sums
- **Fast/slow pointers**: O(n) time, O(1) space - for cycle detection, middle finding

#### Recursive Patterns
- **Linear recursion**: O(n) time, O(n) space - simple traversals
- **Binary recursion**: O(2^n) time, O(n) space - generating permutations
- **Tail recursion**: O(n) time, O(1) space when optimized

#### Sliding Window Variants
- **Fixed window**: O(n) time, O(1) space - maximum in sliding windows
- **Variable window**: O(n) time, O(k) space - longest substring problems
- **Multiple windows**: O(n*k) time, O(k) space - complex pattern matching

#### Backtracking Complexity
- **Subsets**: O(2^n) time, O(n) space - exponential combinations
- **Permutations**: O(n!) time, O(n) space - factorial arrangements  
- **N-Queens**: O(n!) time, O(n) space - constraint satisfaction
- **Sudoku**: O(9^(n*n)) time, O(n*n) space - exponential with pruning

### Performance Tips

1. **Use index-based loops** for maximum performance in tight loops
2. **Prefer range-based loops** for readability when index isn't needed  
3. **Use array methods** (map/filter/reduce) for functional transformations
4. **Consider generators** for memory-efficient processing of large arrays
5. **Use async processing** for I/O bound operations
6. **Avoid nested array methods** when performance is critical

### Memory Considerations

```typescript
// Memory efficient traversal for large arrays
function* chunkedTraversal<T>(array: T[], chunkSize: number): Generator<T[], void, unknown> {
    for (let i = 0; i < array.length; i += chunkSize) {
        yield array.slice(i, i + chunkSize);
    }
}

// Process large arrays in chunks
const largeArray = Array.from({length: 1000000}, (_, i) => i);
for (const chunk of chunkedTraversal(largeArray, 1000)) {
    // Process chunk without loading entire array into memory
    console.log(`Processing chunk of ${chunk.length} items`);
}
```

## Use Cases and Best Practices

### When to Use Each Method

1. **Index-based traversal**: When you need the index or maximum performance
2. **Range-based traversal**: When you only need values and want clean code
3. **Functional methods**: When transforming data or using functional programming
4. **Generators**: When dealing with large datasets or lazy evaluation
5. **Async traversal**: When processing involves I/O operations

### Best Practices

1. **Choose the right method** based on your specific needs
2. **Consider memory usage** for large arrays
3. **Use type annotations** in TypeScript for better code safety
4. **Handle edge cases** like empty arrays
5. **Prefer immutable operations** when possible
6. **Use appropriate error handling** for async operations

## Common Pitfalls

1. **Modifying array during traversal** can lead to unexpected behavior
2. **Off-by-one errors** in index-based loops
3. **Performance issues** with nested array methods
4. **Memory leaks** with closures in callbacks
5. **Race conditions** in async parallel processing