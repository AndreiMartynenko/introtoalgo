package main

import "fmt"

/*
Exercise 2.2-2
Consider sorting n numbers stored in array A by first finding the smallest element of
A and exchanging it with the element in A[1].
Then find the second smallest element of A, and exchange it with A[2].
Continue in this manner for the first n-1 elements of A.
Write pseudocode for this algorithm, which is known as selection sort.
What loop invariant does this algorithm maintain?
Why does it need to run for only the first n-1 elements,
rather than for all n elements?
Give the best-case and worst-case running times of selection sort in ‚theta notation.
*/

func selectionSort(arr []int) []int {
	// Implement selection sort here
	n := len(arr)
	// Outer loop iterates through the array from 0 to n-1
	for i := 0; i < n-1; i++ {
		// Inner loop finds the minimum element in the array
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// Swap the minimum element with the current element
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)

	// Sort the array using selection sort
	sortedArr := selectionSort(arr)
	fmt.Println("Sorted array:", sortedArr)

}
