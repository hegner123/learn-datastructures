/**
 * Finds the minimum element in an array of numbers.
 * Returns the minimum value found in the array.
 * If the array is empty, returns 0.
 */
export function findMinElement(arr: number[]): number {
    // TODO: Implement the function to find the minimum element in the array
    // This is a beginner-level problem that focuses on basic array traversal
    // and comparison operations.
    if (arr.length === 0) {
        return 0
    }
    let min: number = arr[0]
    for (let i = 1; i < arr.length; i++) {
        if (arr[i] < min) {
            min = arr[i]
        }
    }
    return min;
}
