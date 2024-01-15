# class KthLargest:

#     def __init__(self, k: int, nums: List[int]):
        

#     def add(self, val: int) -> int:

import heapq
class KthLargest:

    def __init__(self, k: int, nums: list[int]):
        self.k: int = k
        self._minheap: list[int] = nums[:k]
        heapq.heapify(self._minheap)
        for num in nums[k:]:
            self.add(num)

    def add(self, val: int) -> int:
        # returns kth largest element seen in stream after adding val
        if len(self._minheap) < self.k:
            self._minheap.append(val)
            heapq.heapify(self._minheap)
        elif val > self._minheap[0]:
            heapq.heappop(self._minheap)
            heapq.heappush(self._minheap, val)
        return self._minheap[0]

# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)
