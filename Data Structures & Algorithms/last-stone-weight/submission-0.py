class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        if len(stones) == 1:
            return stones[0]
        while True:
            stones.sort()
            stone_1 = stones.pop()
            stone_2 = stones.pop()
            
                
            if stone_1 > stone_2:
                new_stone = stone_1 - stone_2
                stones.append(new_stone)
            
            if len(stones) <= 1:
                return 0 if len(stones) == 0 else stones.pop()

        
                


