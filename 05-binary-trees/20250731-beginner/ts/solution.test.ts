import { BinaryTree } from './solution';
import testCases from './test-cases.json';

describe('BinaryTree Height', () => {
    const heightTests = testCases.height;

    test.each(heightTests)('$name', (testCase: any) => {
        const { values, expected } = testCase;
        const bt = new BinaryTree();
        
        // Build tree using level-order insertion
        for (const value of values) {
            bt.insert(value);
        }
        
        const result = bt.height(bt.root);
        expect(result).toBe(expected);
    });
});

describe('BinaryTree Size', () => {
    const sizeTests = testCases.size;

    test.each(sizeTests)('$name', (testCase: any) => {
        const { values, expected } = testCase;
        const bt = new BinaryTree();
        
        // Build tree using level-order insertion
        for (const value of values) {
            bt.insert(value);
        }
        
        const result = bt.size();
        expect(result).toBe(expected);
    });
});

describe('BinaryTree IsEmpty', () => {
    test('Empty tree should return true', () => {
        const bt = new BinaryTree();
        expect(bt.isEmpty()).toBe(true);
    });

    test('Non-empty tree should return false', () => {
        const bt = new BinaryTree();
        bt.insert(1);
        expect(bt.isEmpty()).toBe(false);
    });
});
