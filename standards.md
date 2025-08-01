# Standards

*Note: The `00-Reference` directory contains reference implementations and examples demonstrating this format. It is not part of the numbered learning sequence.*

## Format
When generating a problem the output should reflect this format:

**Directory Structure:**
*prefix-subjectmatter*
├─ yyyymmdd-level/
   ├─ question.md
   ├─ .git/
   ├─ .gitignore
   ├─ go/
      ├─ go.mod
      ├─ go.sum
      ├─ solution.go
      ├─ solution_test.go
      └─ Makefile
   └─ ts/
      ├─ package.json
      ├─ tsconfig.json
      ├─ jest.config.js
      ├─ solution.ts
      ├─ solution.test.ts
      └─ test-cases.json

## Levels
Problems are categorized into five difficulty levels:

* **Intro** - Direct replication of subject matter, minimal problem solving
* **Beginner** - Direct replication of subject matter, minimal problem solving
* **Intermediate** - Subject matter + one additional data structure from previous directories
* **Advanced** - Subject matter + up to three additional data structures from previous directories  
* **Senior** - Can use any/all data structures, requires advanced problem solving

*See levels.md for detailed requirements per level.*

## Git Repository Initialization

Each problem directory must be initialized as a Git repository at the same level as `question.md`:

* Run `git init` in the problem directory (yyyymmdd-level/)
* Create a `.gitignore` file with language-specific best practices:
  * **Go**: Include `*.log`, `*.out`, `*.exe`, `*.test`, `vendor/`, `.DS_Store`
  * **TypeScript**: Include `node_modules/`, `dist/`, `build/`, `*.log`, `.DS_Store`, `coverage/`
  * **General**: Include editor/IDE files (`.vscode/`, `.idea/`, `*.swp`, `*.swo`)

## Content

* question.md - includes an overview of the problem, example inputs and outputs, and constraints
  * Must provide **10 different test cases** for the solution to be tested against
  * Test case complexity should match the difficulty level (see levels.md)
* go/ - a directory containing files for a go solution
    * go.mod - module definition and dependency management (use installed Go version)
    * go.sum - dependency checksums (auto-generated)
    * solution.go - file that will contain the user's solution (placeholder with named function signature and proper Go types - no implementation provided)
    * solution_test.go - contains test functions using Go's testing framework
    * Makefile - .PHONY content such that running `make test` runs go test, running `make` shows available commands
* ts/ - a directory containing files for a TypeScript solution
    * package.json - contains scripts for running tests and dependencies (Jest, TypeScript, etc.) - **DO NOT include Jest configuration**
    * tsconfig.json - TypeScript compiler configuration
    * jest.config.js - Jest testing framework configuration (standalone config file only)
    * solution.ts - file that will contain the user's solution in TypeScript (placeholder with named function signature and proper TypeScript types - no implementation provided)
    * solution.test.ts - contains test functions using Jest testing framework
    * test-cases.json - JSON file containing test case data for reuse

## Definitions
* **Subject Matter** - the name of the data structure that matches the current directory name
* **Prefix** - integer reference used to determine which previous data structures can be used in higher levels
