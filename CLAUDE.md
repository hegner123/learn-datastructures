# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a data structures learning repository organized as a comprehensive study guide for developers. It contains 33 different data structures and algorithms, each in its own numbered directory (01-Arrays through 33-bit-manipulation).

## Learning Path Purpose

The purpose of this learning path is to achieve **mastery** of these data structures. Mastery includes:

- **Memory Recall**: Implementing any data structure from memory without reference
- **Implementation Fluency**: Determining when full or partial implementation solves a problem
- **Pattern Recognition**: Reading a problem and instinctively knowing what data structure, algorithm, or combination is best to solve it
- **Optimization Intuition**: Recognizing trade-offs and selecting optimal approaches for given constraints
- **Complexity Analysis**: Instantly recognizing time/space complexity implications and proving complexity bounds, not just memorizing them
- **Edge Case Intuition**: Immediately identifying potential failure points, boundary conditions, and corner cases without systematic analysis
- **Implementation Variants**: Understanding when to use different implementations of the same structure (array-based vs linked-list stacks, collision strategies)
- **Failure Mode Recognition**: Knowing not just what works, but what fails and why - understanding the breaking points of each approach
- **Cross-Domain Application**: Recognizing data structure patterns in non-obvious domains (UI state, database design, distributed systems)
- **Teaching Ability**: Explain the "why" behind structure selection to others, indicating deep conceptual understanding
- **Real-World Constraints**: Understanding how theoretical optimality changes with practical constraints (memory hierarchies, concurrency, distributed systems)
- **Anti-Pattern Recognition**: Knowing when NOT to use a structure, recognizing overengineering or inappropriate applications

The most important objective is developing the intuitive ability to analyze any problem and immediately recognize the most effective solution approach, whether it requires a single data structure, multiple structures, specific algorithms, or hybrid approaches.

## Repository Structure

The repository follows a strict organizational pattern:

- **Numbered directories**: Each data structure has a prefix number (01-33) indicating its order in the learning sequence
- **Documentation files**: Each directory contains a markdown file explaining the data structure (e.g., `arrays.md`, `linked_lists.md`)
- **Problem structure**: Problems are organized by date and difficulty level according to `standards.md`
- **Multi-language support**: Go and TypeScript implementations are supported
- **Reference directory**: `00-Reference` contains reference implementations and examples - it is not part of the learning sequence

## Key Files

- `README.md`: Main overview with list of all data structures and practice counts
- `standards.md`: Defines the format and structure for problem organization
- `levels.md`: Defines detailed difficulty level requirements (Beginner, Intermediate, Advanced, Senior)

## Problem Creation Process

### Prerequisites
Before creating a new problem:
1. **Choose target data structure directory** (e.g., `01-Arrays`, `02-linked-lists`, etc.)
2. **Determine difficulty level** (Beginner, Intermediate, Advanced, Senior)
3. **Select date** for problem creation (format: YYYYMMDD)
4. **Review existing problems** in the target directory to ensure uniqueness

### Step 1: Verify Subject-Matter Documentation
Check if the subject-matter directory contains a self-titled markdown file explaining the subject-matter (e.g., `arrays.md` in `01-Arrays/`, `linked_lists.md` in `02-linked-lists/`, etc.).

**If the documentation file exists**: Move to the next step.

**If the documentation file does NOT exist**: Create a self-titled markdown file explaining the subject matter in accordance with the template and guidelines provided in `explain.md`. The file should include all required sections: definition, key properties, time complexity, Go implementation, TypeScript implementation, common use cases, advantages, and disadvantages.

### Step 2: Review Existing Problems for Uniqueness
Before generating a new problem, review all existing problems in the target subject-matter directory to ensure the new question will not be a duplicate.

**Review Process:**
1. List all existing problem directories in the subject-matter folder
2. Read the `question.md` file in each existing problem directory
3. Analyze problem statements, requirements, and core concepts
4. Identify the specific aspects of the data structure being tested
5. Ensure the new problem concept is unique and doesn't overlap

**Uniqueness Criteria:**
- **Problem Statement**: Core challenge must be different from existing problems
- **Implementation Focus**: Different methods or aspects of using the data structure
- **Use Case Scenario**: Different real-world application contexts
- **Algorithmic Approach**: Different problem-solving strategies required
- **Edge Case Focus**: Different boundary conditions or constraints

**If Duplication Detected**: Modify the problem concept to focus on a different aspect of the subject matter or combine with different supporting data structures as allowed by the difficulty level.

## Problem Organization (from standards.md)

When problems are created, they follow this structure:
```
prefix-subjectmatter/
├─ yyyymmdd-level/
   ├─ question.md
   ├─ go/
   │  ├─ go.mod
   │  ├─ go.sum
   │  ├─ solution.go
   │  ├─ solution_test.go
   │  └─ Makefile
   └─ ts/
      ├─ package.json
      ├─ tsconfig.json
      ├─ jest.config.js
      ├─ solution.ts
      ├─ solution.test.ts
      └─ test-cases.json
```

## Development Commands

Commands for working with problems:

- **Go problems**: `make test` (runs `go test` using Go's testing framework)
  - Always use the installed Go version when creating go.mod files
  - Makefile should include `.PHONY` targets: `help`, `test`, `build`, `clean`
  - `make` should show available commands
- **TypeScript problems**: Use npm scripts (typically `npm test`) defined in `package.json` to run Jest tests
  - Run `npm install` after creating package.json to install dependencies

## Difficulty Levels

Problems are categorized into four levels:

1. **Beginner**: Direct replication of subject matter, minimal problem solving
2. **Intermediate**: Subject matter + one additional data structure from previous directories
3. **Advanced**: Subject matter + up to three additional data structures from previous directories  
4. **Senior**: Can use any/all data structures, requires advanced problem solving

## Learning Sequence

Data structures are ordered by complexity and dependency:
- Basic structures (Arrays, linked-lists, Stacks, Queues) come first
- Tree structures follow (binary-trees, BST, avl-trees, red-black-trees, b-trees)
- Graph representations and hash-based structures
- Advanced structures (Heaps, Tries, segment-trees, etc.)
- Specialized structures (lru-cache, design-patterns, etc.)

## Test Requirements

Each problem must include exactly 10 different test cases for the solution to be tested against, with test case complexity matching the difficulty level.

## Implementation Policy

**No solutions are provided** in the codebase. All `solution.go` and `solution.ts` files contain only placeholder comments with TODO instructions. This is intentional to encourage learning through implementation rather than copying existing solutions.

## Difficulty Validation Process

Before finalizing any problem, validate that the difficulty level matches the actual complexity:

### Validation Steps
1. **Review Problem Requirements**
   - Analyze the core problem statement
   - Identify all data structures needed for solution
   - Evaluate algorithmic complexity required

2. **Check Difficulty Alignment**
   - **Beginner**: Uses only the current subject matter, no additional data structures
   - **Intermediate**: Uses current subject + exactly one previous data structure (prefix 1 to current-1)
   - **Advanced**: Uses current subject + up to three previous data structures (prefix 1 to current-1)
   - **Senior**: May use any/all data structures, requires advanced problem solving

3. **Validate Test Cases**
   - **Beginner**: Straightforward cases, no deliberate edge case traps
   - **Intermediate**: Include cases that test robustness, moderate complexity
   - **Advanced**: Complex edge cases, multiple failure scenarios
   - **Senior**: Comprehensive edge cases, real-world complexity

4. **Adjust if Necessary**
   - If problem is too simple for chosen level: Add complexity or lower difficulty
   - If problem is too complex for chosen level: Simplify or raise difficulty
   - If using wrong data structures: Adjust to match level constraints

### Validation Checklist
- [ ] Problem complexity matches chosen difficulty level
- [ ] Required data structures align with difficulty constraints
- [ ] Test cases match complexity expectations for the level
- [ ] Problem statement clearly indicates expected difficulty

## Final Validation Checklist

Before completion, verify:
- [ ] Subject-matter documentation exists (if not, created according to explain.md)
- [ ] Directory structure matches standards.md format
- [ ] Git repository initialized with appropriate .gitignore
- [ ] Go module uses correct Go version (currently 1.24)
- [ ] All 10 test cases implemented in both Go and TypeScript
- [ ] Test complexity matches difficulty level
- [ ] **Problem difficulty has been validated and confirmed**
- [ ] Solution files contain only placeholders (no implementations)
- [ ] question.md includes all required sections
- [ ] Makefile and package.json configured correctly

## Post-Creation Steps

### Update Practice Count
- Navigate to the root repository directory
- Open `README.md`
- Find the Practice Counts section
- Increment the count for the target data structure by 1
- Save the file

### Commit Changes to Git
- Navigate to the problem directory (yyyymmdd-level)
- Add all files to git: `git add .`
- Create initial commit: `git commit -m "Initial problem setup for [data-structure] [difficulty] level"`

### Final Verification
- Review all files for consistency
- Ensure test cases can be run successfully
- Verify placeholder solutions compile but don't implement functionality

## Problem Creation Guidelines

**Dependency Clarity**: When creating problems for advanced structures (suffix-trees, segment-trees, fenwick-trees, sparse-tables, etc.), explicitly identify and document prerequisite data structures needed for the solution. This helps ensure difficulty levels accurately reflect the complexity and knowledge requirements.

**Todo List Validation**: When Claude generates a todo list during problem creation, it must compare the generated todo items against the step-by-step process outlined in `steps.md` and the requirements specified in this `CLAUDE.md` file. This ensures all necessary steps are included and no critical tasks are missed during the problem creation workflow.
