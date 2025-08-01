import { findMinElement } from './solution';
import testCases from './test-cases.json';

describe('findMinElement', () => {
  testCases.forEach((testCase) => {
    it(testCase.name, () => {
      const result = findMinElement(testCase.input);
      expect(result).toBe(testCase.expected);
    });
  });
});
