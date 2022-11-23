package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func buildTree124() *Node {
	node15 := &Node{Val: 15}
	node7 := &Node{Val: 7}
	node20 := &Node{Val: 20, Left: node15, Right: node7}
	node9 := &Node{Val: 9}
	node10 := &Node{Val: -10, Left: node9, Right: node20}
	node1 := &Node{Val: 1}
	node9.Left = node1
	return node10
}

func buildTree111() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2, Left: node1}
	node3 := &Node{Val: 3, Left: node2}
	node4 := &Node{Val: 4, Left: node3}
	node5 := &Node{Val: 5, Left: node4}
	return node5
}
