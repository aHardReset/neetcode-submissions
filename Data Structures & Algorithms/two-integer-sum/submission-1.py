class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        result = []
        complementaries = dict()

        for i, n in enumerate(nums):
            
            complementary = target - n
            if n in complementaries:
                return [complementaries[n], i]
            complementaries[complementary] = i


        return [0,0]