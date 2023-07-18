import heapq
class LRUCache:

    def __init__(self, capacity: int):
        self.capacity: int = capacity
        self.lookup: dict[int, int] = {}
        self.freq: dict[int, int] = {}
        self.mh: list[tuple[int, int]] = []
        self.timestamp = 0
        

    def get(self, key: int) -> int:
        if key in self.lookup:
            self.freq[key] += 1
            heapq.heappush(self.mh, (timestamp, key))
            timestamp += 1
            return self.lookup[key]
        else:
            return -1
        

    def put(self, key: int, value: int) -> None:
        if key in self.lookup:
            self.freq[key] += 1
            heapq.heappush(self.mh, (timestamp, key))
            timestamp += 1
            self.lookup[key] = value
        else:
            pass # todo

        


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
