# class MedianFinder:

#     def __init__(self):
        

#     def addNum(self, num: int) -> None:
        

#     def findMedian(self) -> float:
        


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()


import heapq

class MaxHeapObj(object):
    def __init__(self, val):
        self.val = val
    def __lt__(self, other):
        return self.val > other.val
    def __eq__(self, other):
        return self.val == other.val
    def __str__(self):
        return str(self.val)

class MedianFinder:

    def __init__(self):
        self._maxheap: list[MaxHeapObj] = [] 
        self._minheap: list[int] = []

    def addNum(self, num: int):

        # find which heap num is to be placed on
        median: float = self.findMedian()
        place_on_maxheap: bool = False
        if float(num) > median:
            # place on minheap
            place_on_maxheap = False
        else:
            # place on maxheap
            place_on_maxheap = True
    
        # place num on correct heap
        if place_on_maxheap:
            heapq.heappush(self._maxheap, MaxHeapObj(num))    
        else:
            heapq.heappush(self._minheap, num)

        # balance heaps if necessary
        if len(self._maxheap) - len(self._minheap) > 1:
            heapq.heappush(self._minheap, heapq.heappop(self._maxheap).val)
        elif  len(self._minheap) - len(self._maxheap) > 1:
            heapq.heappush(self._maxheap, MaxHeapObj(heapq.heappop(self._minheap)))
        

    def findMedian(self) -> float:
        if len(self._maxheap) > len(self._minheap):
            return float(self._maxheap[0].val)
        elif len(self._maxheap) < len(self._minheap):
            return float(self._minheap[0])
        else:
            if len(self._maxheap) == 0:
                return 0.0
            return (self._maxheap[0].val + self._minheap[0]) / 2.0
