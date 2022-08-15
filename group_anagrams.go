package main

import "strconv"

// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	groups := make(map[string][]string)

	for _, str := range strs {
		var freq [26]int
		for _, s := range str {
			freq[int(s)-97]++
		}

		var hash string
		for i, f := range freq {
			hash += strconv.Itoa(f)
			if i != len(freq)-1 {
				hash += "|"
			}
		}

		groups[hash] = append(groups[hash], str)
	}

	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
