# Sliding Window Maximum - Monotonic Deque Solution

## Key Concepts

**Sliding Window**: A subarray of fixed size k that "slides" across the array from left to right.
```
Array: [1, 3, -1, -3, 5, 3, 6, 7], k=3
Windows: [1,3,-1] → [3,-1,-3] → [-1,-3,5] → [-3,5,3] → [5,3,6] → [3,6,7]
```

**Deque (Double-ended Queue)**: A data structure allowing insertion/deletion from both front and back.
```
Operations: push_front(), push_back(), pop_front(), pop_back()
```

**Monotonic Decreasing**: Elements stored in decreasing order of their values.
```
If deque contains indices [i, j, k], then nums[i] ≥ nums[j] ≥ nums[k]
```

**Amortized Complexity**: Average time per operation over a sequence of operations.

## Visual Step-by-Step Walkthrough

**Array**: [1, 3, -1, -3, 5, 3, 6, 7], **k=3**

### Step 1: i=0, nums[0]=1
```
Window: [1, ?, ?] (incomplete)
Deque: [0] (stores index 0, value 1)
Action: Add index 0
```

### Step 2: i=1, nums[1]=3
```
Window: [1, 3, ?] (incomplete)
Deque before: [0] (value 1)
Deque after:  [1] (value 3)
Action: Remove index 0 (3 > 1), add index 1
```

### Step 3: i=2, nums[2]=-1
```
Window: [1, 3, -1] ← First complete window
Deque: [1, 2] (values 3, -1)
Maximum: nums[1] = 3
Action: Keep index 1 (-1 < 3), add index 2
```

### Step 4: i=3, nums[3]=-3
```
Window: [3, -1, -3]
Deque: [1, 2, 3] (values 3, -1, -3)
Maximum: nums[1] = 3
Action: Keep all (-3 < -1), add index 3
```

### Step 5: i=4, nums[4]=5
```
Window: [-1, -3, 5]
Deque before: [1, 2, 3] (values 3, -1, -3)
Deque after:  [4] (value 5)
Maximum: nums[4] = 5
Actions: 
  1. Remove index 1 (outside window: 1 ≤ 4-3)
  2. Remove indices 2,3 (5 > -1, 5 > -3)
  3. Add index 4
```

## Deque State Visualization

```
Step | Window      | Deque Indices | Deque Values | Max
-----|-------------|---------------|--------------|----
  1  | [1,?,?]     | [0]           | [1]          | -
  2  | [1,3,?]     | [1]           | [3]          | -
  3  | [1,3,-1]    | [1,2]         | [3,-1]       | 3
  4  | [3,-1,-3]   | [1,2,3]       | [3,-1,-3]    | 3
  5  | [-1,-3,5]   | [4]           | [5]          | 5
  6  | [-3,5,3]    | [4,5]         | [5,3]        | 5
  7  | [5,3,6]     | [6]           | [6]          | 6
  8  | [3,6,7]     | [7]           | [7]          | 7
```

## Algorithm Operations

### Core Operations in Order:

1. **Remove Outdated Indices**: 
   ```
   while (!deque.empty() && deque.front() <= i - k)
       deque.pop_front()
   ```
   Remove indices that fall outside the current window.

2. **Maintain Monotonic Property**:
   ```
   while (!deque.empty() && nums[deque.back()] <= nums[i])
       deque.pop_back()
   ```
   Remove smaller elements from back since they can never be maximum while current element exists.

3. **Add Current Index**:
   ```
   deque.push_back(i)
   ```
   Always add current index to maintain potential maximum candidates.

4. **Extract Maximum**:
   ```
   if (i >= k - 1)
       result.push_back(nums[deque.front()])
   ```
   Front index always gives maximum for current window.

## Why This Algorithm Works

### Efficiency Principles:

- **Amortized O(n)**: Each element enters and exits the deque at most once
- **No Redundant Work**: Only track elements that could potentially be maximums
- **Early Elimination**: Remove elements that can never be the answer

### Correctness Guarantees:

- **Maximum at Front**: Monotonic decreasing order ensures front is always largest
- **Window Validity**: Outdated indices are removed before each window check  
- **Complete Coverage**: Every possible maximum candidate is considered

## Complexity Analysis

- **Time Complexity**: O(n) - each element processed exactly once
- **Space Complexity**: O(k) - deque stores at most k elements in worst case

## Key Insight

The solution efficiently tracks potential maximums while discarding elements that can never be the answer. By maintaining a monotonic decreasing deque of indices, we ensure the front always contains the maximum for any valid window, making this optimal for sliding window maximum problems.
