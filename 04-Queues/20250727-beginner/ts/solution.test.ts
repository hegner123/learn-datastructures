import { Queue } from './solution';

describe('Queue', () => {
  test('should create a new empty queue', () => {
    const queue = new Queue<string>();
    expect(queue.isEmpty()).toBe(true);
    expect(queue.size()).toBe(0);
  });

  test('should enqueue a single item', () => {
    const queue = new Queue<string>();
    queue.enqueue('hello');
    expect(queue.isEmpty()).toBe(false);
    expect(queue.size()).toBe(1);
    expect(queue.front()).toBe('hello');
  });

  test('should enqueue multiple items', () => {
    const queue = new Queue<number>();
    queue.enqueue(1);
    queue.enqueue(2);
    queue.enqueue(3);
    expect(queue.size()).toBe(3);
    expect(queue.front()).toBe(1);
  });

  test('should dequeue from non-empty queue', () => {
    const queue = new Queue<string>();
    queue.enqueue('first');
    queue.enqueue('second');
    const item = queue.dequeue();
    expect(item).toBe('first');
    expect(queue.size()).toBe(1);
    expect(queue.front()).toBe('second');
  });

  test('should dequeue from empty queue', () => {
    const queue = new Queue<string>();
    const item = queue.dequeue();
    expect(item).toBeNull();
  });

  test('should return null for front on empty queue', () => {
    const queue = new Queue<string>();
    const front = queue.front();
    expect(front).toBeNull();
  });

  test('should return true for isEmpty on empty queue', () => {
    const queue = new Queue<string>();
    expect(queue.isEmpty()).toBe(true);
  });

  test('should return false for isEmpty on non-empty queue', () => {
    const queue = new Queue<string>();
    queue.enqueue('item');
    expect(queue.isEmpty()).toBe(false);
  });

  test('should track size correctly', () => {
    const queue = new Queue<number>();
    
    // Size should be 0 initially
    expect(queue.size()).toBe(0);
    
    // Add items and check size
    queue.enqueue(1);
    expect(queue.size()).toBe(1);
    
    queue.enqueue(2);
    expect(queue.size()).toBe(2);
    
    // Remove items and check size
    queue.dequeue();
    expect(queue.size()).toBe(1);
    
    queue.dequeue();
    expect(queue.size()).toBe(0);
  });

  test('should maintain FIFO behavior', () => {
    const queue = new Queue<string>();
    const items = ['first', 'second', 'third', 'fourth'];
    
    // Enqueue all items
    items.forEach(item => queue.enqueue(item));
    
    // Dequeue all items and verify FIFO order
    items.forEach(expected => {
      const actual = queue.dequeue();
      expect(actual).toBe(expected);
    });
    
    // Queue should be empty now
    expect(queue.isEmpty()).toBe(true);
  });
});