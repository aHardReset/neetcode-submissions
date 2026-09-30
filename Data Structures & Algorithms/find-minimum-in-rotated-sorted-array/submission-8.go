func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	res := min(nums[l], nums[r])
	for l < r {
		m := int((l + r) / 2)
		res = min(res, nums[m], nums[l], nums[r])
		if nums[m] > nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
