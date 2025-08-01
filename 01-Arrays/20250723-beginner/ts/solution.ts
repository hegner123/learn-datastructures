// TODO: Implement the solution
// This function should find and return the maximum element in the given array
// Return 0 for empty arrays
export function arrayMax(arr: number[]): number {
    // TODO: Replace this with your implementation
    if (arr.length === 0) {
        return 0
    }
    let max = 0
    for (let i = 0; i < arr.length; i++) {
        if (i === 0) {
            max = arr[i]

        } else {
            if (arr[i] > max) {
                max = arr[i]
            }
        }
    }
    return max
}
