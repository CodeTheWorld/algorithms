package _91_NumberOf1Bits

func hammingWeight1(num uint32) int {
	count := 0
	for ;num > 0; num = num / 2 {
		count += int(num % 2)
	}
	return count
}
