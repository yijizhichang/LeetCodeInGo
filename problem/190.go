package problem

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	count := 0
	for count < 32 {
		res <<= 1
		res |= num & 1
		num >>= 1
		count++
	}
	return res
}
