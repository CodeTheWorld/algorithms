package _91_NumberOf1Bits

import "testing"

func TestHammingWeight2(t *testing.T) {
	count := hammingWeight2(4)
	t.Log(count)
	count = hammingWeight2(3)
	t.Log(count)
}
