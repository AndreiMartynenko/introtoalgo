package main

// MergeSort recursively divides and sorts the array using merge sort
func mergeSort(arr []int) []int {
	// Base case: if the array has 1 or fewer elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}
	// Step 1: Divide the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])  // Recursively sort the left half
	right := mergeSort(arr[mid:]) // Recursively sort the right half

	// Step 2: Conquer by merging sorted halves
	return merge(left, right)
}

// merge combines two sorted arrays into a single sorted array
func merge(left, right []int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Step 3: Combine by comparing elements and merging
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	// Append any remaining elements from the left or right array
	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted
}

// func main() {
// 	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
// 	fmt.Println("Original array:", arr)
// 	sortedArr := MergeSort(arr)
// 	fmt.Println("Sorted array:", sortedArr)
// }
