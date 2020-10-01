package bitwise

func LeftRotate(n uint8, k uint) uint8 {
	return n<<k | (n >> (8 - k))
}
