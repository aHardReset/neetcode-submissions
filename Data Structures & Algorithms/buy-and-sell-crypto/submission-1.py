class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        cheapest = prices[0]
        max_profit = 0
        for i, price in enumerate(prices):
            if price > cheapest:
                max_profit = max(max_profit, price - cheapest)
            else:
                cheapest = price
        return max_profit
