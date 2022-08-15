package leetcodesolutions

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	var prefix string

	for i, r := range strs[0] {
		for _, str := range strs {
			runes := []rune(str)
			if len(runes) <= i || runes[i] != r {
				return prefix
			}
		}
		prefix += string(r)
	}

	return prefix
}
