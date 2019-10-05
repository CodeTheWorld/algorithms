package _91_NumberOf1Bits

func hammingWeight2(num uint32) int {
	count := 0
	mask := uint32(1)
	for i := 0; i < 32; i++ {
		if mask&num != 0 {
			count++
		}
		mask <<= 1
	}
	return count
}
