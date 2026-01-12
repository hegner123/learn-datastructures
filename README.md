# Data Structures Learning System

An AI-powered workflow for generating comprehensive data structures and algorithms practice problems using Claude Code.

## What is This?

This repository is a **prompt engineering workflow system** designed to work with [Claude Code](https://claude.ai/code) to automatically generate high-quality practice problems across 33 essential data structures and algorithms. Rather than providing pre-made problems, this system empowers you to create custom problems on-demand, tailored to your learning pace and needs.

The goal is to achieve **mastery** through deliberate practice - developing the intuitive ability to analyze any problem and immediately recognize the most effective solution approach.

## Key Features

- **AI-Driven Problem Generation**: Leverages Claude Code to create unique, well-structured problems
- **33 Data Structures**: Comprehensive coverage from arrays to advanced structures like segment trees
- **Progressive Difficulty**: Five levels (Intro, Beginner, Intermediate, Advanced, Senior) with strict constraints
- **Multi-Language Support**: Problems generated with both Go and TypeScript solutions
- **Complete Test Suites**: Each problem includes 10 test cases with proper test harness setup
- **Standardized Format**: Consistent structure across all problems for streamlined learning

## Quick Start

### Prerequisites

- [Claude Code](https://claude.ai/code) CLI installed
- Go 1.24 or later
- Node.js (for TypeScript problems)

### Generate Your First Problem

**Option 1: Using the convenience script**

```bash
# Clone this repository
git clone https://github.com/hegner123/learn-datastructures.git
cd learn-datastructures

# Make the script executable
chmod +x create-problem.sh

# Generate a problem
./create-problem.sh beginner Arrays
./create-problem.sh intermediate linked-lists
./create-problem.sh advanced binary-trees

# Syntax: ./create-problem.sh <difficulty> <data-structure>
# Difficulty: intro, beginner, intermediate, advanced, senior
```

**Option 2: Using Claude Code directly**

```bash
# Navigate to the repository
cd learn-datastructures

# Start Claude Code
claude

# In Claude Code, request a problem
> Create an intermediate-level problem for linked-lists
```

Claude will:
1. Check if subject documentation exists (creates if needed)
2. Review existing problems to ensure uniqueness
3. Generate a problem following STANDARDS.md format
4. Create Go and TypeScript implementations with tests
5. Update practice counts

## Repository Structure

```
learn-datastructures/
├── CLAUDE.md              # AI workflow instructions (the core prompt)
├── STANDARDS.md           # Problem format specifications
├── STEPS.md              # Step-by-step problem creation process
├── LEVELS.md             # Difficulty level definitions
├── EXPLAIN.md            # Template for subject matter documentation
├── 00-Reference/         # Example problem (reference implementation)
├── 01-Arrays/            # Arrays data structure
│   ├── arrays.md        # Subject matter explanation
│   └── [generated problems appear here]
├── 02-linked-lists/      # Linked lists data structure
│   └── linked_lists.md
└── [03-33: Other data structures...]
```

## Data Structures Covered

### Essential Data Structures for Senior Software Developers

**Basic Structures:**
- Arrays (Dynamic/Static)
- linked-lists (Singly, Doubly, Circular)
- Stacks
- Queues (Regular, Priority, Circular, Deque)

**Trees:**
- binary-trees
- binary-search-trees (BST)
- avl-trees
- red-black-trees
- b-trees
- Tries (Prefix Trees)
- segment-trees
- fenwick-trees (Binary Indexed Trees)

**Graphs:**
- adjacency-matrix
- adjacency-list
- directed-graphs/undirected-graphs
- weighted-graphs/unweighted-graphs

**Hash-based:**
- hash-tables/Hash Maps
- hash-sets
- bloom-filters

**Advanced Structures:**
- Heaps (Min/Max, Binary, Fibonacci)
- disjoint-set-union (Union-Find)
- suffix-arrays
- sparse-tables
- skip-lists

**String Structures:**
- suffix-trees
- kmp-algorithm structures
- rolling-hash

**Specialized:**
- lru-cache
- design-patterns (Singleton, Factory, Observer)
- circular-buffers
- bit-manipulation structures

## Difficulty Levels

Problems are categorized into five levels with strict data structure constraints:

- **Intro**: Direct replication of subject matter, minimal problem solving
- **Beginner**: Direct implementation of subject matter concepts
- **Intermediate**: Subject matter + one additional data structure from previous directories
- **Advanced**: Subject matter + up to three additional data structures from previous directories
- **Senior**: Can use any/all data structures, requires advanced problem solving

See `LEVELS.md` for detailed requirements.

## How It Works

This system uses prompt engineering to guide Claude Code through a rigorous problem creation workflow:

1. **Documentation First**: Ensures subject matter explanation exists
2. **Uniqueness Check**: Reviews existing problems to avoid duplication
3. **Structured Generation**: Creates problems following strict format (STANDARDS.md)
4. **Multi-Language**: Generates both Go and TypeScript implementations
5. **Complete Testing**: Includes 10 test cases per problem
6. **Validation**: Verifies difficulty level matches actual complexity

The workflow is defined in `CLAUDE.md` and `STEPS.md` - these files contain the prompts that guide Claude Code.

## Practice Counts

Track your progress as you generate and complete problems:

- 00-Reference: 1 (Example)
- 01-Arrays: 0
- 02-linked-lists: 0
- 03-Stacks: 0
- 04-Queues: 0
- 05-binary-trees: 0
- 06-binary-search-trees: 0
- 07-avl-trees: 0
- 08-red-black-trees: 0
- 09-b-trees: 0
- 10-Tries: 0
- 11-segment-trees: 0
- 12-fenwick-trees: 0
- 13-adjacency-matrix: 0
- 14-adjacency-list: 0
- 15-directed-graphs: 0
- 16-undirected-graphs: 0
- 17-weighted-graphs: 0
- 18-unweighted-graphs: 0
- 19-hash-tables: 0
- 20-hash-sets: 0
- 21-bloom-filters: 0
- 22-Heaps: 0
- 23-disjoint-set-union: 0
- 24-suffix-arrays: 0
- 25-sparse-tables: 0
- 26-skip-lists: 0
- 27-suffix-trees: 0
- 28-kmp-algorithm: 0
- 29-rolling-hash: 0
- 30-lru-cache: 0
- 31-design-patterns: 0
- 32-circular-buffers: 0
- 33-bit-manipulation: 0

**Legend:** i = Intro, B = Beginner, I = Intermediate, A = Advanced, S = Senior

## Key Workflow Files

- **CLAUDE.md**: The main prompt that instructs Claude Code on problem generation workflow
- **STANDARDS.md**: Defines the exact directory structure and file format for problems
- **STEPS.md**: Step-by-step process for creating a problem (used by Claude)
- **LEVELS.md**: Detailed difficulty level requirements and constraints
- **EXPLAIN.md**: Template for creating subject matter documentation
- **practice-counts.json**: Machine-readable practice count tracking

## Running Tests

### Go Problems
```bash
cd [data-structure]/[date]-[level]/go
make test
```

### TypeScript Problems
```bash
cd [data-structure]/[date]-[level]/ts
npm install
npm test
```

## Contributing

This is a personal learning workflow system. Feel free to fork and adapt for your own use. The workflow files (CLAUDE.md, STANDARDS.md, etc.) can be customized to match your learning preferences.

## License

MIT License - See LICENSE file for details

## Reference

Check `00-Reference/20250718-beginner/` to see an example of a generated problem structure.

---

**Note**: This system is designed specifically for use with Claude Code. The markdown files in the root (CLAUDE.md, STANDARDS.md, STEPS.md) are the workflow prompts that guide the AI through problem generation.
