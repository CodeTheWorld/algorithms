package _06_ReverseLinkedList

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := reverseListRecursion(head.Next)
	if tail == nil {
		return head
	}
	head.Next.Next = head
	head.Next = nil
	return tail
}
