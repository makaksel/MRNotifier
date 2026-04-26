package main

import (
	"fmt"
)

// quicksort sorts the array in place using the quicksort algorithm.
func quicksort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

// partition rearranges the elements in the array such that all elements less than the pivot are on the left,
// and all elements greater than the pivot are on the right.
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Original array:", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
