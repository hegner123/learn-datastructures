# Sliding Window Maximum - Complete Beginner's Guide

## What is This Problem About?

Imagine you're looking through a window that can only show you 3 numbers at a time from a long list of numbers. You want to slide this window from left to right across the entire list and find the biggest number visible through the window at each position.

For example, if you have numbers `[1, 3, -1, -3, 5, 3, 6, 7]` and your window shows 3 numbers:
- Position 1: You see `[1, 3, -1]` → biggest is `3`
- Position 2: You see `[3, -1, -3]` → biggest is `3`  
- Position 3: You see `[-1, -3, 5]` → biggest is `5`
- And so on...

## Why Can't We Just Check Each Window Simply?

**The Naive Approach** would be to look at each window and check every number in it to find the maximum. This works but is slow:

```
For each window position:
    Look at all k numbers in the window
    Find the maximum among them
    Add it to our result
```

**Time taken**: If we have n numbers and window size k, this takes O(n × k) time.
**Problem**: For large arrays, this becomes very slow!

## Key Concepts Explained Simply

### 1. Sliding Window
**What it is**: Think of a magnifying glass that can only show a fixed number of items at once. As you move it across a row of objects, different objects come into view while others go out of view.

```
Array: [1, 3, -1, -3, 5, 3, 6, 7], window size = 3

Position 1: [1, 3, -1] _ _ _ _ _
Position 2: _ [3, -1, -3] _ _ _ _
Position 3: _ _ [-1, -3, 5] _ _ _
Position 4: _ _ _ [-3, 5, 3] _ _
Position 5: _ _ _ _ [5, 3, 6] _
Position 6: _ _ _ _ _ [3, 6, 7]
```

### 2. What is a Deque?
**Deque** stands for "Double-Ended Queue" (pronounced "deck").

**Think of it like**: A line of people where you can add or remove people from either the front OR the back of the line.

```
Normal Queue: Only add to back, only remove from front
    Front [Person1 Person2 Person3] Back
          ↑                        ↑
       Remove                    Add

Deque: Can add/remove from BOTH ends
    Front [Person1 Person2 Person3] Back
          ↕                        ↕
      Add/Remove              Add/Remove
```

**Operations**:
- `push_front()`: Add to the front
- `push_back()`: Add to the back  
- `pop_front()`: Remove from the front
- `pop_back()`: Remove from the back

### 3. Monotonic Decreasing
**What it means**: We arrange things in decreasing order (biggest to smallest).

**In our deque**: If we have indices `[i, j, k]`, then the values they point to satisfy:
`nums[i] ≥ nums[j] ≥ nums[k]` (biggest first, smallest last)

**Why this helps**: The front always has the biggest element!

### 4. Amortized Complexity
**Simple explanation**: Even though some individual operations might be expensive, when you average out the cost over many operations, it becomes cheap.

**Analogy**: Imagine you're moving houses. Some days you pack many boxes (expensive), some days you pack few boxes (cheap). But over the entire month, the average work per day is reasonable.

## The Smart Solution: Step-by-Step

### The Big Idea
Instead of recalculating the maximum for each window, we'll keep track of "potential maximums" using a deque. We'll be smart about:
1. **Removing outdated candidates** (numbers that left our window)
2. **Removing useless candidates** (smaller numbers that can never be the maximum while bigger numbers exist)

### Why Store Indices Instead of Values?
We store the **position** (index) of numbers, not the numbers themselves, because:
- We need to know **when** a number goes out of the window
- We can always get the actual value using `nums[index]`

### Detailed Algorithm Walkthrough

Let's trace through `[1, 3, -1, -3, 5, 3, 6, 7]` with window size `k=3`:

---

## Step 1: i=0, nums[0]=1

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | **1** | 3 | -1 | -3 | 5 | 3 | 6 | 7 |

**Window (size=3, not complete yet):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 1 | ? | ? |

**Deque (stores indices):**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 0 | ← remove | ← add here |

Action: `push_back(0)`

---

## Step 2: i=1, nums[1]=3

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | **1** | **3** | -1 | -3 | 5 | 3 | 6 | 7 |

**Window (size=3, not complete yet):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 1 | 3 | ? |

**Deque BEFORE processing:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 0 | ← remove | ← add here |

Actions:
1. `pop_back()` - remove index 0 because nums[1]=3 > nums[0]=1
2. `push_back(1)` - add current index

**Deque AFTER:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 1 | ← remove | ← add here |

---

## Step 3: i=2, nums[2]=-1

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | **1** | **3** | **-1** | -3 | 5 | 3 | 6 | 7 |

**Window (first complete window!):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 1 | 3 | -1 |

**Deque (stores indices):**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 1 | 2 ← remove | ← add here |

Action: `push_back(2)` - nums[2]=-1 < nums[1]=3, so keep monotonic order

**Maximum for this window: 3**

---

## Step 4: i=3, nums[3]=-3

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | 1 | **3** | **-1** | **-3** | 5 | 3 | 6 | 7 |

**Window (size=3):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 3 | -1 | -3 |

**Deque (stores indices):**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 1 | 2 | 3 ← remove | ← add here |

Actions:
1. Check window bounds - index 1 still valid (1 >= 3-3+1=1)
2. `push_back(3)` - nums[3]=-3 < nums[2]=-1, maintain order

**Maximum for this window: 3**

---

## Step 5: i=4, nums[4]=5 (The Interesting Case!)

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | 1 | 3 | **-1** | **-3** | **5** | 3 | 6 | 7 |

**Window (size=3):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| -1 | -3 | 5 |

**Deque BEFORE processing:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 1 | 2 | 3 ← remove | ← add here |

Actions:
1. `pop_front()` - remove index 1 (outside window: 1 < 4-3+1=2)
2. `pop_back()` - remove index 2 because nums[4]=5 > nums[2]=-1
3. `pop_back()` - remove index 3 because nums[4]=5 > nums[3]=-3
4. `push_back(4)` - add current index

**Deque AFTER:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 4 | ← remove | ← add here |

**Maximum for this window: 5**

---

## Step 6: i=5, nums[5]=3

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | 1 | 3 | -1 | **-3** | **5** | **3** | 6 | 7 |

**Window (size=3):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| -3 | 5 | 3 |

**Deque (stores indices):**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 4 | 5 ← remove | ← add here |

Action: `push_back(5)` - nums[5]=3 < nums[4]=5, maintain order

**Maximum for this window: 5**

---

## Step 7: i=6, nums[6]=6

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | 1 | 3 | -1 | -3 | **5** | **3** | **6** | 7 |

**Window (size=3):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 5 | 3 | 6 |

**Deque BEFORE processing:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 4 | 5 ← remove | ← add here |

Actions:
1. `pop_back()` - remove index 4 because nums[6]=6 > nums[4]=5
2. `pop_back()` - remove index 5 because nums[6]=6 > nums[5]=3
3. `push_back(6)` - add current index

**Deque AFTER:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 6 | ← remove | ← add here |

**Maximum for this window: 6**

---

## Step 8: i=7, nums[7]=7

**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|-------|---|---|---|---|---|---|---|---|
| Value | 1 | 3 | -1 | -3 | 5 | **3** | **6** | **7** |

**Window (size=3):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 3 | 6 | 7 |

**Deque BEFORE processing:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 6 | ← remove | ← add here |

Actions:
1. `pop_back()` - remove index 6 because nums[7]=7 > nums[6]=6
2. `push_back(7)` - add current index

**Deque AFTER:**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 7 | ← remove | ← add here |

**Maximum for this window: 7**

## The Four Key Operations (In Order)

### 1. Remove Outdated Indices
```
while deque is not empty AND front index is outside current window:
    remove from front
```
**Purpose**: Get rid of indices that are no longer in our current window.

### 2. Maintain Monotonic Property  
```
while deque is not empty AND value at back index ≤ current value:
    remove from back
```
**Purpose**: Remove smaller elements that can never be the maximum while the current (larger) element exists.

### 3. Add Current Index
```
add current index to back of deque
```
**Purpose**: Every element is a potential future maximum.

### 4. Extract Maximum (if window is complete)
```
if we have processed at least k elements:
    the maximum is the value at the front index
```
**Purpose**: Get the answer for the current window.

## Why This Algorithm is Brilliant

### 1. Each Element is Processed At Most Twice
- Once when added to the deque
- Once when removed from the deque
- Total operations: O(2n) = O(n)

### 2. No Redundant Work
- We never recalculate maximums from scratch
- We only keep elements that could potentially be maximums
- We discard elements as soon as we know they're useless

### 3. Always Correct
- The monotonic property ensures the front is always the maximum
- We remove outdated elements before checking each window
- Every potential maximum candidate is considered

## Common Beginner Mistakes

### 1. Storing Values Instead of Indices
**Wrong**: Store the actual numbers in the deque
**Right**: Store the positions (indices) of the numbers
**Why**: We need to know when numbers go out of the window

### 2. Not Understanding the Monotonic Property
**Confusion**: "Why remove smaller elements?"
**Answer**: If we have a bigger element that comes later, the smaller element can never be the maximum while the bigger one is still in the window.

### 3. Incorrect Window Boundary Checking
**Wrong**: `if (i >= k)`
**Right**: `if (i >= k - 1)`
**Why**: We start getting complete windows at index `k-1` (0-indexed)

## Practice Exercise

Try tracing through this example yourself:
- Array: `[4, 1, 3, 2, 5]`
- Window size: `k = 2`
- Expected result: `[4, 3, 3, 5]`

For each step, track:
1. What's the current window?
2. What indices are in the deque?
3. What values do those indices represent?
4. What operations do you perform?
5. What's the maximum for this window?

## Time and Space Complexity

**Time Complexity**: O(n)
- Each of the n elements is added to the deque exactly once
- Each of the n elements is removed from the deque at most once
- Total operations: O(2n) = O(n)

**Space Complexity**: O(k)
- In the worst case, the deque can contain at most k elements
- This happens when the array is in increasing order

## Real-World Applications

This technique appears in many practical scenarios:
- **Stock prices**: Find the highest price in any rolling time window
- **System monitoring**: Track maximum CPU usage over sliding time periods  
- **Gaming**: Find the strongest enemy in your current view range
- **Data analysis**: Calculate rolling maximums in time series data