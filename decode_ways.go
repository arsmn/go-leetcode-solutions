package leetcodesolutions

import "strconv"

// https://leetcode.com/problems/decode-ways/

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1

	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i < len(dp); i++ {
		oneD, _ := strconv.Atoi(s[i-1 : i])
		twoD, _ := strconv.Atoi(s[i-2 : i])

		if oneD >= 1 {
			dp[i] += dp[i-1]
		}

		if twoD >= 10 && twoD <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
