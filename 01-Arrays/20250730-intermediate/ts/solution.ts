// TODO: Implement the mergeSortedArrays function
// This function should merge two pre-sorted arrays into a single sorted array
// Requirements:
// - Take two sorted arrays as input
// - Return a new array containing all elements in sorted order
// - Handle different array sizes
// - Preserve the sorted order property
// - Use O(m+n) time complexity where m and n are array lengths
export function mergeSortedArrays(arr1: number[], arr2: number[]): number[] {
    // TODO: Implement merge logic using two-pointer technique;
    // 1. Create result array with capacity arr1.length + arr2.length
    // 2. Use two pointers to traverse both arrays simultaneously
    // 3. Compare elements and add the smaller one to result
    // 4. Handle remaining elements when one array is exhausted
    const out = []
    let [c1, c2] = [0, 0]

    while (c1 < arr1.length && c2 < arr2.length) {
        if (arr1[c1] <= arr2[c2]) {
            out.push(arr1[c1])
            c1++
        } else {
            out.push(arr2[c2])
            c2++
        }
    }
    return out.concat(arr1.slice(c1), arr2.slice(c2))

}
