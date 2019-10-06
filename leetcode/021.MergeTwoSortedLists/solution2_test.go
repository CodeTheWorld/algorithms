package _21_MergeTwoSortedLists

import "testing"

func TestMergeTwoLists2(t *testing.T) {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node4 := &ListNode{4, nil}
	l1 := node1
	l1.Next = node2
	l1.Next.Next = node4

	node11 := &ListNode{1, nil}
	node3 := &ListNode{3, nil}
	node41 := &ListNode{4, nil}
	l2 := node11
	l2.Next = node3
	l2.Next.Next = node41

	res := mergeTwoLists2(l1, l2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
