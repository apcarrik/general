# class SmallestInfiniteSet:

#     def __init__(self):
        

#     def popSmallest(self) -> int:
        

#     def addBack(self, num: int) -> None:
        


# Your SmallestInfiniteSet object will be instantiated and called as such:
# obj = SmallestInfiniteSet()
# param_1 = obj.popSmallest()
# obj.addBack(num)

import heapq

class SmallestInfiniteSet:

	def __init__(self):
		self._smallest_int: int = 1
		self._added_back: List[int] = []
		self._added_back_set: Set[int] = set()

	def popSmallest(self) -> int:
		if len(self._added_back) != 0:
			popped: int = heapq.heappop(self._added_back)
			self._added_back_set.remove(popped)
			return popped
		else:
			self._smallest_int += 1
			return self._smallest_int - 1

	def addBack(self, num: int) -> None:
		if num < self._smallest_int and num not in self._added_back_set:
			self._added_back_set.add(num)
			heapq.heappush(self._added_back, num)
