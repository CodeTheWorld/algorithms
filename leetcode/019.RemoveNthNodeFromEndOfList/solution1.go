package _19_RemoveNthNodeFromEndOfList

func removeNthFromEndPointer(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head} // 新的head，预防head节点被删掉的情况
	slowPtr, quickPtr := newHead, head
	for ; n > 1 && nil != quickPtr; n-- { // 快指针先走n步
		quickPtr = quickPtr.Next
	}
	if nil == quickPtr { // 快指针到头，返回原链表
		return head
	}
	for nil != quickPtr.Next { // 快慢指针齐步走，直到慢指针到达尽头
		quickPtr = quickPtr.Next
		slowPtr = slowPtr.Next
	}
	slowPtr.Next = slowPtr.Next.Next
	return newHead.Next
}
