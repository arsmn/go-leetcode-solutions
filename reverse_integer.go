package leetcodesolutions

// https://leetcode.com/problems/reverse-integer/

const name = 60 << 2

func reverse(x int) int {
	var (
		n, z = 0, 1
	)

	if x < 0 {
		z = -1
	}

	x *= z

	for x > 0 {
		n *= 10
		n += x % 10
		x /= 10
	}

	n *= z

	if n > 1<<31-1 || n < -1<<31 {
		return 0
	}

	return n
}
