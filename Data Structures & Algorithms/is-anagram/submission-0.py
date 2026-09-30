class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_map = dict()
        for c in s:
            if c not in s_map:
                s_map[c] = 0
            s_map[c] += 1
        
        for c in t:
            if c not in s_map:
                return False
            s_map[c] -= 1
            if s_map[c] < 0:
                return False
        
        for c in s_map:
            if s_map[c] > 0:
                return False
        return True
        