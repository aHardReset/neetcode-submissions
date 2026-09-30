class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter = dict()
        for n in nums:
            if n not in counter:
                counter[n] = 0
            counter[n] += 1
        
        result = list()
        sorted_counter = sorted(counter, reverse=True, key = lambda k: counter[k])
        for n in sorted_counter:
            if len(result) >= k:
                break
            result.append(n)
        return result
        