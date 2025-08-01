import { TreeNode, insert, search, inorderTraversal } from './solution';

describe('Insert', () => {
  const testCases = [
    {
      name: 'Insert into empty tree',
      initial: [],
      insertVal: 5,
      expected: [5]
    },
    {
      name: 'Insert smaller value as left child',
      initial: [10],
      insertVal: 5,
      expected: [5, 10]
    },
    {
      name: 'Insert larger value as right child',
      initial: [10],
      insertVal: 15,
      expected: [10, 15]
    },
    {
      name: 'Insert multiple values in order',
      initial: [10, 5],
      insertVal: 15,
      expected: [5, 10, 15]
    },
    {
      name: 'Insert value creating left subtree',
      initial: [10, 5, 15],
      insertVal: 3,
      expected: [3, 5, 10, 15]
    },
    {
      name: 'Insert value creating right subtree',
      initial: [10, 5, 15],
      insertVal: 18,
      expected: [5, 10, 15, 18]
    },
    {
      name: 'Insert duplicate value',
      initial: [10, 5, 15],
      insertVal: 10,
      expected: [5, 10, 15]
    },
    {
      name: 'Insert creating balanced tree',
      initial: [10, 5, 15, 3],
      insertVal: 7,
      expected: [3, 5, 7, 10, 15]
    },
    {
      name: 'Insert negative value',
      initial: [10, 5],
      insertVal: -2,
      expected: [-2, 5, 10]
    },
    {
      name: 'Insert large value',
      initial: [10, 5, 15],
      insertVal: 100,
      expected: [5, 10, 15, 100]
    }
  ];

  testCases.forEach(({ name, initial, insertVal, expected }) => {
    test(name, () => {
      let root: TreeNode | null = null;
      
      // Build initial tree
      for (const val of initial) {
        root = insert(root, val);
      }
      
      // Insert new value
      root = insert(root, insertVal);
      
      // Get inorder traversal
      const result = inorderTraversal(root);
      
      expect(result).toEqual(expected);
    });
  });
});

describe('Search', () => {
  const testCases = [
    {
      name: 'Search in empty tree',
      values: [],
      searchVal: 5,
      expected: false
    },
    {
      name: 'Search root value',
      values: [10],
      searchVal: 10,
      expected: true
    },
    {
      name: 'Search non-existent value in single node',
      values: [10],
      searchVal: 5,
      expected: false
    },
    {
      name: 'Search left child',
      values: [10, 5, 15],
      searchVal: 5,
      expected: true
    },
    {
      name: 'Search right child',
      values: [10, 5, 15],
      searchVal: 15,
      expected: true
    },
    {
      name: 'Search deep left value',
      values: [10, 5, 15, 3, 7],
      searchVal: 3,
      expected: true
    },
    {
      name: 'Search deep right value',
      values: [10, 5, 15, 3, 7, 12, 18],
      searchVal: 18,
      expected: true
    },
    {
      name: 'Search non-existent small value',
      values: [10, 5, 15],
      searchVal: 1,
      expected: false
    },
    {
      name: 'Search non-existent large value',
      values: [10, 5, 15],
      searchVal: 20,
      expected: false
    },
    {
      name: 'Search negative value',
      values: [10, 5, 15, -2],
      searchVal: -2,
      expected: true
    }
  ];

  testCases.forEach(({ name, values, searchVal, expected }) => {
    test(name, () => {
      let root: TreeNode | null = null;
      
      // Build tree
      for (const val of values) {
        root = insert(root, val);
      }
      
      // Search for value
      const result = search(root, searchVal);
      
      expect(result).toBe(expected);
    });
  });
});