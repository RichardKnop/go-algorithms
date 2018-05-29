package bitwise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnOffKthBit(t *testing.T) {
	var (
		n        = bin2int("00001111")
		k        = uint(2)
		expected = bin2int("00001101")
	)

	actual := TurnOffKthBit(n, k)

	assert.Equal(t, expected, actual)
}
