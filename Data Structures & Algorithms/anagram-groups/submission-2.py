from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams = defaultdict(list)
        for s in strs:
            word = "".join(sorted(s.lower()))
            anagrams[word].append(s)
        return list(anagrams.values())

    def groupAnagrams2(self, strs: List[str]) -> List[List[str]]:
        anagrams = dict()
        for s in strs:
            word = "".join(sorted(s.lower()))
            if word not in anagrams:
                anagrams[word] = list()
            anagrams[word].append(s)
        return list(anagrams.values())
        