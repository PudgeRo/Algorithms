package main

import "fmt"

func main() {
	arr := []int{0, 2, 3, 6, 9, 14}
	needNumber := 9
	fmt.Println("Index of desired number is", binarySearch(arr, needNumber))
}

func binarySearch(arr []int, number int) int {
	left := 0
	right := len(arr)
	var middle int
	for {
		middle = (left + right) / 2             
		if number > arr[middle] {                 
			left = middle + 1                     
		} else if number < arr[middle] {
			right = middle - 1                  
		} else {
			return middle 
		}
		if left > right {
			return -1
		}
	}
}