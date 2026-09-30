import heapq
class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.heap = [-n for n in nums]
        heapq.heapify(self.heap)
        self.k = k

    def add(self, val: int) -> int:
        heapq.heappush(self.heap, -val)
        return -self.partial_k_heap_sort()
        
    def partial_k_heap_sort(self):
        val = None
        heap = self.heap.copy()
        for i in range(self.k):
            val = heapq.heappop(heap)
        return val

