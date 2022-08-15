package leetcodesolutions

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var (
		start, cmax int
		runes       = []rune(s)
		mp          = make(map[rune]bool)
	)

	for i := 0; i < len(s); {
		c := runes[i]
		if _, ok := mp[c]; ok {
			cmax = max(cmax, i-start)
			mp = make(map[rune]bool)
			start++
			i = start
		} else {
			mp[c] = true
			i++
		}
	}

	return max(cmax, len(s)-start)
}
