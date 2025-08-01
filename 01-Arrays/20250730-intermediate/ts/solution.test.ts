import { mergeSortedArrays } from './solution';
import testCases from './test-cases.json';

describe('mergeSortedArrays', () => {
    test('Basic merge', () => {
        const result = mergeSortedArrays([1, 3, 5], [2, 4, 6]);
        expect(result).toEqual([1, 2, 3, 4, 5, 6]);
    });

    test('Empty first array', () => {
        const result = mergeSortedArrays([], [1, 2, 3]);
        expect(result).toEqual([1, 2, 3]);
    });

    test('Empty second array', () => {
        const result = mergeSortedArrays([1, 2, 3], []);
        expect(result).toEqual([1, 2, 3]);
    });

    test('Both arrays empty', () => {
        const result = mergeSortedArrays([], []);
        expect(result).toEqual([]);
    });

    test('Different lengths', () => {
        const result = mergeSortedArrays([1, 5, 9, 10, 15, 20], [2, 3, 8, 13]);
        expect(result).toEqual([1, 2, 3, 5, 8, 9, 10, 13, 15, 20]);
    });

    test('First array larger elements', () => {
        const result = mergeSortedArrays([4, 5, 6], [1, 2, 3]);
        expect(result).toEqual([1, 2, 3, 4, 5, 6]);
    });

    test('Duplicate elements', () => {
        const result = mergeSortedArrays([1, 3, 5], [1, 3, 5]);
        expect(result).toEqual([1, 1, 3, 3, 5, 5]);
    });

    test('Negative numbers', () => {
        const result = mergeSortedArrays([-5, -1, 2], [-3, 0, 4]);
        expect(result).toEqual([-5, -3, -1, 0, 2, 4]);
    });

    test('Single elements', () => {
        const result = mergeSortedArrays([5], [3]);
        expect(result).toEqual([3, 5]);
    });

    test('Large ranges', () => {
        const result = mergeSortedArrays([100, 200, 300], [150, 250, 350, 400, 500]);
        expect(result).toEqual([100, 150, 200, 250, 300, 350, 400, 500]);
    });
});