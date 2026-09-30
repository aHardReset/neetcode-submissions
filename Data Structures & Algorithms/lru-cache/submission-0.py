from collections import OrderedDict
class LRUCache:

    def __init__(self, capacity: int):
        self.lru = OrderedDict()
        self.capacity = max(0, capacity)

    def get(self, key: int) -> int:
        if key in self.lru:
            self.lru.move_to_end(key)
        return self.lru.get(key, -1)

    def put(self, key: int, value: int) -> None:
        if key in self.lru:
            self.lru.move_to_end(key)
            self.lru[key] = value
        else:
            self.lru[key] = value
            while len(self.lru) > self.capacity:
                self.lru.popitem(last=False)
        
