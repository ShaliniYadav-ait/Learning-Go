package queue

// TreeNode declaration
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Que declaration
type Que struct {
	val []*TreeNode
}

func initial() Que {
	return Que{val: []*TreeNode{}}
}

// Push tree node
func (Q *Que) Push(node *TreeNode) {

	Q.val = append(Q.val, node)

}

//Pop function
func (Q *Que) Pop() *TreeNode {
	temp := Q.val[0]
	Q.val = Q.val[1:]
	return temp
}

func bfs(root *TreeNode) [][]*TreeNode {

	if root == nil {
		return [][]*TreeNode{}
	}

	queue := initial()
	result := [][]*TreeNode{}
	queue.Push(root)
	size := len(queue.val)
	currLevel := 0

	for size != 0 {
        result = append(result,nil)
		for i := 0; i < size; i++ {
			val := queue.Pop()
			result[currLevel][i] = val

			if val.Left != nil {
				queue.Push(val.Left)
			}
			if val.Right != nil {
				queue.Push(val.Right)
			}
		}
        currLevel++
		size = len(queue.val)
	}
return result
}
