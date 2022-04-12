package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 6, 2, 0, 8}
	fmt.Printf("Sorted array: %v\n", bubleSort(arr))
}

func bubleSort(arr []int) []int {
	swapped := true 
	for swapped {
		swapped = false 
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}