package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{4, 202, 3, 9, 6, 5, 1, 43, 506}

	quickSort(arr)

	fmt.Println(arr)
}

func quickSort(arr []int) {
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
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}
