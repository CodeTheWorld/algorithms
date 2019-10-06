package _02_AddTwoNumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	node0 := &ListNode{0, nil}
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}
	node6 := &ListNode{6, nil}

	l1 := node1
	l1.Next = node0
	l1.Next.Next = node6
	l2 := node2
	l2.Next = node3
	l2.Next.Next = node4
	l2.Next.Next.Next = node5

	sum := addTwoNumbers(l1, l2)
	for sum != nil {
		t.Log(sum.Val)
		sum = sum.Next
	}
}
