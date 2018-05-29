package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506}

	insertionSort(items)

	fmt.Println(items)
}

func insertionSort(items []int) {
	var (
		n = len(items)
		i = 1
	)
	for i < n {
		j := i
		for j > 0 && items[j-1] > items[j] {
			items[j-1], items[j] = items[j], items[j-1]
			j = j - 1
		}
		i += 1
	}

}
