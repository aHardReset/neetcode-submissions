class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        window = set()
        l = 0
        res = 0
        for r in range(len(s)): 
            if s[r] in window:
                while l <= r:
                    window.remove(s[l])
                    l += 1
                    if s[l-1] == s[r]:
                        break
            window.add(s[r])
            res = max(res, r-l+1)
        return res
        