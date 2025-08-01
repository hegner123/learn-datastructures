import { findMaxElement } from './solution';
import testCases from './test-cases.json';

describe('findMaxElement', () => {
  testCases.forEach((testCase) => {
    it(testCase.name, () => {
      const result = findMaxElement(testCase.input);
      expect(result).toBe(testCase.expected);
    });
  });

  it('should throw error for empty array', () => {
    expect(() => findMaxElement([])).toThrow('Array cannot be empty');
  });
});