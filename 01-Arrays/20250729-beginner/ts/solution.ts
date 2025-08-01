// TODO: Implement the arrayContains function
// The function should return true if the target element is found in the array, false otherwise
// Signature: function arrayContains(arr: number[], target: number): boolean
export function arrayContains(arr: number[], target: number): boolean {
    // TODO: Implement linear search through the array
    // Check each element to see if it matches the target
    // Return true if found, false if not found after checking all elements
    if (arr.length === 0) {
        return false
    }
    for (let i=0;i< arr.length;i++){
        if (arr[i]=== target){
            return true
        }
    }
    return false;
}
