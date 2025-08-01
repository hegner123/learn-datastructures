import { arrayContains } from './solution';

describe('arrayContains', () => {
    test('Empty array', () => {
        expect(arrayContains([], 5)).toBe(false);
    });

    test('Single element - found', () => {
        expect(arrayContains([42], 42)).toBe(true);
    });

    test('Single element - not found', () => {
        expect(arrayContains([42], 10)).toBe(false);
    });

    test('Multiple elements - target at beginning', () => {
        expect(arrayContains([1, 2, 3, 4, 5], 1)).toBe(true);
    });

    test('Multiple elements - target at end', () => {
        expect(arrayContains([1, 2, 3, 4, 5], 5)).toBe(true);
    });

    test('Multiple elements - target in middle', () => {
        expect(arrayContains([1, 2, 3, 4, 5], 3)).toBe(true);
    });

    test('Multiple elements - target not found', () => {
        expect(arrayContains([1, 2, 3, 4, 5], 10)).toBe(false);
    });

    test('Array with negative numbers - found', () => {
        expect(arrayContains([-5, -2, -8, -1], -2)).toBe(true);
    });

    test('Array with negative numbers - not found', () => {
        expect(arrayContains([-5, -2, -8, -1], 2)).toBe(false);
    });

    test('Array with duplicates - found', () => {
        expect(arrayContains([7, 7, 7, 7], 7)).toBe(true);
    });
});