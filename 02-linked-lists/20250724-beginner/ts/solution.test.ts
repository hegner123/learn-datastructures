import { LinkedList } from './solution';
import testCases from './test-cases.json';

describe('LinkedList', () => {
  testCases.forEach((testCase) => {
    it(testCase.name, () => {
      const ll = new LinkedList<number>();
      
      // Execute setup operations if they exist
      if (testCase.input.setup) {
        testCase.input.setup.forEach((op: string) => {
          executeOperation(ll, op);
        });
      }
      
      // Execute main operations if they exist
      if (testCase.input.operations) {
        testCase.input.operations.forEach((op: string) => {
          executeOperation(ll, op);
        });
      }
      
      // Execute single operation if it exists
      if (testCase.input.operation) {
        const result = executeOperation(ll, testCase.input.operation);
        if (testCase.input.method === 'delete' || testCase.input.method === 'search') {
          expect(result).toBe(testCase.expected);
          return;
        }
      }
      
      // Check final result based on method
      let result: any;
      switch (testCase.input.method) {
        case 'toArray':
          result = ll.toArray();
          break;
        case 'size':
          result = ll.size();
          break;
        default:
          throw new Error(`Unknown method: ${testCase.input.method}`);
      }
      
      expect(result).toEqual(testCase.expected);
    });
  });
});

function executeOperation(ll: LinkedList<number>, operation: string): any {
  const match = operation.match(/(\w+)\((\d+)?\)/);
  if (!match) {
    throw new Error(`Invalid operation: ${operation}`);
  }
  
  const [, method, arg] = match;
  const value = arg ? parseInt(arg, 10) : undefined;
  
  switch (method) {
    case 'insert':
      return ll.insert(value!);
    case 'append':
      return ll.append(value!);
    case 'delete':
      return ll.delete(value!);
    case 'search':
      return ll.search(value!);
    case 'size':
      return ll.size();
    case 'toArray':
      return ll.toArray();
    default:
      throw new Error(`Unknown method: ${method}`);
  }
}