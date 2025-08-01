import { LinkedList } from './solution';
import { testCases } from './test-cases.json';

describe('LinkedList Core Methods', () => {
    test('append should add elements to the list', () => {
        const ll = new LinkedList();
        ll.append(1);
        expect(ll.size()).toBe(1);
        expect(ll.toArray()).toEqual([1]);

        ll.append(2);
        expect(ll.size()).toBe(2);
        expect(ll.toArray()).toEqual([1, 2]);

        ll.append(3);
        expect(ll.size()).toBe(3);
        expect(ll.toArray()).toEqual([1, 2, 3]);
    });

    test('size should return correct count', () => {
        const ll = new LinkedList();
        expect(ll.size()).toBe(0);

        ll.append(10);
        expect(ll.size()).toBe(1);

        ll.append(20);
        ll.append(30);
        expect(ll.size()).toBe(3);
    });

    test('toArray should return elements in order', () => {
        const ll = new LinkedList();
        expect(ll.toArray()).toEqual([]);

        ll.append(5);
        expect(ll.toArray()).toEqual([5]);

        ll.append(10);
        ll.append(15);
        expect(ll.toArray()).toEqual([5, 10, 15]);
    });

    test('empty list should have size 0 and empty array', () => {
        const ll = new LinkedList();
        expect(ll.size()).toBe(0);
        expect(ll.toArray()).toEqual([]);
    });

});


describe('LinkedList Middle Element', () => {
    testCases.forEach((testCase: any) => {
        test(testCase.name, () => {
            const ll = new LinkedList();

            // Build the list
            testCase.elements.forEach((element: number) => {
                ll.append(element);
            });

            // Test findMiddle
            const middleResult = ll.findMiddle();
            expect(middleResult.data).toBe(testCase.expectedMiddle);
            expect(middleResult.found).toBe(testCase.expectedFound);

            // Test toArray
            const array = ll.toArray();
            expect(array).toEqual(testCase.expectedArray);

            // Test size
            const size = ll.size();
            expect(size).toBe(testCase.expectedSize);
        });
    });
});
