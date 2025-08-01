# Deque Algorithm Visualization Format

This document describes the standardized table format for visualizing deque-based sliding window algorithms step-by-step.

## Format Structure

Each step should contain three main table groups:

### 1. Array State Table
Shows the current position in the array with bold formatting for elements in the current window.

```markdown
**Array State:**
| Index | 0 | 1 | 2 | 3 | 4 | 5 |
|-------|---|---|---|---|---|---|
| Value | **8** | **1** | **3** | 2 | 6 | 4 |
```

**Rules:**
- Use `**bold**` formatting for elements currently in the sliding window
- Include all array indices and values
- Keep consistent column width

### 2. Window Table
Shows only the elements currently visible in the sliding window.

```markdown
**Window (size=3, not complete yet):**
| Position 1 | Position 2 | Position 3 |
|------------|------------|------------|
| 8 | 1 | ? |
```

**Rules:**
- Use `?` for positions not yet filled (incomplete windows)
- Label with window size information
- Note if window is complete or incomplete

### 3. Deque State Table
Shows the deque contents with function availability headers.

```markdown
**Deque (stores indices):**
| pop_front() | front() | pop_back() | push_back() |
|-------------|---------|------------|-------------|
| ← remove | 0 | 1 ← remove | ← add here |
```

**Rules:**
- Headers indicate which functions are available at each position
- Use arrows (`←`) to show direction of operations
- Include descriptive text: "remove", "add here"
- Show actual indices stored, not values
- Note in parentheses that deque stores indices

## Step Organization

### Step Header Format
```markdown
## Step X: i=Y, nums[Y]=Z
```
Where:
- X = step number
- Y = current array index
- Z = current array value

### Action Documentation
After each set of tables, include:

```markdown
Action: `function_name(parameter)`
```

For multiple actions:
```markdown
Actions:
1. `function_name(parameter)` - explanation
2. `function_name(parameter)` - explanation
```

### Before/After States
For complex steps with multiple operations:

```markdown
**Deque BEFORE processing:**
[table showing initial state]

Actions: [list of operations]

**Deque AFTER:**
[table showing final state]
```

## Special Cases

### Incomplete Windows
For steps where window isn't complete yet:
```markdown
**Window (size=3, not complete yet):**
```

### First Complete Window
For the first step with a complete window:
```markdown
**Window (first complete window!):**
```

### Maximum Extraction
When extracting the maximum for a complete window:
```markdown
**Maximum for this window: X**
```

## Table Spacing and Separators

Use horizontal rules (`---`) between steps:

```markdown
---

## Step 2: i=1, nums[1]=1
```

## Formatting Consistency

- **Bold** for current window elements in array
- `Code formatting` for function names and parameters
- **Bold** for table headers and important notes
- Consistent column alignment using proper markdown table syntax
- Clear visual hierarchy with appropriate header levels

## Example Usage

This format is ideal for:
- Deque-based sliding window algorithms
- Monotonic deque explanations  
- Step-by-step algorithm walkthroughs
- Educational content requiring detailed state visualization

The format provides clear visual separation between data structures while showing their relationships and the available operations at each step.