import { maxSlidingWindow } from './solution';

describe('maxSlidingWindow', () => {
    test('Basic case', () => {
        const nums = [1, 3, -1, -3, 5, 3, 6, 7];
        const k = 3;
        const expected = [3, 3, 5, 5, 6, 7];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Single element', () => {
        const nums = [1];
        const k = 1;
        const expected = [1];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Window size equals array length', () => {
        const nums = [1, -1, 0, 5, 4];
        const k = 5;
        const expected = [5];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('All same elements', () => {
        const nums = [4, 4, 4, 4];
        const k = 2;
        const expected = [4, 4, 4];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Decreasing array', () => {
        const nums = [9, 8, 7, 6, 5];
        const k = 3;
        const expected = [9, 8, 7];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Increasing array', () => {
        const nums = [1, 2, 3, 4, 5];
        const k = 3;
        const expected = [3, 4, 5];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Negative numbers', () => {
        const nums = [-7, -8, 7, 5, 7, 1, 6, 0];
        const k = 4;
        const expected = [7, 7, 7, 7, 7];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Large window size', () => {
        const nums = [1, 2, 3, 4];
        const k = 3;
        const expected = [3, 4];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Mixed positive negative', () => {
        const nums = [-1, 2, -3, 4, -5, 6];
        const k = 2;
        const expected = [2, 2, 4, 4, 6];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });

    test('Peak in middle', () => {
        const nums = [1, 3, 1, 2, 0, 5];
        const k = 3;
        const expected = [3, 3, 2, 5];
        expect(maxSlidingWindow(nums, k)).toEqual(expected);
    });
});