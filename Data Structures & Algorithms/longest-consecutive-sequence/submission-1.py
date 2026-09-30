class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not nums:
            return 0
        table = sorted(list(set(nums)))
        longest = 1
        prev = table[0]
        streak = 1
        for num in table:
            if num  == (prev + 1):
                streak += 1
                longest = max(longest, streak)
            else:
                streak = 1
            prev = num

        return longest

        