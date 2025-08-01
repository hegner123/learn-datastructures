# Arrays

## What is an Array?

An array is a data structure that stores elements of the same type in contiguous memory locations. Arrays provide constant-time access to elements using an index and are one of the most fundamental data structures in computer science.

## Key Properties

- **Fixed size**: Arrays have a predetermined size that cannot be changed after creation
- **Indexed access**: Elements are accessed using zero-based indexing
- **Homogeneous**: All elements must be of the same data type
- **Contiguous memory**: Elements are stored in consecutive memory locations

## Time Complexity

- **Access**: O(1)
- **Search**: O(n)
- **Insertion**: O(n) - requires shifting elements
- **Deletion**: O(n) - requires shifting elements

## Go Code Examples

### Basic Array Declaration and Initialization

```go
package main

import "fmt"

func main() {
    // Declare an array of 5 integers
    var numbers [5]int
    
    // Initialize with values
    fruits := [3]string{"apple", "banana", "orange"}
    
    // Let Go infer the size
    colors := [...]string{"red", "green", "blue", "yellow"}
    
    fmt.Println("Numbers:", numbers)  // [0 0 0 0 0]
    fmt.Println("Fruits:", fruits)    // [apple banana orange]
    fmt.Println("Colors:", colors)    // [red green blue yellow]
}
```

### Accessing and Modifying Elements

```go
package main

import "fmt"

func main() {
    scores := [5]int{85, 92, 78, 96, 88}
    
    // Access elements
    fmt.Println("First score:", scores[0])
    fmt.Println("Last score:", scores[4])
    
    // Modify elements
    scores[2] = 82
    fmt.Println("Updated scores:", scores)
    
    // Get array length
    fmt.Println("Array length:", len(scores))
}
```

### Iterating Through Arrays

```go
package main

import "fmt"

func main() {
    temperatures := [7]float64{23.5, 25.1, 22.8, 26.3, 24.7, 21.9, 23.2}
    
    // Using traditional for loop
    fmt.Println("Using index-based loop:")
    for i := 0; i < len(temperatures); i++ {
        fmt.Printf("Day %d: %.1f°C\n", i+1, temperatures[i])
    }
    
    // Using range
    fmt.Println("\nUsing range:")
    for index, temp := range temperatures {
        fmt.Printf("Day %d: %.1f°C\n", index+1, temp)
    }
    
    // Using range with blank identifier
    fmt.Println("\nJust values:")
    for _, temp := range temperatures {
        fmt.Printf("%.1f°C ", temp)
    }
    fmt.Println()
}
```

### Array Functions

```go
package main

import "fmt"

func main() {
    data := [6]int{15, 3, 9, 1, 12, 7}
    
    fmt.Println("Original array:", data)
    fmt.Println("Sum:", sum(data))
    fmt.Println("Maximum:", findMax(data))
    fmt.Println("Minimum:", findMin(data))
    
    // Arrays are passed by value (copied)
    modifyArray(data)
    fmt.Println("After modifyArray:", data) // Original unchanged
    
    // To modify, pass by reference
    modifyArrayByRef(&data)
    fmt.Println("After modifyArrayByRef:", data) // Modified
}

func sum(arr [6]int) int {
    total := 0
    for _, value := range arr {
        total += value
    }
    return total
}

func findMax(arr [6]int) int {
    max := arr[0]
    for _, value := range arr {
        if value > max {
            max = value
        }
    }
    return max
}

func findMin(arr [6]int) int {
    min := arr[0]
    for _, value := range arr {
        if value < min {
            min = value
        }
    }
    return min
}

func modifyArray(arr [6]int) {
    arr[0] = 999 // This won't affect the original
}

func modifyArrayByRef(arr *[6]int) {
    arr[0] = 999 // This will modify the original
}
```

### Multi-dimensional Arrays

```go
package main

import "fmt"

func main() {
    // 2D array (matrix)
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("2D Array:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
    
    // 3D array
    cube := [2][2][2]int{
        {
            {1, 2},
            {3, 4},
        },
        {
            {5, 6},
            {7, 8},
        },
    }
    
    fmt.Println("\n3D Array:")
    for i := range cube {
        fmt.Printf("Layer %d:\n", i)
        for j := range cube[i] {
            for k := range cube[i][j] {
                fmt.Printf("%d ", cube[i][j][k])
            }
            fmt.Println()
        }
    }
}
```

## TypeScript Implementation

### Basic Array Operations

```typescript
// Array declaration and initialization with type annotations
const numbers: number[] = [1, 2, 3, 4, 5];
const fruits: string[] = ['apple', 'banana', 'orange'];
const mixed: (number | string | boolean | null)[] = [1, 'hello', true, null];

// Alternative array type syntax
const scores: Array<number> = [85, 92, 78, 96, 88];

// Access elements
console.log('First element:', numbers[0]);
console.log('Last element:', numbers[numbers.length - 1]);

// Modify elements
numbers[2] = 10;
console.log('Modified array:', numbers);

// Array length
console.log('Array length:', numbers.length);
```

### Array Methods and Operations

```typescript
// Common array methods with type annotations
const data: number[] = [15, 3, 9, 1, 12, 7];

console.log('Original array:', data);

// Adding elements
data.push(20);           // Add to end
data.unshift(5);         // Add to beginning
console.log('After adding:', data);

// Removing elements
const lastElement: number | undefined = data.pop();      // Remove from end
const firstElement: number | undefined = data.shift();   // Remove from beginning
console.log('After removing:', data);
console.log('Removed elements:', firstElement, lastElement);

// Finding elements
const index: number = data.indexOf(9);       // Find index of element
const exists: boolean = data.includes(12);    // Check if element exists
console.log('Index of 9:', index);
console.log('Contains 12:', exists);

// Array iteration
console.log('Using for loop:');
for (let i: number = 0; i < data.length; i++) {
    console.log(`Element ${i}: ${data[i]}`);
}

console.log('Using forEach:');
data.forEach((element: number, index: number) => {
    console.log(`Element ${index}: ${element}`);
});

console.log('Using for...of:');
for (const element of data) {
    console.log('Element:', element);
}
```

### Array Utility Functions

```typescript
// Utility functions for arrays with type annotations
const numbers: number[] = [15, 3, 9, 1, 12, 7];

// Sum of array elements
function sum(arr: number[]): number {
    return arr.reduce((total: number, num: number) => total + num, 0);
}

// Find maximum element
function findMax(arr: number[]): number {
    return Math.max(...arr);
    // Alternative: return arr.reduce((max: number, num: number) => Math.max(max, num), -Infinity);
}

// Find minimum element
function findMin(arr: number[]): number {
    return Math.min(...arr);
    // Alternative: return arr.reduce((min: number, num: number) => Math.min(min, num), Infinity);
}

// Calculate average
function average(arr: number[]): number {
    return arr.length > 0 ? sum(arr) / arr.length : 0;
}

// Filter even numbers
function filterEven(arr: number[]): number[] {
    return arr.filter((num: number) => num % 2 === 0);
}

// Sort array (creates new array)
function sortArray(arr: number[]): number[] {
    return [...arr].sort((a: number, b: number) => a - b);
}

// Reverse array (creates new array)
function reverseArray<T>(arr: T[]): T[] {
    return [...arr].reverse();
}

// Remove duplicates (generic function)
function removeDuplicates<T>(arr: T[]): T[] {
    return [...new Set(arr)];
}

// Example usage
console.log('Original:', numbers);
console.log('Sum:', sum(numbers));
console.log('Max:', findMax(numbers));
console.log('Min:', findMin(numbers));
console.log('Average:', average(numbers));
console.log('Even numbers:', filterEven(numbers));
console.log('Sorted:', sortArray(numbers));
console.log('Reversed:', reverseArray(numbers));

const withDuplicates: number[] = [1, 2, 2, 3, 3, 3, 4];
console.log('With duplicates:', withDuplicates);
console.log('Without duplicates:', removeDuplicates(withDuplicates));
```

### Multi-dimensional Arrays

```typescript
// 2D array (matrix) with type annotations
const matrix: number[][] = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
];

console.log('2D Array:');
for (let i: number = 0; i < matrix.length; i++) {
    let row: string = '';
    for (let j: number = 0; j < matrix[i].length; j++) {
        row += matrix[i][j] + ' ';
    }
    console.log(row);
}

// Alternative using forEach
console.log('2D Array using forEach:');
matrix.forEach((row: number[], i: number) => {
    console.log(`Row ${i}: ${row.join(' ')}`);
});

// 3D array with type annotations
const cube: number[][][] = [
    [
        [1, 2],
        [3, 4]
    ],
    [
        [5, 6],
        [7, 8]
    ]
];

console.log('3D Array:');
cube.forEach((layer: number[][], i: number) => {
    console.log(`Layer ${i}:`);
    layer.forEach((row: number[]) => {
        console.log(row.join(' '));
    });
});

// Matrix operations with type annotations
function addMatrices(matrix1: number[][], matrix2: number[][]): number[][] {
    if (matrix1.length !== matrix2.length || 
        matrix1[0].length !== matrix2[0].length) {
        throw new Error('Matrices must have same dimensions');
    }
    
    const result: number[][] = [];
    for (let i: number = 0; i < matrix1.length; i++) {
        result[i] = [];
        for (let j: number = 0; j < matrix1[i].length; j++) {
            result[i][j] = matrix1[i][j] + matrix2[i][j];
        }
    }
    return result;
}

const matrix1: number[][] = [[1, 2], [3, 4]];
const matrix2: number[][] = [[5, 6], [7, 8]];
const sum: number[][] = addMatrices(matrix1, matrix2);
console.log('Matrix addition result:', sum);
```

### Array Searching and Sorting

```typescript
// Linear search with type annotations
function linearSearch(arr: number[], target: number): number {
    for (let i: number = 0; i < arr.length; i++) {
        if (arr[i] === target) {
            return i;
        }
    }
    return -1;
}

// Binary search (requires sorted array)
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

// Bubble sort with type annotations
function bubbleSort(arr: number[]): number[] {
    const sorted: number[] = [...arr]; // Create copy
    
    for (let i: number = 0; i < sorted.length - 1; i++) {
        for (let j: number = 0; j < sorted.length - i - 1; j++) {
            if (sorted[j] > sorted[j + 1]) {
                // Swap elements
                [sorted[j], sorted[j + 1]] = [sorted[j + 1], sorted[j]];
            }
        }
    }
    
    return sorted;
}

// Generic search function
function genericSearch<T>(arr: T[], target: T): number {
    for (let i: number = 0; i < arr.length; i++) {
        if (arr[i] === target) {
            return i;
        }
    }
    return -1;
}

// Example usage
const searchArray: number[] = [64, 34, 25, 12, 22, 11, 90];
console.log('Original array:', searchArray);

// Linear search
console.log('Linear search for 25:', linearSearch(searchArray, 25));

// Sort and binary search
const sortedArray: number[] = bubbleSort(searchArray);
console.log('Sorted array:', sortedArray);
console.log('Binary search for 25:', binarySearch(sortedArray, 25));

// Generic search example
const stringArray: string[] = ['apple', 'banana', 'orange'];
console.log('Generic search for "banana":', genericSearch(stringArray, 'banana'));
```

### Advanced Array Operations

```typescript
// Array destructuring with type annotations
const colors: string[] = ['red', 'green', 'blue', 'yellow'];
const [first, second, ...rest]: [string, string, ...string[]] = colors;
console.log('First:', first);
console.log('Second:', second);
console.log('Rest:', rest);

// Spread operator
const arr1: number[] = [1, 2, 3];
const arr2: number[] = [4, 5, 6];
const combined: number[] = [...arr1, ...arr2];
console.log('Combined:', combined);

// Array methods chaining with type annotations
const numbers: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
const result: number = numbers
    .filter((num: number) => num % 2 === 0)    // Get even numbers
    .map((num: number) => num * 2)             // Double each number
    .reduce((sum: number, num: number) => sum + num, 0); // Sum all numbers

console.log('Chained operations result:', result);

// Map and Set from arrays
const arrayWithDuplicates: number[] = [1, 2, 2, 3, 3, 3, 4, 4, 4, 4];
const uniqueArray: number[] = [...new Set(arrayWithDuplicates)];
console.log('Unique array:', uniqueArray);

// Array from string
const stringArray: string[] = Array.from('hello');
console.log('Array from string:', stringArray);

// Array with specific length and fill
const filledArray: number[] = new Array(5).fill(0);
console.log('Filled array:', filledArray);

// Array.from with mapping function
const squares: number[] = Array.from({length: 5}, (_: undefined, i: number) => i * i);
console.log('Squares:', squares);

// Type-safe array operations with interfaces
interface Person {
    name: string;
    age: number;
    city: string;
}

const people: Person[] = [
    { name: 'Alice', age: 25, city: 'New York' },
    { name: 'Bob', age: 30, city: 'Los Angeles' },
    { name: 'Charlie', age: 35, city: 'Chicago' }
];

// Filter and map with type safety
const adults: Person[] = people.filter((person: Person) => person.age >= 18);
const names: string[] = people.map((person: Person) => person.name);
const averageAge: number = people.reduce((sum: number, person: Person) => sum + person.age, 0) / people.length;

console.log('Adults:', adults);
console.log('Names:', names);
console.log('Average age:', averageAge);

// Tuple arrays
const coordinates: [number, number][] = [[1, 2], [3, 4], [5, 6]];
const point: [number, number] = [10, 20];
coordinates.push(point);
console.log('Coordinates:', coordinates);

// Readonly arrays
const readonlyNumbers: readonly number[] = [1, 2, 3, 4, 5];
// readonlyNumbers.push(6); // This would cause a TypeScript error
console.log('Readonly array:', readonlyNumbers);
```

## Common Use Cases

- **Data storage**: Storing collections of similar data
- **Mathematical operations**: Matrix operations, mathematical computations
- **Lookup tables**: Fast access to predefined values
- **Buffer management**: Fixed-size buffers for data processing
- **Implementation of other data structures**: Foundation for stacks, queues, etc.
- **Data processing**: Filtering, mapping, and transforming datasets

## TypeScript vs Go Arrays

### Key Differences

| Feature | TypeScript | Go |
|---------|------------|-----|
| **Size** | Dynamic | Fixed at compile time |
| **Type Safety** | Static typing with runtime flexibility | Static typing |
| **Memory** | Managed by JS engine | Explicit memory layout |
| **Methods** | Rich built-in methods | Basic operations, slices for flexibility |
| **Performance** | Optimized by JS engine | Predictable performance |
| **Generics** | Full generic support | Limited generic support |

### TypeScript Dynamic Arrays
```typescript
// TypeScript arrays are dynamic with type safety
let arr: number[] = [1, 2, 3];
arr.push(4);           // Can grow
arr.pop();             // Can shrink
// arr.push('string'); // This would cause a TypeScript error
console.log(arr.length); // 4

// TypeScript provides compile-time type checking
const mixedArray: (number | string)[] = [1, 'hello', 2, 'world'];
```

### Go Fixed Arrays vs Slices
```go
// Go arrays are fixed
var arr [3]int = [3]int{1, 2, 3}
// arr[3] = 4  // This would cause a compile error

// Go slices are dynamic (similar to TypeScript arrays)
slice := []int{1, 2, 3}
slice = append(slice, 4) // Can grow dynamically
```

## Advantages

- Constant-time access to elements by index
- Memory efficient for storing homogeneous data
- Cache-friendly due to contiguous memory layout
- Simple and intuitive data structure
- Foundation for implementing other data structures

## Disadvantages

- Fixed size in many languages (like Go)
- Expensive insertion/deletion operations (O(n))
- Memory waste if not fully utilized
- Lack of flexibility for dynamic data requirements
