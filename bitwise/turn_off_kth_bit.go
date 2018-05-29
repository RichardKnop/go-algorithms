package bitwise

import (
	"strconv"
)

const intBits = uint(32)

type uintType uint32

// TurnOffKthBit ...
func TurnOffKthBit(n uintType, k uint) uintType {
	return (n & ^(1 << (k - 1)))
}

func bin2int(binStr string) uintType {
	// base 2 for binary
	result, _ := strconv.ParseUint(binStr, 2, int(intBits))
	return uintType(result)
}
