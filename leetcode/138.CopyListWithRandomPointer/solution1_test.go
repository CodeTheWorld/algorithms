package _38_CopyListWithRandomPointer

import "testing"

func TestCopyLinkedList1(t *testing.T) {
	node5 := &ListNode{5, nil, nil}
	node4 := &ListNode{4, node5, nil}
	node3 := &ListNode{3, node4, node5}
	node2 := &ListNode{2, node3, node4}
	node1 := &ListNode{1, node2, node3}
	node5.Random = node1
	newHead := copyLinkedList1(node1)
	for newHead != nil {
		t.Log(newHead.Val)
		if newHead.Random != nil {
			t.Log("random:", newHead.Random.Val)
		} else {
			t.Log("random:nil")
		}
		if newHead.Next != nil {
			t.Log("next:", newHead.Next.Val)
		} else {
			t.Log("next:nil")
		}
		newHead = newHead.Next
	}
}
