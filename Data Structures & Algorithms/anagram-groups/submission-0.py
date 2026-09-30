from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams = defaultdict(list)
        for s in strs:
            word = "".join(sorted(s.lower()))
            anagrams[word].append(s)
        return list(anagrams.values())
        