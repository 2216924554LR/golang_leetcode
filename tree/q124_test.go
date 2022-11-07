package tree

import (
	"math"
	"testing"
)

func maxPathSum(root *Node) int {
	maxSum := math.MinInt32
	var maxGain func(*Node) int
	maxGain = func(node *Node) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		priceNewPath := node.Val + leftGain + rightGain

		maxSum = max(maxSum, priceNewPath)

		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum

}

func Test124(t *testing.T) {
	root := buildTree124()
	res := maxPathSum(root)
	println(res)
}
