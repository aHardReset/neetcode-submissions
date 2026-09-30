class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        slow = 0
        fast = 1
        speed = 2

        while True:
            if slow != fast and nums[slow] == nums[fast]:
                return nums[slow]

            fast = (fast + speed) % len(nums)
            if (slow + 1) >= len(nums):
                speed += 1
            slow = (slow + 1) % len(nums)
        