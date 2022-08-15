package main

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/reorder-data-in-log-files/

func reorderLogFiles(logs []string) []string {
	dLog := make([]string, 0)
	lLog := make([]string, 0)

	for _, v := range logs {
		temp := strings.Fields(v)
		tmp := []rune(temp[1])

		if tmp[0] < 'a' || tmp[0] > 'z' {
			dLog = append(dLog, v)
			continue
		}

		lLog = append(lLog, v)
	}

	sort.SliceStable(lLog, func(i, j int) bool {
		iIndex := strings.Index(lLog[i], " ")
		jIndex := strings.Index(lLog[j], " ")

		iLog := lLog[i][iIndex:]
		jLog := lLog[j][jIndex:]

		if iLog == jLog {
			return lLog[i] < lLog[j]
		}

		return iLog < jLog
	})

	return append(lLog, dLog...)
}
