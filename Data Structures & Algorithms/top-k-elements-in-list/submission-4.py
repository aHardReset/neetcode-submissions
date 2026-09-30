from collections import Counter
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        c = Counter(nums)
        l = [[] for i in range(len(nums)+1)]
        print(c)
        for n, t in c.items():
            l[t].append(n)
        
        res = []
        for i in range(len(l)-1, -1, -1):
            if v := l[i]:
                res.extend(v)
            if len(res) >= k:
                break
        return res[:k]
        