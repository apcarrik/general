class Solution:
	def isMonotonic(self, nums: List[int]) -> bool:
		isIncreasing: bool = False
		isDecreasing: bool = False
		prevNum: int = nums[0]
		for num in nums[1:]:
			if num < prevNum:
				isDecreasing = True
			if num > prevNum:
				isIncreasing = True
			if isIncreasing and isDecreasing:
				return False
			prevNum = num
		return True
