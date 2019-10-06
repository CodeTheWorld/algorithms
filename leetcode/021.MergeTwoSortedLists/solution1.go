package _21_MergeTwoSortedLists

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	iter := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			iter.Next = l1
			l1 = l1.Next
		} else {
			iter.Next = l2
			l2 = l2.Next
		}
		iter = iter.Next
	}
	if l1 != nil {
		iter.Next = l1
	}
	if l2 != nil {
		iter.Next = l2
	}
	return head.Next
}
