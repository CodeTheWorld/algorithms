package _38_CopyListWithRandomPointer

/**
 * 思路：迭代
 */
func copyLinkedList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	iter := head
	nodeMap := make(map[*ListNode]*ListNode)
	nodeMap[iter] = &ListNode{Val:iter.Val}
	for iter != nil {
		if iter.Next != nil {
			if _, ok := nodeMap[iter.Next]; !ok {
				nodeMap[iter.Next] = &ListNode{Val:iter.Next.Val}
			}
			nodeMap[iter].Next = nodeMap[iter.Next]
		}
		if iter.Random != nil {
			if _, ok := nodeMap[iter.Random]; !ok {
				nodeMap[iter.Random] = &ListNode{Val:iter.Random.Val}
			}
			nodeMap[iter].Random = nodeMap[iter.Random]
		}
		iter = iter.Next
	}
	return nodeMap[head]
}
