from collections import Counter

class Solution:
    def topKFrequent2(self, nums: List[int], k: int) -> List[int]:
        counter = Counter(nums)
        
        result = list()
        sorted_counter = sorted(counter, reverse=True, key = lambda k: counter[k])
        for n in sorted_counter:
            if len(result) >= k:
                break
            result.append(n)
        return result

    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        return [pair[0] for pair in Counter(nums).most_common(k)]

        
    
        