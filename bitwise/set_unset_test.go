package bitwise_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RichardKnop/go-algorithms/bitwise"
)

func Test_UnsetBit(t *testing.T) {
	t.Parallel()

	n := bin2uint("00001111")

	// We will be turning off the 3rd bit (index starts at 0)
	k := 2

	expected := strings.Repeat("0", 56) + "00001011"

	actual := bitwise.UnsetBit(n, k)

	assert.Equal(t, expected, fmt.Sprintf("%.64b", actual))
}

func Test_SetBit(t *testing.T) {
	t.Parallel()

	n := bin2uint("00001111")

	// We will be turning on the 8th bit (index starts at 0)
	k := 7

	expected := strings.Repeat("0", 56) + "10001111"

	actual := bitwise.SetBit(n, k)

	assert.Equal(t, expected, fmt.Sprintf("%.64b", actual))
}

func Test_IsSet(t *testing.T) {
	t.Parallel()

	n := bin2uint("10001111")

	assert.True(t, bitwise.IsSet(n, 0))
	assert.True(t, bitwise.IsSet(n, 1))
	assert.True(t, bitwise.IsSet(n, 2))
	assert.True(t, bitwise.IsSet(n, 3))
	assert.False(t, bitwise.IsSet(n, 4))
	assert.False(t, bitwise.IsSet(n, 5))
	assert.False(t, bitwise.IsSet(n, 6))
	assert.True(t, bitwise.IsSet(n, 7))

	for k := 8; k < 64; k++ {
		assert.False(t, bitwise.IsSet(n, k))
	}
}

func bin2uint(binStr string) uint64 {
	// base 2 for binary
	result, _ := strconv.ParseUint(binStr, 2, 64)
	return uint64(result)
}

func reverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
	}
	return string(runes)
}
