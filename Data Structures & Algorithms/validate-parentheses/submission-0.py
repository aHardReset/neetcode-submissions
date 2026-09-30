class Solution:
    def isValid(self, s: str) -> bool:
        stack = list()
        order = {
            ')': '(',
            '}': '{',
            ']': '['
        }
        for char in s:
            if char in order.values():
                stack.append(char)
            else:
                if not stack or stack.pop() != order[char]:
                    return False
        return len(stack) == 0