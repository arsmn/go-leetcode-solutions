package main

// https://leetcode.com/problems/di-string-match/

func diStringMatch(s string) []int {
	var (
		max, min = len(s), 0
		res      = make([]int, len(s)+1)
	)

	for i := 0; i < len(res); i++ {
		if i >= len(s) {
			res[i] = min
			break
		}

		if s[i] == 'I' {
			res[i] = min
			min++
		} else if s[i] == 'D' {
			res[i] = max
			max--
		}
	}

	return res
}
