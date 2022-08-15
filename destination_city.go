package leetcodesolutions

// https://leetcode.com/problems/destination-city/

func destCity(paths [][]string) string {
	m := make(map[string]bool)

	for _, v := range paths {
		m[v[0]] = true
		if !m[v[1]] {
			m[v[1]] = false
		}
	}

	for k, v := range m {
		if !v {
			return k
		}
	}

	return ""
}
