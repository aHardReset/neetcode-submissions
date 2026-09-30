class Solution:
    """
    def isPalindrome2(self, s: str) -> bool:
        s = "".join([c.lower() for c in s if c.isalnum()])
        left = 0
        right = len(s) - 1
        while left < right:
            if s[left] != s[right]:
                return False
            left += 1
            right -= 1
        return True
    """

    def isPalindrome(self, s: str) -> bool:
        s = "".join([c.lower() for c in s if c.isalnum()])
        left = 0
        right = len(s) - 1
        while left < right:
            s_left = s[left].lower()
            while not s_left.isalnum():
                left += 1
                s_left = s[left].lower()
            s_right = s[right]
            while not s_left.isalnum():
                right -= 1
                s_right = s[right].lower()
            if left < right and s[left] != s[right]:
                return False
            left += 1
            right -= 1
        return True

        