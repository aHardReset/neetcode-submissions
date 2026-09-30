class MinStack:

    def __init__(self):
        self.stack = list()
        self.min_values = [float('inf')]
        

    def push(self, val: int) -> None:
        self.stack.append(val)
        if self.stack[-1] <= self.min_values[-1]:
            self.min_values.append(self.stack[-1]) 

    def pop(self) -> None:
        value = self.stack.pop()
        if value <= self.min_values[-1]:
            self.min_values.pop()
        

    def top(self) -> int:
        return self.stack[-1]
        

    def getMin(self) -> int:
        return self.min_values[-1]
        
