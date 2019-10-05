package _19_RemoveNthNodeFromEndOfList

import (
	"testing"
)

func TestRemoveNthFromEndPointer(t *testing.T) {
	node2 := &ListNode{2, nil}
	node1 := &ListNode{1, node2}
	head := new(ListNode)
	head.Next = node1

	node := removeNthFromEndPointer(head, 2)
	for ; node != nil; node = node.Next {
		t.Log(node.Val)
	}
	t.Error("hhah")
}
