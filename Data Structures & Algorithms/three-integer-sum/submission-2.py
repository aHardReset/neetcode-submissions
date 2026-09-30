class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = list()
        for i in range(len(nums)):
            if i>0 and nums[i] == nums[i-1]:
                continue
            twoDigits = self.twoSum(nums[i+1:], -nums[i])
            print(twoDigits)
            for _result in twoDigits:
                result.append([nums[i], _result[0], _result[1]])
        return result

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        left = 0
        right = len(nums) - 1
        results = []
        while left < right:
            current_sum = nums[left] + nums[right]
            if current_sum == target:
                results.append([nums[left], nums[right]])
                left += 1
                while nums[left] == nums[left-1] and left < right:
                    left += 1
            elif current_sum > target:
                right -= 1
            else:
                left += 1
        return results