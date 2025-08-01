/**
 * Find the maximum element in an array
 * @param arr Array of numbers
 * @returns The maximum number in the array
 * @throws Error if array is empty
 */
export function findMaxElement(arr: number[]): number {
    // TODO: Implement your solution here
    if (arr.length < 1) {
        throw new Error("Method not implemented");
    }
    let max: number | null = null;
    for (let index = 0; index < arr.length; index++) {
        if (index === 0) {
            max = arr[index]
        }
        const element = arr[index];
        if (max !== null && max < element) {
            max = element
        }
    }
    if (max === null) {
        throw new Error("Max cannot be null")
    }
    return max
}
