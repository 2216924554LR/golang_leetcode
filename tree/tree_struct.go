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

func buildTree124() *Node {
	node15 := &Node{Val: 15}
	node7 := &Node{Val: 7}
	node20 := &Node{Val: 20, Left: node15, Right: node7}
	node9 := &Node{Val: 9}
	nodeN10 := &Node{Val: -10, Left: node9, Right: node20}
	return nodeN10
}
