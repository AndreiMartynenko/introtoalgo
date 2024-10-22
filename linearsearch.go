package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	target := 6
	result := linearSearch(arr, target)
	if result != -1 {
		fmt.Println("Element found at index: ", result)
	} else {
		fmt.Println("Element not found")
	}

}
