class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        stack = [prices[0]]
        max_profit = 0
        for i, price in enumerate(prices):
            if price > stack[-1]:
                max_profit = max(max_profit, price - stack[-1])
            else:
                stack[-1] = price
        return max_profit
