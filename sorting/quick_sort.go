package sorting

import (
	"math/rand"
)

// QuickSort ...
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(arr)

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
