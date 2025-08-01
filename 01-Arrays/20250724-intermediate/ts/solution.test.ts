import { twoSum } from './solution';

describe('twoSum', () => {
  test('Basic case', () => {
    const nums = [2, 7, 11, 15];
    const target = 9;
    const expected = [0, 1];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Reverse order', () => {
    const nums = [15, 11, 7, 2];
    const target = 9;
    const expected = [2, 3];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Same elements', () => {
    const nums = [3, 3];
    const target = 6;
    const expected = [0, 1];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Negative numbers', () => {
    const nums = [-1, -2, -3, -4, -5];
    const target = -8;
    const expected = [2, 4];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Mixed signs', () => {
    const nums = [-3, 4, 3, 90];
    const target = 0;
    const expected = [0, 2];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Zero target', () => {
    const nums = [-2, 0, 1, 1];
    const target = 2;
    const expected = [2, 3];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Large numbers', () => {
    const nums = [999, 1, 2];
    const target = 1000;
    const expected = [0, 1];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Adjacent elements', () => {
    const nums = [1, 2, 3, 4];
    const target = 7;
    const expected = [2, 3];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('First and last', () => {
    const nums = [5, 1, 2, 8];
    const target = 13;
    const expected = [0, 3];
    expect(twoSum(nums, target)).toEqual(expected);
  });

  test('Middle elements', () => {
    const nums = [10, 20, 30, 40, 50];
    const target = 70;
    const expected = [2, 3];
    expect(twoSum(nums, target)).toEqual(expected);
  });
});