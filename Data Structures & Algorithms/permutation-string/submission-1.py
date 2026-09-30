from collections import Counter

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        counter = Counter(s1)
        windowCounter = Counter()
        l = 0
        for r in range(len(s2)):
            if s2[r] in counter:
                windowCounter[s2[r]] += 1
                while counter[s2[r]] < windowCounter[s2[r]]:
                    windowCounter[s2[l]] -= 1
                    l += 1
                print(windowCounter, counter)
                if windowCounter == counter:
                    return True                    
            else:
                l = r
                windowCounter = Counter()
        return False