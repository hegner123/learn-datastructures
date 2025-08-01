# Queues

## What is a Queue?

A queue is a linear data structure that follows the First In, First Out (FIFO) principle. Elements are added at one end (rear/back) and removed from the other end (front). Think of it like a line of people waiting - the first person in line is the first to be served.

## Key Properties

- **FIFO (First In, First Out)**: The first element added is the first one to be removed
- **Two access points**: Elements are added at the rear and removed from the front
- **Dynamic size**: Can grow and shrink during runtime
- **Sequential processing**: Elements are processed in the order they arrive

## Core Operations

- **Enqueue**: Add an element to the rear of the queue
- **Dequeue**: Remove and return the front element
- **Front/Peek**: View the front element without removing it
- **Rear**: View the rear element without removing it
- **IsEmpty**: Check if the queue is empty
- **Size**: Get the number of elements in the queue

## Time Complexity

- **Enqueue**: O(1)
- **Dequeue**: O(1)
- **Front/Peek**: O(1)
- **Search**: O(n)

## Types of Queues

- **Simple Queue**: Basic FIFO queue
- **Circular Queue**: Rear connects back to front when space is available
- **Priority Queue**: Elements have priorities, highest priority served first
- **Deque (Double-ended Queue)**: Insertion and deletion at both ends

## Go Implementation

### Basic Queue using Slice

```go
package main

import (
    "errors"
    "fmt"
)

// Queue represents a basic queue data structure
type Queue struct {
    items []int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
    return &Queue{
        items: make([]int, 0),
    }
}

// Enqueue adds an element to the rear of the queue
func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

// Dequeue removes and returns the front element
func (q *Queue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, nil
}

// Front returns the front element without removing it
func (q *Queue) Front() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    
    return q.items[0], nil
}

// Rear returns the rear element without removing it
func (q *Queue) Rear() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    
    return q.items[len(q.items)-1], nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
    return len(q.items)
}

// Display prints all elements in the queue
func (q *Queue) Display() {
    if q.IsEmpty() {
        fmt.Println("Queue is empty")
        return
    }
    
    fmt.Print("Front -> ")
    for i, item := range q.items {
        if i == len(q.items)-1 {
            fmt.Printf("%d", item)
        } else {
            fmt.Printf("%d -> ", item)
        }
    }
    fmt.Println(" <- Rear")
}

func main() {
    queue := NewQueue()
    
    // Enqueue elements
    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)
    queue.Enqueue(40)
    
    fmt.Println("After enqueuing 10, 20, 30, 40:")
    queue.Display()
    
    // Check front and rear
    if front, err := queue.Front(); err == nil {
        fmt.Printf("Front element: %d\n", front)
    }
    
    if rear, err := queue.Rear(); err == nil {
        fmt.Printf("Rear element: %d\n", rear)
    }
    
    // Dequeue elements
    if item, err := queue.Dequeue(); err == nil {
        fmt.Printf("Dequeued: %d\n", item)
    }
    
    if item, err := queue.Dequeue(); err == nil {
        fmt.Printf("Dequeued: %d\n", item)
    }
    
    fmt.Println("After dequeuing twice:")
    queue.Display()
    
    fmt.Printf("Queue size: %d\n", queue.Size())
    fmt.Printf("Is empty: %t\n", queue.IsEmpty())
}
```

### Circular Queue Implementation

```go
package main

import (
    "errors"
    "fmt"
)

// CircularQueue represents a circular queue data structure
type CircularQueue struct {
    items []int
    front int
    rear  int
    size  int
    capacity int
}

// NewCircularQueue creates a new circular queue with given capacity
func NewCircularQueue(capacity int) *CircularQueue {
    return &CircularQueue{
        items:    make([]int, capacity),
        front:    0,
        rear:     -1,
        size:     0,
        capacity: capacity,
    }
}

// Enqueue adds an element to the rear of the circular queue
func (cq *CircularQueue) Enqueue(item int) error {
    if cq.IsFull() {
        return errors.New("queue is full")
    }
    
    cq.rear = (cq.rear + 1) % cq.capacity
    cq.items[cq.rear] = item
    cq.size++
    return nil
}

// Dequeue removes and returns the front element
func (cq *CircularQueue) Dequeue() (int, error) {
    if cq.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    
    item := cq.items[cq.front]
    cq.front = (cq.front + 1) % cq.capacity
    cq.size--
    return item, nil
}

// Front returns the front element without removing it
func (cq *CircularQueue) Front() (int, error) {
    if cq.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    
    return cq.items[cq.front], nil
}

// Rear returns the rear element without removing it
func (cq *CircularQueue) Rear() (int, error) {
    if cq.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    
    return cq.items[cq.rear], nil
}

// IsEmpty checks if the queue is empty
func (cq *CircularQueue) IsEmpty() bool {
    return cq.size == 0
}

// IsFull checks if the queue is full
func (cq *CircularQueue) IsFull() bool {
    return cq.size == cq.capacity
}

// Size returns the number of elements in the queue
func (cq *CircularQueue) Size() int {
    return cq.size
}

// Display prints all elements in the circular queue
func (cq *CircularQueue) Display() {
    if cq.IsEmpty() {
        fmt.Println("Circular queue is empty")
        return
    }
    
    fmt.Print("Front -> ")
    for i := 0; i < cq.size; i++ {
        index := (cq.front + i) % cq.capacity
        if i == cq.size-1 {
            fmt.Printf("%d", cq.items[index])
        } else {
            fmt.Printf("%d -> ", cq.items[index])
        }
    }
    fmt.Println(" <- Rear")
}

func main() {
    cq := NewCircularQueue(5)
    
    // Enqueue elements
    cq.Enqueue(10)
    cq.Enqueue(20)
    cq.Enqueue(30)
    cq.Enqueue(40)
    cq.Enqueue(50)
    
    fmt.Println("After enqueuing 5 elements:")
    cq.Display()
    
    // Try to enqueue when full
    if err := cq.Enqueue(60); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
    
    // Dequeue some elements
    cq.Dequeue()
    cq.Dequeue()
    
    fmt.Println("After dequeuing 2 elements:")
    cq.Display()
    
    // Add more elements (demonstrating circular nature)
    cq.Enqueue(60)
    cq.Enqueue(70)
    
    fmt.Println("After enqueuing 2 more elements:")
    cq.Display()
}
```

### Priority Queue Implementation

```go
package main

import (
    "container/heap"
    "fmt"
)

// PriorityItem represents an item with priority
type PriorityItem struct {
    value    string
    priority int
    index    int
}

// PriorityQueue implements a priority queue using heap
type PriorityQueue []*PriorityItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority > pq[j].priority // Higher priority first
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*PriorityItem)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    item.index = -1
    *pq = old[0 : n-1]
    return item
}

func main() {
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    
    // Add items with different priorities
    heap.Push(&pq, &PriorityItem{
        value:    "Task 1",
        priority: 3,
    })
    heap.Push(&pq, &PriorityItem{
        value:    "Task 2",
        priority: 1,
    })
    heap.Push(&pq, &PriorityItem{
        value:    "Task 3",
        priority: 5,
    })
    heap.Push(&pq, &PriorityItem{
        value:    "Task 4",
        priority: 2,
    })
    
    fmt.Println("Processing tasks by priority:")
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*PriorityItem)
        fmt.Printf("Processing: %s (Priority: %d)\n", item.value, item.priority)
    }
}
```

## TypeScript Implementation

### Basic Queue

```typescript
class Queue<T> {
    private items: T[];
    
    constructor() {
        this.items = [];
    }
    
    // Add element to rear of queue
    public enqueue(item: T): void {
        this.items.push(item);
    }
    
    // Remove and return front element
    public dequeue(): T {
        if (this.isEmpty()) {
            throw new Error("Queue is empty");
        }
        return this.items.shift()!;
    }
    
    // View front element without removing
    public front(): T {
        if (this.isEmpty()) {
            throw new Error("Queue is empty");
        }
        return this.items[0];
    }
    
    // View rear element without removing
    public rear(): T {
        if (this.isEmpty()) {
            throw new Error("Queue is empty");
        }
        return this.items[this.items.length - 1];
    }
    
    // Check if queue is empty
    public isEmpty(): boolean {
        return this.items.length === 0;
    }
    
    // Get queue size
    public size(): number {
        return this.items.length;
    }
    
    // Display queue contents
    public display(): void {
        if (this.isEmpty()) {
            console.log("Queue is empty");
            return;
        }
        
        console.log("Front -> " + this.items.join(" -> ") + " <- Rear");
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
const queue = new Queue<number>();

// Enqueue elements
queue.enqueue(10);
queue.enqueue(20);
queue.enqueue(30);
queue.enqueue(40);

console.log("After enqueuing 10, 20, 30, 40:");
queue.display();

// Check front and rear
console.log("Front element:", queue.front());
console.log("Rear element:", queue.rear());

// Dequeue elements
console.log("Dequeued:", queue.dequeue());
console.log("Dequeued:", queue.dequeue());

console.log("After dequeuing twice:");
queue.display();

console.log("Queue size:", queue.size());
console.log("Is empty:", queue.isEmpty());
```

### Circular Queue

```typescript
class CircularQueue<T> {
    private items: (T | undefined)[];
    private front: number;
    private rear: number;
    private size: number;
    private capacity: number;
    
    constructor(capacity: number) {
        this.items = new Array(capacity);
        this.front = 0;
        this.rear = -1;
        this.size = 0;
        this.capacity = capacity;
    }
    
    // Add element to rear
    public enqueue(item: T): void {
        if (this.isFull()) {
            throw new Error("Queue is full");
        }
        
        this.rear = (this.rear + 1) % this.capacity;
        this.items[this.rear] = item;
        this.size++;
    }
    
    // Remove and return front element
    public dequeue(): T {
        if (this.isEmpty()) {
            throw new Error("Queue is empty");
        }
        
        const item = this.items[this.front]!;
        this.items[this.front] = undefined;
        this.front = (this.front + 1) % this.capacity;
        this.size--;
        return item;
    }
    
    // View front element
    public getFront(): T {
        if (this.isEmpty()) {
            throw new Error("Queue is empty");
        }
        return this.items[this.front]!;
    }
    
    // View rear element
    public getRear(): T {
        if (this.isEmpty()) {
            throw new Error("Queue is empty");
        }
        return this.items[this.rear]!;
    }
    
    // Check if empty
    public isEmpty(): boolean {
        return this.size === 0;
    }
    
    // Check if full
    public isFull(): boolean {
        return this.size === this.capacity;
    }
    
    // Get current size
    public getSize(): number {
        return this.size;
    }
    
    // Display queue
    public display(): void {
        if (this.isEmpty()) {
            console.log("Circular queue is empty");
            return;
        }
        
        const elements: T[] = [];
        for (let i = 0; i < this.size; i++) {
            const index = (this.front + i) % this.capacity;
            elements.push(this.items[index]!);
        }
        
        console.log("Front -> " + elements.join(" -> ") + " <- Rear");
    }
}

// Example usage
const cq = new CircularQueue<number>(5);

// Fill the queue
for (let i = 1; i <= 5; i++) {
    cq.enqueue(i * 10);
}

console.log("After filling circular queue:");
cq.display();

// Dequeue some elements
cq.dequeue();
cq.dequeue();

console.log("After dequeuing 2 elements:");
cq.display();

// Add more elements (demonstrating circular nature)
cq.enqueue(60);
cq.enqueue(70);

console.log("After enqueuing 2 more elements:");
cq.display();
```

### Priority Queue

```typescript
interface PriorityItem<T> {
    item: T;
    priority: number;
}

class PriorityQueue<T> {
    private items: PriorityItem<T>[];
    
    constructor() {
        this.items = [];
    }
    
    // Add element with priority
    public enqueue(item: T, priority: number): void {
        const queueElement: PriorityItem<T> = { item, priority };
        let added = false;
        
        for (let i = 0; i < this.items.length; i++) {
            if (queueElement.priority > this.items[i].priority) {
                this.items.splice(i, 0, queueElement);
                added = true;
                break;
            }
        }
        
        if (!added) {
            this.items.push(queueElement);
        }
    }
    
    // Remove and return highest priority element
    public dequeue(): T {
        if (this.isEmpty()) {
            throw new Error("Priority queue is empty");
        }
        return this.items.shift()!.item;
    }
    
    // View highest priority element
    public front(): T {
        if (this.isEmpty()) {
            throw new Error("Priority queue is empty");
        }
        return this.items[0].item;
    }
    
    // Check if empty
    public isEmpty(): boolean {
        return this.items.length === 0;
    }
    
    // Get size
    public size(): number {
        return this.items.length;
    }
    
    // Display queue with priorities
    public display(): void {
        if (this.isEmpty()) {
            console.log("Priority queue is empty");
            return;
        }
        
        console.log("Priority Queue (highest to lowest priority):");
        this.items.forEach((element, index) => {
            console.log(`${index + 1}. ${element.item} (Priority: ${element.priority})`);
        });
    }
}

// Example usage
const pq = new PriorityQueue<string>();

// Add tasks with different priorities
pq.enqueue("Task 1", 3);
pq.enqueue("Task 2", 1);
pq.enqueue("Task 3", 5);
pq.enqueue("Task 4", 2);

console.log("Priority Queue:");
pq.display();

console.log("\nProcessing tasks by priority:");
while (!pq.isEmpty()) {
    console.log("Processing:", pq.dequeue());
}
```

### Queue Applications

```typescript
// BFS (Breadth-First Search) implementation
class Graph {
    private adjacencyList: { [key: string]: string[] };
    
    constructor() {
        this.adjacencyList = {};
    }
    
    public addVertex(vertex: string): void {
        if (!this.adjacencyList[vertex]) {
            this.adjacencyList[vertex] = [];
        }
    }
    
    public addEdge(vertex1: string, vertex2: string): void {
        this.adjacencyList[vertex1].push(vertex2);
        this.adjacencyList[vertex2].push(vertex1);
    }
    
    public bfs(startVertex: string): string[] {
        const queue = new Queue<string>();
        const visited: { [key: string]: boolean } = {};
        const result: string[] = [];
        
        queue.enqueue(startVertex);
        visited[startVertex] = true;
        
        while (!queue.isEmpty()) {
            const currentVertex: string = queue.dequeue();
            result.push(currentVertex);
            
            this.adjacencyList[currentVertex].forEach(neighbor => {
                if (!visited[neighbor]) {
                    visited[neighbor] = true;
                    queue.enqueue(neighbor);
                }
            });
        }
        
        return result;
    }
}

// Task scheduler using queue
interface Task {
    name: string;
    duration: number;
}

class TaskScheduler {
    private taskQueue: Queue<Task>;
    private isProcessing: boolean;
    
    constructor() {
        this.taskQueue = new Queue<Task>();
        this.isProcessing = false;
    }
    
    public addTask(task: Task): void {
        this.taskQueue.enqueue(task);
        console.log(`Task added: ${task.name}`);
        
        if (!this.isProcessing) {
            this.processTasks();
        }
    }
    
    private async processTasks(): Promise<void> {
        this.isProcessing = true;
        
        while (!this.taskQueue.isEmpty()) {
            const task: Task = this.taskQueue.dequeue();
            console.log(`Processing: ${task.name}`);
            
            // Simulate task processing
            await new Promise(resolve => setTimeout(resolve, task.duration));
            
            console.log(`Completed: ${task.name}`);
        }
        
        this.isProcessing = false;
        console.log("All tasks completed!");
    }
}

// Cache implementation using queue (LRU-like behavior)
class Cache<K, V> {
    private capacity: number;
    private cache: Map<K, V>;
    private queue: Queue<K>;
    
    constructor(capacity: number) {
        this.capacity = capacity;
        this.cache = new Map<K, V>();
        this.queue = new Queue<K>();
    }
    
    public get(key: K): V | null {
        if (this.cache.has(key)) {
            // Move to end of queue (most recently used)
            this.updateQueue(key);
            return this.cache.get(key)!;
        }
        return null;
    }
    
    public set(key: K, value: V): void {
        if (this.cache.has(key)) {
            this.cache.set(key, value);
            this.updateQueue(key);
        } else {
            if (this.cache.size >= this.capacity) {
                // Remove least recently used
                const lru: K = this.queue.dequeue();
                this.cache.delete(lru);
            }
            
            this.cache.set(key, value);
            this.queue.enqueue(key);
        }
    }
    
    private updateQueue(key: K): void {
        // Remove key from queue and add to end
        const tempQueue = new Queue<K>();
        while (!this.queue.isEmpty()) {
            const item: K = this.queue.dequeue();
            if (item !== key) {
                tempQueue.enqueue(item);
            }
        }
        
        while (!tempQueue.isEmpty()) {
            this.queue.enqueue(tempQueue.dequeue());
        }
        
        this.queue.enqueue(key);
    }
    
    public display(): void {
        console.log("Cache contents:", [...this.cache.entries()]);
        console.log("Usage order:", this.queue.toArray());
    }
}

// Example usage
console.log("BFS Example:");
const graph = new Graph();
graph.addVertex("A");
graph.addVertex("B");
graph.addVertex("C");
graph.addVertex("D");
graph.addVertex("E");

graph.addEdge("A", "B");
graph.addEdge("A", "C");
graph.addEdge("B", "D");
graph.addEdge("C", "E");

console.log("BFS traversal from A:", graph.bfs("A"));

console.log("\nTask Scheduler Example:");
const scheduler = new TaskScheduler();
scheduler.addTask({ name: "Task 1", duration: 1000 });
scheduler.addTask({ name: "Task 2", duration: 500 });
scheduler.addTask({ name: "Task 3", duration: 1500 });

console.log("\nCache Example:");
const cache = new Cache<string, number>(3);
cache.set("a", 1);
cache.set("b", 2);
cache.set("c", 3);
cache.display();

cache.get("a"); // Move 'a' to end
cache.set("d", 4); // Should evict 'b'
cache.display();
```

## Common Use Cases

- **Process scheduling**: Operating system task management
- **BFS traversal**: Graph and tree traversal algorithms
- **Print queue**: Managing print jobs in order
- **Buffer management**: Streaming data, keyboard buffer
- **Web server requests**: Handling HTTP requests in order
- **Cache implementation**: LRU cache eviction policies
- **Simulation**: Modeling real-world queuing systems
- **Asynchronous programming**: Task queues in event loops

## Advantages

- Fair processing (FIFO order)
- Efficient insertion and deletion
- Natural modeling of real-world scenarios
- Good for managing resources and scheduling
- Simple and intuitive behavior

## Disadvantages

- No random access to elements
- Limited access (only front and rear)
- Search operations require dequeuing elements
- Memory overhead for linked implementations
- Potential for unbounded growth without size limits