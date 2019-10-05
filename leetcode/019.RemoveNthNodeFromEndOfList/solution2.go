package _19_RemoveNthNodeFromEndOfList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node, _ := remove(head, n)
	return node
}

func remove(node *ListNode, n int) (*ListNode, int) {
	if node == nil { // 节点为nil，到达了链表尾端
		return nil, 0
	}
	next, depth := remove(node.Next, n) // 下探
	depth++ // 回退时，depth++
	if depth == n { // 倒数第n个，返回next
		return next, depth
	}
	if depth == n+1 { // 倒数第n+1个，改变next指针
		node.Next = next
	}
	return node, depth
}
