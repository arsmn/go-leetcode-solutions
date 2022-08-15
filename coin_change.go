package leetcodesolutions

import "math"

// https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = minInt(dp[i-c]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
