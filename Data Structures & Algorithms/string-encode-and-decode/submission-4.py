class Solution:
    sep = "\n"
    def encode(self, strs: List[str]) -> str:
        if len(strs) == 0:
            return ""
        encoded = self.sep.join(strs)
        if not encoded:
            return self.sep
        return encoded

    def decode(self, s: str) -> List[str]:
        if s == "":
            return []
        if s == self.sep:
            return [""]
        return s.split(self.sep)
