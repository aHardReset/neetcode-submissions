import heapq
class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.heap = nums.copy()
        heapq.heapify(self.heap)
        self.k = k
        while len(self.heap) > self.k:
            heapq.heappop(self.heap)

    def add(self, val: int) -> int:
        if self.heap:
            if val >= self.heap[0]:
                heapq.heappush(self.heap, val)
                if len(self.heap) > self.k:
                    heapq.heappop(self.heap)
        else:
            heapq.heappush(self.heap, val)
        return self.heap[0]

