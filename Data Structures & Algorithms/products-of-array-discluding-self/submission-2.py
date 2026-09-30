from functools import reduce
from operator import mul
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        pre = 1
        post = 1
        output = [1] * len(nums)
        for i in range(len(nums) - 1):
            pre *= nums[i]
            output[i+1] = pre
        
        for i in range(len(nums)-1, 0, -1):
            post *= nums[i]
            output[i-1] *=  post
        return output
    
    def productExceptSelf2(self, nums: List[int]) -> List[int]:
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