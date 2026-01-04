package main

func selectionSort(unsorted []int) []int {
	// make a copy so original is untouched as Go slices behave like pass by reference
	arr := make([]int, len(unsorted))
	copy(arr, unsorted)

	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}

	return arr
}
