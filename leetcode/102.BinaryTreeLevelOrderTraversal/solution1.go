package _02_BinaryTreeLevelOrderTraversal

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	leftOrder := levelOrder1(root.Left)
	rightOrder := levelOrder1(root.Right)
	leftLen := len(leftOrder)
	rightLen := len(rightOrder)

	maxLen := leftLen
	if rightLen > maxLen {
		maxLen = rightLen
	}

	res := make([][]int, maxLen+1)
	res[0] = []int{root.Val}
	for i := 0; i < maxLen; i++ {
		leftArr := make([]int, 0)
		if i < leftLen {
			leftArr = leftOrder[i]
		}
		rightArr := make([]int, 0)
		if i <  rightLen {
			rightArr = rightOrder[i]
		}
		res[i+1] = append(leftArr, rightArr...)
	}
	return res
}
