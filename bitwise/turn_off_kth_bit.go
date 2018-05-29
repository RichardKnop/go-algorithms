package bitwise

// TurnOffKthBit ...
func TurnOffKthBit(n uint8, k uint) uint8 {
	return (n & ^(1 << k))
}
