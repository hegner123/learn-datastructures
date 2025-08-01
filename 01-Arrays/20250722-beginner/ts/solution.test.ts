import { arraySum } from './solution';
import testCases from './test-cases.json';

describe('arraySum', () => {
  testCases.forEach((testCase) => {
    it(testCase.name, () => {
      const result = arraySum(testCase.input);
      expect(result).toBe(testCase.expected);
    });
  });
});