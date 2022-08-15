package leetcodesolutions

// https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
	const plus = 1

	var carry int

	for i := len(digits) - 1; i >= 0; i-- {
		res := digits[i] + carry
		if i == len(digits)-1 {
			res += plus
		}

		if res < 10 {
			carry = 0
			digits[i] = res
		} else {
			carry = 1
			digits[i] = res % 10
		}
	}

	if carry == 0 {
		return digits
	}

	return append([]int{carry}, digits...)
}
