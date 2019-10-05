package _06_ReverseLinkedList

func reverseList(head *ListNode) *ListNode {
	var reverseHead *ListNode
	for head != nil {
		reverseHead, head.Next, head = head, reverseHead, head.Next
	}
	return reverseHead
}