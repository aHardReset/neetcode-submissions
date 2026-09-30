func search(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l <= r {
		m := int((l+r)/2)
		if nums[m] == target{
			return m
		}
		if nums[l] == target{
			return l
		}
		if nums[r] == target{
			return r
		}
		if target < nums[m] && target >= nums[l]{
			r = r - 1
		} else {
			l = l + 1
		}
	}

	return -1

}


// [3,4,5,6,1,2]
// [1,2,3,4,5,6]