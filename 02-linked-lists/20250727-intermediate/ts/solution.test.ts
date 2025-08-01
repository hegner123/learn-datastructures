import { LinkedList, mergeSortedLists } from './solution';

describe('MergeSortedLists', () => {
    test('Test 1: Both lists empty', () => {
        const list1 = new LinkedList();
        const list2 = new LinkedList();
        const result = mergeSortedLists(list1, list2);
        const expected: number[] = [];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 2: First list empty', () => {
        const list1 = new LinkedList();
        const list2 = LinkedList.createFromArray([1, 3, 5]);
        const result = mergeSortedLists(list1, list2);
        const expected = [1, 3, 5];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 3: Second list empty', () => {
        const list1 = LinkedList.createFromArray([2, 4, 6]);
        const list2 = new LinkedList();
        const result = mergeSortedLists(list1, list2);
        const expected = [2, 4, 6];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 4: Simple merge', () => {
        const list1 = LinkedList.createFromArray([1, 3, 5]);
        const list2 = LinkedList.createFromArray([2, 4, 6]);
        const result = mergeSortedLists(list1, list2);
        const expected = [1, 2, 3, 4, 5, 6];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 5: Lists with different lengths', () => {
        const list1 = LinkedList.createFromArray([1, 5, 9, 15]);
        const list2 = LinkedList.createFromArray([2, 3]);
        const result = mergeSortedLists(list1, list2);
        const expected = [1, 2, 3, 5, 9, 15];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 6: All elements from first list come first', () => {
        const list1 = LinkedList.createFromArray([1, 2, 3, 4]);
        const list2 = LinkedList.createFromArray([5, 6, 7, 8]);
        const result = mergeSortedLists(list1, list2);
        const expected = [1, 2, 3, 4, 5, 6, 7, 8];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 7: All elements from second list come first', () => {
        const list1 = LinkedList.createFromArray([10, 20, 30]);
        const list2 = LinkedList.createFromArray([1, 2, 3, 4, 5]);
        const result = mergeSortedLists(list1, list2);
        const expected = [1, 2, 3, 4, 5, 10, 20, 30];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 8: Duplicate elements', () => {
        const list1 = LinkedList.createFromArray([1, 3, 5, 5]);
        const list2 = LinkedList.createFromArray([2, 3, 4, 5]);
        const result = mergeSortedLists(list1, list2);
        const expected = [1, 2, 3, 3, 4, 5, 5, 5];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 9: Single element lists', () => {
        const list1 = LinkedList.createFromArray([5]);
        const list2 = LinkedList.createFromArray([3]);
        const result = mergeSortedLists(list1, list2);
        const expected = [3, 5];
        expect(result.toArray()).toEqual(expected);
    });

    test('Test 10: Negative numbers', () => {
        const list1 = LinkedList.createFromArray([-5, -1, 3]);
        const list2 = LinkedList.createFromArray([-3, 0, 2, 4]);
        const result = mergeSortedLists(list1, list2);
        const expected = [-5, -3, -1, 0, 2, 3, 4];
        expect(result.toArray()).toEqual(expected);
    });
});