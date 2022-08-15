package leetcodesolutions

// https://leetcode.com/problems/two-sum/

// O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, xa := range nums {
		xb := target - xa
		if v, ok := m[xb]; ok {
			return []int{v, i}
		}
		m[xa] = i
	}

	return nil
}

// O(n^2)
// func twoSum(nums []int, target int) []int {

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return nil
// }
