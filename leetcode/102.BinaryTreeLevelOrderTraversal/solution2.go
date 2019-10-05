package _02_BinaryTreeLevelOrderTraversal

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	slice := []*TreeNode{root}
	depth := 1
	res := make([][]int, 0)
	for len(slice) != 0 {
		levelLength := len(slice)
		level := make([]int, 0)
		for i := 0; i < levelLength; i++ {
			if slice[i].Left != nil {
				slice = append(slice, slice[i].Left)
			}
			if slice[i].Right != nil {
				slice = append(slice, slice[i].Right)
			}
			level = append(level, slice[i].Val)
		}
		slice = slice[levelLength:]
		res = append(res, level)
		depth++
	}
	return res
}
