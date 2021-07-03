package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var (
		start, max int
		runes      = []rune(s)
		mp         = make(map[rune]bool)
	)

	for i := 0; i < len(s); {
		c := runes[i]
		if _, ok := mp[c]; ok {
			max = maxInt(max, i-start)
			mp = make(map[rune]bool)
			start++
			i = start
		} else {
			mp[c] = true
			i++
		}
	}

	return maxInt(max, len(s)-start)
}
