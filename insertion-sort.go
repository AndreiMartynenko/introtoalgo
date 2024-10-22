package main

import "fmt"

// non-decreasing order
//func insertionSort(arr []int) []int {
//	// Implement insertion sort here
//	for j := 1; j < len(arr); j++ {
//		key := arr[j]
//		i := j - 1
//		for i >= 0 && arr[i] > key {
//			arr[i+1] = arr[i]
//			i = i - 1
//		}
//		arr[i+1] = key
//	}
//	return arr
//
//}

// non-increasing order

func insertionSort(arr []int) []int {
	// Implement insertion sort here
	for j := 1; j < len(arr); j++ {
		//fmt.Println(j)
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] < key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}
	return arr
}

func main() {
	arr := []int{8, 3, 9, 6, 1, 7}
	result := insertionSort(arr)
	fmt.Println(result)

}
