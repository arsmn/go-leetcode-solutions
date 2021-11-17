package main

func isMonotonic(nums []int) bool {
	var status int

	for i, n := range nums {
		if status == 0 && i >= 1 {
			if nums[i-1] > n {
				status = 1
			} else if nums[i-1] < n {
				status = 2
			}
		} else if status > 0 {
			if nums[i-1] > n && status == 2 {
				return false
			} else if nums[i-1] < n && status == 1 {
				return false
			}
		}
	}

	return true
}
