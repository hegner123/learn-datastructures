# Bug Report: TypeScript Jest Test Import Issues

## Problem Description

The most recent TypeScript Jest test file (`20250724-beginner/ts/solution.test.ts`) had two import-related issues preventing tests from running:

1. **JSON Module Import Error**: TypeScript compiler could not resolve JSON imports
2. **Jest test.each() Syntax Error**: Incompatible parameter types when using `test.each()` with destructured parameters

## Error Messages

```
TS2732: Cannot find module './test-cases.json'. Consider using '--resolveJsonModule' to import module with '.json' extension.

TS2345: Argument of type '({ input, expected }: any) => void' is not assignable to parameter of type '(...args: any[] | [any]) => any'.
```

## Root Cause

1. **Missing TypeScript Configuration**: The `tsconfig.json` was missing the `resolveJsonModule: true` compiler option
2. **Inconsistent Test Syntax**: Used `test.each()` syntax instead of the established `forEach()` + `it()` pattern

## Solution Applied

### 1. Fixed TypeScript Configuration
**File**: `20250724-beginner/ts/tsconfig.json`

Added missing compiler option:
```json
"resolveJsonModule": true
```

Removed unnecessary options:
- `declaration: true`
- `declarationMap: true` 
- `sourceMap: true`

### 2. Standardized Test Syntax
**File**: `20250724-beginner/ts/solution.test.ts`

Changed from:
```typescript
test.each(testCases)('$name', ({ input, expected }) => {
  expect(findMinElement(input)).toBe(expected);
});
```

To:
```typescript
testCases.forEach((testCase) => {
  it(testCase.name, () => {
    const result = findMinElement(testCase.input);
    expect(result).toBe(testCase.expected);
  });
});
```

## Prevention

- Ensure all TypeScript projects include `resolveJsonModule: true` when importing JSON files
- Use consistent test patterns across the codebase (follow the `forEach() + it()` pattern established in earlier tests)
- Copy configuration from working examples rather than creating from scratch

## Status

✅ **RESOLVED** - Tests now run successfully without import errors.