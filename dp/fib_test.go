package dp

import (
	"testing"
)

func TestFib(t *testing.T) {
	res := fib(10)
	println(res)
	res2 := fib2(10)
	println(res2)
}

func fib(n int) int {
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	if n < 2 {
		return dp[n]
	}
	for i := 2; i <= n; i++ {
		dp[1], dp[0] = dp[1]+dp[0], dp[1]
	}
	return dp[1]
}

func fib2(n int) int {
	f := FibClosure()
	var res int
	for i := 0; i <= n; i++ {
		res = f(i)
	}
	return res
}

func FibClosure() func(int) int {
	var back1, back2 = 0, 1
	return func(i int) (res int) {
		res = back1
		back1, back2 = back2, back1+back2
		return
	}
}
