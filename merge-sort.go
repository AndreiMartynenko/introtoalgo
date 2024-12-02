package main

import "fmt"

/*

To illustrate the operation of **Merge Sort** on the array `A = (3, 41, 52, 26, 38, 57, 9, 49)`,
we follow the steps of the **Divide and Conquer** approach. In **Merge Sort**,
the array is repeatedly divided into two halves until each subarray has one element,
then the subarrays are merged back together in sorted order.

Here’s a step-by-step explanation, matching Figure 2.4 (which I assume shows the recursive division and merging):

---

### Initial Array:
`A = (3, 41, 52, 26, 38, 57, 9, 49)`

---

### Step 1: Divide the Array
Divide the array into two halves:
- Left: `(3, 41, 52, 26)`
- Right: `(38, 57, 9, 49)`

---

### Step 2: Divide Again
Each half is divided further:
- Left half `(3, 41, 52, 26)` →
  - `(3, 41)` and `(52, 26)`
- Right half `(38, 57, 9, 49)` →
  - `(38, 57)` and `(9, 49)`

---

### Step 3: Continue Dividing
We keep dividing until we have single-element subarrays:
- `(3, 41)` → `(3)` and `(41)`
- `(52, 26)` → `(52)` and `(26)`
- `(38, 57)` → `(38)` and `(57)`
- `(9, 49)` → `(9)` and `(49)`

At this point, we have:
- `(3)`, `(41)`, `(52)`, `(26)`, `(38)`, `(57)`, `(9)`, `(49)`

---

### Step 4: Start Merging
Now, we begin merging the subarrays back together, sorting as we go.

1. Merge `(3)` and `(41)`:
   - Result: `(3, 41)`

2. Merge `(52)` and `(26)`:
   - Result: `(26, 52)`

3. Merge `(38)` and `(57)`:
   - Result: `(38, 57)`

4. Merge `(9)` and `(49)`:
   - Result: `(9, 49)`

Now, we have:
- `(3, 41)`, `(26, 52)`, `(38, 57)`, `(9, 49)`

---

### Step 5: Merge Further
Now, we merge the larger subarrays:

1. Merge `(3, 41)` and `(26, 52)`:
   - Compare and merge:
     - Result: `(3, 26, 41, 52)`

2. Merge `(38, 57)` and `(9, 49)`:
   - Compare and merge:
     - Result: `(9, 38, 49, 57)`

Now, we have:
- `(3, 26, 41, 52)`, `(9, 38, 49, 57)`

---

### Step 6: Final Merge
Finally, merge the two sorted subarrays `(3, 26, 41, 52)` and `(9, 38, 49, 57)`:

- Compare elements from both arrays and merge:
  - Result: `(3, 9, 26, 38, 41, 49, 52, 57)`

---

### Final Sorted Array:
After all the recursive divisions and merges, the sorted array is:
`(3, 9, 26, 38, 41, 49, 52, 57)`

---

### Illustration Summary:
1. **Divide** the array into two halves repeatedly.
2. **Merge** subarrays back together, sorting at each step.
3. The time complexity is **O(n log n)** due to the dividing and merging process.

This is a step-by-step example of how Merge Sort works on an array. If you need a diagram or further clarification, feel free to ask!

*/

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
	// sorted := make([]int, 0, len(left)+len(right))
	result := []int{}
	i, j := 0, 0

	// Step 3: Combine by comparing elements and merging
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left or right array
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println("Original array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

/*

Time Complexity
Best Case:
𝑂(𝑛log𝑛)

Worst Case:
𝑂(𝑛log𝑛)

Average Case:
𝑂(𝑛log𝑛)

*/
