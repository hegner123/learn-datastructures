/**
 * Finds two numbers in the array that add up to the target sum
 * and returns their indices.
 *
 * @param nums - array of integers
 * @param target - target sum to find
 * @returns array containing the indices of the two numbers [index1, index2]
 *
 * TODO: Implement the twoSum function
 * Hint: Use a Map to store values and their indices for O(n) time complexity
 */
export function twoSum(nums: number[], target: number): number[] {
    // TODO: Implement this function
    const store: Map<number, number> = new Map()
    for (let i = 0; i < nums.length; i++) {
        let complement: number = target - nums[i];
        if (store.has(complement)) {
            const c = store.get(complement)
            return [c as number, i]
        } else {
            store.set(nums[i], i)
        }
    }
    return []; // Placeholder return
}
