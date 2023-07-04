# class Solution:
#     def singleNumber(self, nums: List[int]) -> int:

class Solution:
	def singleNumber(self, nums: list[int]) -> int:
		# sort nums in place
		nums.sort()
		
		# edge case: first element is singleton
		if len(nums) == 1 or nums[0] != nums[1]:
			return nums[0]
			
		# move sliding window across to find middle element different than outside elements
		for i in range(1,len(nums)-1):
			if nums[i-1] != nums[i] and nums[i] != nums[i+1]:
				return nums[i]

		# edge case: last element is singleton
		return nums[-1]
