class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        days = len(temperatures)
        res = [0] * days
        stack = list()
        for i in range(days):            
            while stack and temperatures[i] > temperatures[stack[-1]]:
                colder = stack.pop()
                res[colder] = i - colder
            stack.append(i)
            
        return res