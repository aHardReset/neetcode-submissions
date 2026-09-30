from functools import reduce
from operator import mul
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        zeros = 0
        for num in nums:
            if num == 0:
                zeros += 1
            if zeros > 1:
                break
        if zeros > 1:
            return [0] * len(nums)

        productAll = reduce(mul, (num for num in nums if num != 0))
        if zeros == 1:
            return [productAll if num == 0 else 0 for num in nums ]
        return [productAll//num for num in nums]