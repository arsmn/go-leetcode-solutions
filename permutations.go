package main

// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	ans := make([][]int, 0)
	backtrack(nums, nil, &ans)

	return ans
}

func backtrack(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrack(append(append([]int{}, nums[0:i]...), nums[i+1:]...),
			append(prev, nums[i]),
			ans)
	}
}
