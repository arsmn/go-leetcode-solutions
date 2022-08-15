package leetcodesolutions

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	m := map[int]*Node{}
	traverse(root, 0, m)
	return root
}

func traverse(root *Node, height int, m map[int]*Node) {
	if root == nil {
		return
	}

	traverse(root.Left, height+1, m)
	traverse(root.Right, height+1, m)

	if _, ok := m[height]; ok {
		m[height].Next = root
		m[height] = m[height].Next
	} else {
		m[height] = root
	}
}
