package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's sort stuff!")

	unsorted := []int{42, 7, 19, 3, 88, 15, 60, 2, 37, 9}

	sorted := selectionSort(unsorted)

	fmt.Printf("Sorted by selection sort: %#v", sorted)
}
