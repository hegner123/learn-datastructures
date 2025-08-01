package main


// TODO: Implement the MergeSortedArrays function
// This function should merge two pre-sorted arrays into a single sorted array
// Requirements:
// - Take two sorted arrays as input
// - Return a new array containing all elements in sorted order
// - Handle different array sizes
// - Preserve the sorted order property
// - Use O(m+n) time complexity where m and n are array lengths
func MergeSortedArrays(arr1, arr2 []int) []int {
	// TODO: Implement merge logic using two-pointer technique
	n := len(arr1) + len(arr2)
	out := make([]int, n)
	c1 := 0
	c2 := 0
	for i := 0; i < n; i++ {
		if c1 == len(arr1) {
			for c2 != len(arr2) {
				out[i] = arr2[c2]
				i++
				c2++
			}
			break
		}
		if c2 == len(arr2) {
			for c1 != len(arr1) {
				out[i] = arr1[c1]
				i++
				c1++
			}
			break
		}
		if arr1[c1] < arr2[c2] {
			out[i] = arr1[c1]
			c1++
		} else {
			out[i] = arr2[c2]
			c2++
		}

	}
	return out
}

// Industry standard solution using cleaner two-pointer approach
func MergeSortedArraysStandard(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	
	// Merge while both arrays have elements
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	
	// Append remaining elements
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	
	return result
}

