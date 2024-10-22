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
	target := 4
	result := linearSearch(arr, target)
	fmt.Println(result)
}
