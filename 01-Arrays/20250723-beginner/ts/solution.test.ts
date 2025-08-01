import { arrayMax } from './solution';
import testCases from './test-cases.json';

describe('arrayMax', () => {
  testCases.forEach((testCase) => {
    it(testCase.name, () => {
      const result = arrayMax(testCase.input);
      expect(result).toBe(testCase.expected);
    });
  });
});