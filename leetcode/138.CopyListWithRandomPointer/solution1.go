package _38_CopyListWithRandomPointer

/**
 * 思路：回溯
 */
func copyLinkedList1(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]*ListNode)
	return copyIter(head, nodeMap)
}

func copyIter(head *ListNode, nodeMap map[*ListNode]*ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if node, ok := nodeMap[head]; ok {
		return node
	}
	newNode := &ListNode{head.Val, nil, nil}
	nodeMap[head] = newNode
	newNode.Next = copyIter(head.Next, nodeMap)
	newNode.Random = copyIter(head.Random, nodeMap)
	return newNode
}
