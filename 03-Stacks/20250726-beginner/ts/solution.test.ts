import { isBalanced } from './solution';

describe('isBalanced', () => {
    // Test Case 1: Empty string
    test('should return true for empty string', () => {
        expect(isBalanced('')).toBe(true);
    });

    // Test Case 2: Single pair of parentheses
    test('should return true for single pair of parentheses', () => {
        expect(isBalanced('()')).toBe(true);
    });

    // Test Case 3: Single pair of brackets
    test('should return true for single pair of brackets', () => {
        expect(isBalanced('[]')).toBe(true);
    });

    // Test Case 4: Single pair of braces
    test('should return true for single pair of braces', () => {
        expect(isBalanced('{}')).toBe(true);
    });

    // Test Case 5: Nested brackets of same type
    test('should return true for nested brackets of same type', () => {
        expect(isBalanced('((()))')).toBe(true);
    });

    // Test Case 6: Mixed balanced brackets
    test('should return true for mixed balanced brackets', () => {
        expect(isBalanced('([{}])')).toBe(true);
    });

    // Test Case 7: Unmatched opening bracket
    test('should return false for unmatched opening bracket', () => {
        expect(isBalanced('(()')).toBe(false);
    });

    // Test Case 8: Unmatched closing bracket
    test('should return false for unmatched closing bracket', () => {
        expect(isBalanced('())')).toBe(false);
    });

    // Test Case 9: Wrong order of brackets
    test('should return false for wrong order of brackets', () => {
        expect(isBalanced('([)]')).toBe(false);
    });

    // Test Case 10: Complex nested structure
    test('should return true for complex nested structure', () => {
        expect(isBalanced('{[(){}][(())]}')).toBe(true);
    });
});