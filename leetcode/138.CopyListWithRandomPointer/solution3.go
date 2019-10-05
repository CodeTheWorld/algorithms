package _38_CopyListWithRandomPointer

func copyLinkedList3(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	iter := head
	for iter != nil {
		newNode := &ListNode{Val:iter.Val}
		newNode.Next, iter.Next, iter = iter.Next, newNode, iter.Next
	}
	iter = head
	for iter != nil {
		if iter.Random != nil {
			iter.Next.Random = iter.Random.Next
		}
		iter = iter.Next.Next
	}
	newHead := head.Next
	iter = head
	for iter.Next != nil {
		iter.Next, iter = iter.Next.Next, iter.Next
	}
	return newHead
}