package _91_NumberOf1Bits

import "testing"

func TestHammingWeight3(t *testing.T) {
	count := hammingWeight3(4)
	t.Log(count)
	count = hammingWeight3(3)
	t.Log(count)
}
