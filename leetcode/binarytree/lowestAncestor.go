package binarytree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	found := 0
	var val *TreeNode

	if root == p || root == q {
		return root
	}

	val, found = lowestAncestor(root, p, q, found)
	if found == -1 {
		return val
	}

	return nil
}

func lowestAncestor(node, p, q *TreeNode, found int) (*TreeNode, int) {

	if node == nil{
		return nil, found 
	}

	if node == p || node == q {
		found++
	}

	if found == 2 {
		return nil, -2
	}

	//var left, right *TreeNode

	_, found = lowestAncestor(node.Left, p, q, found)
	if found == 2 {
		return nil, -2
	}
	if found == -1{
		return node,-1
	}

	_, found = lowestAncestor(node.Right, p, q, found)
	if found == 2 {
		return node,-1
	}

	if found == -2{
		return node,-1
	}

	return nil,found
}
