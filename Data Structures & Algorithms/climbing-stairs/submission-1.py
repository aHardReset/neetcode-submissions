class Solution:
    def climbStairsRecursion(self, n: int) -> int:
        if n < 1:
            if n == 0:
                return 1
            else:
                return 0
        return self.climbStairs(n-1) + self.climbStairs(n-2)
    
    def climbStairs(self, n: int) -> int:
        if n == 1:
            return 1
        dp = [-1] * (n+1)
        dp[-1] = 1
        dp[-2] = 1
        for i in range(n-2, -1, -1):
            dp[i] = dp[i+1] + dp[i+2]
        return dp[0]
