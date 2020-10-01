package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/sorting"
)

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	arr := []int{2, 3, 0, 12, -5, -100, 7, 12, 4, 45, 7, 1}
	expected := []int{-100, -5, 0, 1, 2, 3, 4, 7, 7, 12, 12, 45}

	sorting.InsertionSort(arr)

	assert.Equal(t, expected, arr)
}
