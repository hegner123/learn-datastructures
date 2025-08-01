/**
 * Reverses the elements of an integer array in-place
 * TODO: Implement this function to reverse the array elements
 * The function should modify the original array and also return it
 * @param arr - The array to reverse
 * @returns The reversed array (same reference as input)
 */
export function reverseArray(arr: number[]): number[] {
    // TODO: Implement array reversal logic here
    // Hint: Use two pointers - one at the beginning and one at the end
    // Swap elements and move pointers towards center
    if (arr.length === 0) {
        return arr
    }
    let p1 = 0
    let p2 = arr.length - 1
    while (p1 < p2){
        let swap = arr[p1]
        arr[p1] = arr[p2]
        arr[p2] = swap
        p1++
        p2--
    }
    return arr;
}
