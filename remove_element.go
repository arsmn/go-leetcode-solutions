package main

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	var i int

	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}

	return i
}
