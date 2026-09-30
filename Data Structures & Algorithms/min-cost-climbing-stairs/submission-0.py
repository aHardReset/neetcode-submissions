class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        # [10, 15]
        min_cost = 0
        table = [1000 for _ in range(len(cost)+1)]
        table[0] = 0
        table[1] = 0
        for i in range(2, len(cost)+1):
            table[i] = min(cost[i-2] + table[i-2], cost[i-1] + table[i-1])
        return table[-1]