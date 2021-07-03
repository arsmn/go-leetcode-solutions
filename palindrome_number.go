package main

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	var (
		n, z = 0, x
	)

	if x < 0 {
		return false
	}

	for x > 0 {
		n *= 10
		n += x % 10
		x /= 10
	}

	return n == z
}
