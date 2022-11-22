package dp

import (
	"testing"
)

func TestCoin(t *testing.T) {
	coins := []int{1, 2, 5}
	res := coinChangeDp(coins, 11)
	println(res)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChangeDp(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChangeCache(coins []int, amount int) int {
	memo := map[int]int{}
	var dp func(n int) int
	dp = func(n int) int {
		_, ok := memo[n]
		if ok {
			return memo[n]
		}
		if n == 0 {
			return 0
		}
		if n < 1 {
			return -1
		}
		res := amount + 1
		for _, coin := range coins {
			problem := dp(n - coin)
			if problem == -1 {
				continue
			}
			res = min(res, 1+problem)
		}
		if res != amount+1 {
			memo[n] = res
		} else {
			memo[n] = -1
		}
		return memo[n]
	}
	return dp(amount)
}

func coinChange(coins []int, amount int) int {
	var dp func(n int) int

	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := amount + 1
		for _, coin := range coins {
			problem := dp(n - coin)
			if problem == -1 {
				continue
			}
			res = min(res, 1+problem)
		}
		if res != amount+1 {
			return res
		}
		return -1
	}
	return dp(amount)
}
