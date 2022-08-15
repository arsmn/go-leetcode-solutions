package leetcodesolutions

import "math"

// https://leetcode.com/problems/maximum-number-of-balloons/

func maxNumberOfBalloons(text string) int {
	const word = "balloon"

	wmap := make(map[rune]int)
	for _, v := range word {
		wmap[v]++
	}

	tmap := make(map[rune]int)
	for _, v := range text {
		if _, ok := wmap[v]; ok {
			tmap[v]++
		}
	}

	if len(tmap) != len(wmap) {
		return 0
	}

	res := math.MaxInt

	for k, v := range wmap {
		if x := tmap[k] / v; x < res {
			res = x
		}
	}

	return res
}
