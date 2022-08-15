package leetcodesolutions

// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	var counter, last int

	for _, r := range s {
		if r == ' ' {
			if counter > 0 {
				last = counter
				counter = 0
			}
		} else {
			counter++
		}
	}

	if counter > 0 {
		last = counter
	}

	return last
}
