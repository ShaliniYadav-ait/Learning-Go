package binarytree

//https://leetcode.com/problems/binary-tree-level-order-traversal/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{}
	result := [][]int{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		level := []int{}
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			level = append(level, queue[i].Val)
		}
		queue = queue[size:]
		result = append(result, level)
	}
	return result

}
