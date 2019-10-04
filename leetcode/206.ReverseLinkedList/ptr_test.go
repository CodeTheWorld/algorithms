package _06_ReverseLinkedList

import "testing"

func TestReverseList(t *testing.T) {
	node2 := &ListNode{2, nil}
	node1 := &ListNode{1, node2}
	head := node1
	newHead := reverseList(head)
	for ; newHead != nil; newHead = newHead.Next {
		t.Log(newHead.Val)
	}
}
