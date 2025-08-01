# Steps for Creating New Problems

This document outlines the chronological process for creating new problems in the data structures repository.

## Prerequisites

1. **Choose target data structure directory** (e.g., `01-Arrays`, `02-linked-lists`, etc.)
2. **Determine difficulty level** (Intro, Beginner, Intermediate, Advanced, Senior)
3. **Select date** for problem creation (format: YYYYMMDD)

## Step-by-Step Process

### 1. Verify Subject-Matter Documentation
Check if the subject-matter directory contains a self-titled markdown file explaining the subject-matter (e.g., `arrays.md` in `01-Arrays/`, `linked_lists.md` in `02-linked-lists/`, etc.).

**If the documentation file exists**: Move to the next step.

**If the documentation file does NOT exist**: Create a self-titled markdown file explaining the subject matter in accordance with the template and guidelines provided in `explain.md`. The file should include all required sections: definition, key properties, time complexity, Go implementation, TypeScript implementation, common use cases, advantages, and disadvantages.

### 2. Review Existing Problems
Before creating a new problem, review all existing problems in the target subject-matter directory to ensure the new question will not be a duplicate.

**Process:**
1. List all existing problem directories in the subject-matter folder
2. Read the `question.md` file in each existing problem directory
3. Analyze the problem statements, requirements, and core concepts
4. Ensure the new problem idea is unique and doesn't overlap with existing problems
5. If duplication is detected, modify the problem concept or approach a different aspect of the subject matter

**Duplication Check:**
- **Problem Statement**: The core challenge should be different
- **Implementation Approach**: Different methods of using the data structure
- **Use Case**: Different real-world application scenarios
- **Complexity Focus**: Different aspects of time/space optimization

### 3. Create Problem Directory Structure
```bash
mkdir "yyyymmdd-level"
cd "yyyymmdd-level"
```

### 4. Create .gitignore File
Create `.gitignore` with appropriate exclusions:
- **Go**: `*.log`, `*.out`, `*.exe`, `*.test`, `vendor/`, `.DS_Store`
- **TypeScript**: `node_modules/`, `dist/`, `build/`, `*.log`, `.DS_Store`, `coverage/`
- **General**: `.vscode/`, `.idea/`, `*.swp`, `*.swo`

### 6. Create Go Implementation Structure
```bash
mkdir go
cd go
```

#### 6.1 Create go.mod
```bash
go mod init [module-name]
```
*Use the installed Go version (currently 1.24)*

#### 6.2 Create solution.go
- Add placeholder file with TODO comments
- **No actual implementation** - learning through doing

#### 6.3 Create solution_test.go
- Implement all 10 test cases using Go's testing framework
- Test complexity must match difficulty level

#### 6.4 Create Makefile
- `.PHONY` targets: `help`, `test`, `build`, `clean`
- `make test` runs `go test -v`
- `make` shows available commands

### 7. Create TypeScript Implementation Structure
```bash
mkdir ts
cd ts
```

#### 7.1a Create package.json
- Include Jest and TypeScript dependencies
- Add test script (`npm test`)

#### 7.1b  install packages
- Run `npm install` to install dependencies


#### 7.1b  check if you skipped step 7.1b
- Did you skip step 7.1b?


#### 7.2 Create tsconfig.json
- TypeScript compiler configuration

#### 7.3 Create jest.config.js
- Jest testing framework configuration

#### 7.4 Create solution.ts
- Add placeholder file with TODO comments
- **No actual implementation** - learning through doing

#### 7.5 Create solution.test.ts
- Implement all 10 test cases using Jest
- Test complexity must match difficulty level

#### 7.6 Create test-cases.json
- JSON file with test case data for reuse between Go and TypeScript

### 8. Create question.md
Write comprehensive problem description including:
- **Problem Description**: Clear explanation of what to implement
- **Requirements**: Specific implementation requirements
- **Example Usage**: Code examples showing expected behavior
- **Constraints**: Input/output constraints and edge cases
- **Test Cases**: Exactly 10 test cases with expected inputs/outputs

### 9. Test Case Requirements by Level

#### Intro
- Direct implementation of subject matter
- Straightforward test cases
- Minimal additional problem solving

#### Beginner
- Direct implementation of subject matter
- Straightforward test cases
- Minimal additional problem solving

#### Intermediate
- Subject matter + one previous data structure
- Include cases that could trip up less robust solutions
- Moderate problem solving required

#### Advanced
- Subject matter + up to three previous data structures
- Include edge cases and tricky scenarios
- Advanced problem solving required

#### Senior
- Can use any/all data structures
- Complex test cases with multiple edge cases
- Advanced problem solving required

### 10. Validate Problem Difficulty

Before finalizing the problem, validate that the difficulty level matches the actual complexity:

#### Difficulty Validation Process
1. **Review Problem Requirements**
   - Analyze the core problem statement
   - Identify all data structures needed for solution
   - Evaluate algorithmic complexity required

2. **Check Difficulty Alignment**
   - **Intro**: Uses only the current subject matter, no additional data structures
   - **Beginner**: Uses only the current subject matter, no additional data structures
   - **Intermediate**: Uses current subject + exactly one previous data structure (prefix 1 to current-1)
   - **Advanced**: Uses current subject + up to three previous data structures (prefix 1 to current-1)
   - **Senior**: May use any/all data structures, requires advanced problem solving

3. **Validate Test Cases**
   - **Intro**: Straightforward cases, no deliberate edge case traps
   - **Beginner**: Straightforward cases, no deliberate edge case traps
   - **Intermediate**: Include cases that test robustness, moderate complexity
   - **Advanced**: Complex edge cases, multiple failure scenarios
   - **Senior**: Comprehensive edge cases, real-world complexity

4. **Adjust if Necessary**
   - If problem is too simple for chosen level: Add complexity or lower difficulty
   - If problem is too complex for chosen level: Simplify or raise difficulty
   - If using wrong data structures: Adjust to match level constraints

#### Validation Checklist
- [ ] Problem complexity matches chosen difficulty level
- [ ] Required data structures align with difficulty constraints
- [ ] Test cases match complexity expectations for the level
- [ ] Problem statement clearly indicates expected difficulty

### 11. Final Validation Checklist

Before completion, verify:
- [ ] Subject-matter documentation exists (if not, created according to explain.md)
- [ ] Directory structure matches standards.md format
- [ ] .gitignore file created with appropriate exclusions
- [ ] Go module uses correct Go version (1.24)
- [ ] All 10 test cases implemented in both Go and TypeScript
- [ ] Test complexity matches difficulty level
- [ ] **Problem difficulty has been validated and confirmed**
- [ ] Solution files contain only placeholders (no implementations)
- [ ] question.md includes all required sections
- [ ] Makefile and package.json configured correctly
- [ ] TypeScript directory has node_modules directory

### 12a. Update Practice Count
- Navigate to the root repository directory
- Open `README.md`
- Find the Practice Counts section
- Increment the count for the target data structure by 1
- Save the file

### 12b. Update Practice Count
- Open practice-counts.json
- Increment the count for the target data structure by 1
- Save the file

### 13. Final Steps
- Review all files for consistency
- Ensure test cases can be run successfully
- Verify placeholder solutions compile but don't implement functionality

### 15. Did you skip step 7.1b?
- Did you skip step 7.1b?

## Notes

- **No solutions provided**: All solution files contain only TODO placeholders
- **Learning through implementation**: Students must implement solutions themselves
- **Test-driven approach**: Complete test suites guide implementation
- **Multi-language support**: Both Go and TypeScript implementations required
- **Proper file organization**: Each problem has proper .gitignore for clean development
