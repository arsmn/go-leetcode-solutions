package leetcodesolutions

import "strings"

// https://leetcode.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	l1 := len(num1) - 1
	l2 := len(num2) - 1

	carry := 0
	var result strings.Builder

	for l1 >= 0 || l2 >= 0 {
		res := carry

		if l1 >= 0 {
			res += int(num1[l1] - '0')
			l1--
		}

		if l2 >= 0 {
			res += int(num2[l2] - '0')
			l2--
		}

		x := res % 10
		carry = res / 10

		result.WriteByte(byte(x) + '0')
	}

	if carry == 1 {
		result.WriteByte(byte(carry) + '0')
	}

	return reverseStr(result.String())
}

func reverseStr(str string) string {
	var result strings.Builder
	i := len(str) - 1

	for i >= 0 {
		result.WriteByte(str[i])
		i--
	}

	return result.String()
}
