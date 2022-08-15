package leetcodesolutions

import "strings"

// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	return helper(n, n, "", nil)
}

func helper(leftLeft, rightLeft int, queue string, res []string) []string {
	if leftLeft == 0 {
		return append(res, queue+strings.Repeat(")", rightLeft))
	}

	if rightLeft > leftLeft {
		res = helper(leftLeft, rightLeft-1, queue+")", res)
	}

	return helper(leftLeft-1, rightLeft, queue+"(", res)
}
