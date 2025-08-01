import { Calculator } from './solution';
import testCases from './test-cases.json';

describe('Calculator', () => {
  let calc: Calculator;

  beforeEach(() => {
    calc = new Calculator();
  });

  testCases.forEach((testCase) => {
    it(testCase.name, () => {
      const [a, b] = testCase.input;
      
      if (testCase.expected === "error") {
        expect(() => {
          switch (testCase.operation) {
            case 'add':
              calc.add(a, b);
              break;
            case 'subtract':
              calc.subtract(a, b);
              break;
            case 'multiply':
              calc.multiply(a, b);
              break;
            case 'divide':
              calc.divide(a, b);
              break;
          }
        }).toThrow();
      } else {
        let result: number;
        switch (testCase.operation) {
          case 'add':
            result = calc.add(a, b);
            break;
          case 'subtract':
            result = calc.subtract(a, b);
            break;
          case 'multiply':
            result = calc.multiply(a, b);
            break;
          case 'divide':
            result = calc.divide(a, b);
            break;
          default:
            throw new Error(`Unknown operation: ${testCase.operation}`);
        }
        expect(result).toBe(testCase.expected);
      }
    });
  });
});