func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	res := min(nums[l], nums[r])
	for l < r {
		m := int((l + r) / 2)
		
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
		res = min(res, nums[r])
	}

	return res
}
