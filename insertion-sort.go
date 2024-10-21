package main

import "fmt"

func insertionSort(arr []int) []int {
	// Implement insertion sort here
	for j := 0; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i > 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}
	return arr

}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	result := insertionSort(arr)
	fmt.Println(result)

}
