package linkedlist

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	curr := head
	vals := make(map[*Node]*Node)

	for curr != nil {
		newNode := Node{
			Val: curr.Val,
		}
		vals[curr] = &newNode
		curr = curr.Next
	}

	curr = head
	for curr != nil {

		vals[curr].Next = vals[curr.Next]
		vals[curr].Random = vals[curr.Random]
		curr = curr.Next

	}
	return vals[head]
}
