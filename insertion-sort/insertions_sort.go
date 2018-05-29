package insertionsort

// InsertionSort ...
func InsertionSort(items []int) {
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
