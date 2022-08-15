package leetcodesolutions

// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}

	return -1
}
