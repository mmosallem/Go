package main

//reverseArray reverses an array
func reverseArray(arr []int) []int {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

//isArraySortedMax1Swap checks to see if the array is sorted in ascending order, or if it can be sorted with 1 item change
//scan through the array using a window of 3 elements
func isArraySortedMax1Swap(arr []int) bool {
	var swapped bool = false

	for i := 0; i < len(arr)-2; i = i + 2 {
		if arr[i] > arr[i+1] {
			if swapped {
				return false
			}
			swapped = true

		}
		if arr[i+1] > arr[i+2] {
			if swapped {
				return false
			}
			swapped = true
		}
	}
	if arr[len(arr)-2] > arr[len(arr)-1] {
		return !swapped
	}

	return true
}
