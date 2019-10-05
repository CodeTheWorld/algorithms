package _19_RemoveNthNodeFromEndOfList

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	node2 := &ListNode{2, nil}
	node1 := &ListNode{1, node2}
	head := new(ListNode)
	head.Next = node1

	node := removeNthFromEnd(head, 1)
	for ; node != nil; node = node.Next {
		t.Log(node.Val)
	}
	t.Error("hhah")
}
