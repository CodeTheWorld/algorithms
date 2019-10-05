package _02_BinaryTreeLevelOrderTraversal

import "testing"

func TestLevelOrder2(t *testing.T) {
	node9 := &TreeNode{9, nil, nil}
	node15 := &TreeNode{15, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node20 := &TreeNode{20, node15, node7}
	node3 := &TreeNode{3, node9, node20}
	res := levelOrder2(node3)

	t.Log(res)
}
