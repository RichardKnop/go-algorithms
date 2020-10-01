package bitwise_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/bitwise"
)

func TestLeftRotate(t *testing.T) {
	t.Parallel()

	// Inputs
	input, _ := strconv.ParseUint("01001111", 2, 8)
	n := uint8(input)
	k := uint(2)

	// Expected output
	output, _ := strconv.ParseUint("00111101", 2, 8)
	expected := uint8(output)

	actual := bitwise.LeftRotate(n, k)

	assert.Equal(t, fmt.Sprintf("%.8b", expected), fmt.Sprintf("%.8b", actual))
}
