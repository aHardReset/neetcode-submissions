func findDuplicate(nums []int) int {
    slow := 0
    fast := 1

    for nums[slow] != nums[fast]  {
        slow = (slow + 1) % len(nums)
        fast = (fast + 2) % len(nums)
        if fast == slow {
            fast = (fast + 1) % len(nums)
        }
    }
    return nums[slow]
}
