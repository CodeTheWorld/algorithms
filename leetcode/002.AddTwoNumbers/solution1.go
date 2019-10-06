package _02_AddTwoNumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{0, nil}
	iter := sum
	var bit int
	for l1 != nil || l2 != nil || bit != 0 {
		val1, val2 := 0, 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val := val1 + val2 + bit
		bit = val / 10
		iter.Next = &ListNode{val%10, nil}
		iter = iter.Next
	}
	return sum.Next
}
