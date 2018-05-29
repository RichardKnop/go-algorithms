package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 202, 3, 9, 6, 5, 1, 43, 506}

	mergeSort(arr, 0, len(arr))

	fmt.Println(arr)
}

func mergeSort(arr []int, l, r int) {
	if r-l <= 1 {
		return
	}

	// Recursively divide into halves
	m := int(l+r) / 2
	mergeSort(arr, l, m)
	mergeSort(arr, m, r)

	// Merge halves
	merge(arr, l, m, r)
}

func merge(arr []int, l, m, r int) {
	// Create temp slices for left and right half
	left := make([]int, m-l)
	right := make([]int, r-m)

	// Copy data to temp slices
	copy(left, arr[l:m])
	copy(right, arr[m:r])

	i, j, k := 0, 0, l
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// Left over elements in the left half
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	// Left over elements in the right half
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
