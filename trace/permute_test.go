package trace

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Printf("%v", res)
}

func permute(nums []int) [][]int {
	var ans [][]int
	var dfs func(l []int, temp []int)
	dfs = func(l []int, temp []int) {
		if len(l) == 0 {
			ans = append(ans, temp)
		}
		for i := 0; i < len(l); i++ {
			n := append([]int{}, l...)
			dfs(append(n[:i], n[i+1:]...), append(temp, l[i]))
		}
	}
	dfs(nums, []int{})
	return ans
}
