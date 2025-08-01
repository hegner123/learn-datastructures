// TODO: Implement the solution
// This function should calculate the sum of all elements in the given array
export function arraySum(arr: number[]): number {
    // TODO: Replace this with your implementation
    let sum: number = 0;
    if (arr.length < 1) {
        return 0
    }
    for (let i = 0; i < arr.length; i++) {
        if (i === 0) {
            sum = arr[i]
        } else {
            sum += arr[i]
        }
    }
    return sum;
}
