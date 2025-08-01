import { reverseArray } from './solution';
import { testCases } from './test-cases.json';

describe('reverseArray', () => {
    test('Empty array', () => {
        const input = [];
        const expected = [];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Single element', () => {
        const input = [42];
        const expected = [42];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Two elements', () => {
        const input = [1, 2];
        const expected = [2, 1];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Odd length array', () => {
        const input = [1, 2, 3, 4, 5];
        const expected = [5, 4, 3, 2, 1];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Even length array', () => {
        const input = [10, 20, 30, 40];
        const expected = [40, 30, 20, 10];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Array with negative numbers', () => {
        const input = [-1, -2, -3, -4];
        const expected = [-4, -3, -2, -1];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Array with mixed positive and negative', () => {
        const input = [-5, 10, -3, 8, 2];
        const expected = [2, 8, -3, 10, -5];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Array with zeros', () => {
        const input = [0, 0, 1, 0, 0];
        const expected = [0, 0, 1, 0, 0];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Array with duplicate values', () => {
        const input = [7, 7, 7, 7];
        const expected = [7, 7, 7, 7];
        expect(reverseArray([...input])).toEqual(expected);
    });

    test('Large numbers', () => {
        const input = [1000, 2000, 3000];
        const expected = [3000, 2000, 1000];
        expect(reverseArray([...input])).toEqual(expected);
    });

    // Test with data from JSON file
    testCases.forEach((testCase, index) => {
        test(`Test case ${index + 1}: ${testCase.name}`, () => {
            expect(reverseArray([...testCase.input])).toEqual(testCase.expected);
        });
    });
});
