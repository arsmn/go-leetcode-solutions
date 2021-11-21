package main

// https://leetcode.com/problems/valid-anagram/

func isAnagram(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	m := make(map[rune]int)

	for _, r := range x {
		m[r]++
	}

	for _, r := range y {
		if _, ok := m[r]; !ok {
			return false
		}
		m[r]--
		if m[r] == 0 {
			delete(m, r)
		}
	}

	return len(m) == 0
}
