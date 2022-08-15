package leetcodesolutions

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	var (
		n     int
		last  int
		roman = map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}
	)

	for _, r := range s {
		a := roman[r]

		if a <= last {
			n += a
		} else {
			n -= (2 * last) - a
		}

		last = a
	}

	return n
}
