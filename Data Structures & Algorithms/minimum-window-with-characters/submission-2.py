from collections import Counter

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        t_counter = Counter(t)
        sub_s_counter = Counter()
        l=0
        min_char = s + "*"
        for r in range(len(s)):
            if s[r] in t_counter:
                sub_s_counter[s[r]] += 1
            
            while self.has_at_least_elements(t_counter, sub_s_counter):
                sub_s = s[l:r+1]
                min_char = min_char if len(min_char) < len(sub_s) else sub_s
                if s[l] in t_counter:
                    sub_s_counter[s[l]] -= 1
                l += 1

        return min_char if min_char[-1] != "*" else ""

    @staticmethod
    def has_at_least_elements(t_counter, sub_s_counter):
        for char, counter in t_counter.items():
            if char not in sub_s_counter:
                return False
            if counter > sub_s_counter[char]:
                return False
        return True

            
        