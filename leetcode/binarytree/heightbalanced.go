package binarytree

func isBalanced(root *TreeNode) bool {

    if root == nil{
        return true
    }
    
	diff := height(root)

	if diff != -1 {
		return true
	}

	return false
}

func height(node *TreeNode) int {

	h,d := 0,0
	if node.Left == nil && node.Right == nil{
		return 0
	}

	left := height(node.Left)
	if left == -1 {
		return -1
	}

	right := height(node.Right)
	if right == -1 {
		return -1
	}


	if left > right {
		d = left - right
        h = left + 1
	} else {
		d = right - left
		h = right +1
	}

	if d > 1 {
		return -1
	} 
		return h
}
