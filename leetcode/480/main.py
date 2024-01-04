# class Solution:
#     def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
        
import heapq
class FindMedian:
    def __init__(self):
        self.maxh: list[int] = []
        self.minh: list[int] = []

    def rebalance(self):
        if len(self.minh) - len(self.maxh) > 1:
            elem: tuple(int, int) = heapq.heappop(self.minh)
            heapq.heappush(self.maxh, (-1*elem[0], elem[1]))
        elif len(self.maxh) - len(self.minh) > 0:
            elem: tuple(int, int) = heapq.heappop(self.maxh)
            heapq.heappush(self.minh, (-1*elem[0], elem[1]))

    def addNum(self, n: int, i: int):
        # add new element and rebalance
        if len(self.minh) == 0 or n >= self.minh[0][0]:
            heapq.heappush(self.minh,(n, i))
        else:
            heapq.heappush(self.maxh,(-1*n, i))
        self.rebalance()

    def addNumWithReplace(self, n: int, oldN: int, i: int, k: int):
        # add new element and rebalance if necessary
        if n >= self.minh[0][0]:
            heapq.heappush(self.minh,(n, i+k))
            if oldN <= self.minh[0][0]:
                # move from minh to maxh
                elem: tuple(int, int) = heapq.heappop(self.minh)
                heapq.heappush(self.maxh, (-1*elem[0], elem[1]))
        else:
            heapq.heappush(self.maxh,(-1*n, i+k))
            if oldN >= self.minh[0][0]:
                # move from maxh to minh
                elem: tuple(int, int) = heapq.heappop(self.maxh)
                heapq.heappush(self.minh, (-1*elem[0], elem[1]))

        # remove stale elements
        while len(self.maxh) > 0 and self.maxh[0][1] <= i:
            heapq.heappop(self.maxh)
        while len(self.minh) > 0 and self.minh[0][1] <= i:
            heapq.heappop(self.minh)

    def getMedian(self, k: int) -> float:
        if k % 2 == 1:
            return float(self.minh[0][0])
        else:
            if len(self.maxh) == 0:
                return 0.0
            return ((self.maxh[0][0]*-1) + self.minh[0][0]) / 2
        

class Solution:
    def medianSlidingWindow(self, nums: list[int], k: int) -> list[float]:
        n: int = len(nums)
        if n == 0 or k > n:
            return []

        fm: FindMedian = FindMedian()
        
        # place all of initial window on mf
        for i in range(k):
            fm.addNum(nums[i], i)

        # start moving window
        medians: list[float] = [fm.getMedian(k)]

        for i in range(0,n-k):
            # push new item
            fm.addNumWithReplace(nums[i+k], nums[i],i,k)

            # find new median
            medians.append(fm.getMedian(k))
        return medians
