package leetcodesolutions

// https://leetcode.com/problems/decode-ways/

func numDecodings(s string) int {
	var (
		count = make([]int, len(s)+1)
	)

	count[0] = 1
	count[1] = 1

	for i := 2; i <= len(s); i++ {
		count[i] = 0

		if s[i-1] > '0' {
			count[i] = count[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			count[i] += count[i-2]
		}
	}

	return count[len(s)]
}
