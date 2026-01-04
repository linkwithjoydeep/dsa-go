package main

// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stable
func bubbleSort(unsorted []int) []int {
	// make a copy so original is untouched as Go slices behave like pass by reference
	arr := make([]int, len(unsorted))
	copy(arr, unsorted)
	swapped := false
	for i := 0; i < len(arr)-1; i++ {
		swapped = false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return arr
}
