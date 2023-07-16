# class Solution:
#     def findSmallestInteger(self, nums: List[int], value: int) -> int:

from collections import Counter
class Solution:
	def findSmallestInteger(self, nums: List[int], value: int) -> int:
		for i in range(len(nums)):
			nums[i] = nums[i] % value
		c: Counter = Counter(nums)
		if len(c) < value:
			for i in range(value+1):
				if i not in c:
					return i
		minfreq: int = min(c.values())
		if len(set(c.values())) == 1:
			return value*minfreq		
		for i in range(value+1):
			if c[i] == minfreq:
				return c[i]*value  + i % value 
