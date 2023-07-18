from collections import deque
class LRUCache:

    def __init__(self, capacity: int):
        self.capacity: int = capacity
        self.lookup: dict[int, int] = {}
        self.freq: dict[int, int] = {}
        self.dq: deque[int] = deque([])
        

    def get(self, key: int) -> int:
        if key in self.lookup:
            self.freq[key] += 1
            self.dq.append(key)
            return self.lookup[key]
        else:
            return -1
        

    def put(self, key: int, value: int) -> None:
        if key in self.lookup:
            self.freq[key] += 1
        else:
            if len(self.lookup) == self.capacity:
                while len(self.dq) > 0:
                    minkey = self.dq.popleft()
                    self.freq[minkey] -= 1
                    if self.freq[minkey] == 0:
                        del self.freq[minkey]
                        del self.lookup[minkey]
                        break
            self.freq[key] = 1
        self.dq.append(key)
        self.lookup[key] = value
                

        


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
