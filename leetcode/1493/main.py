# class Solution:
#     def longestSubarray(self, nums: List[int]) -> int:

class Solution:
	def longestSubarray(self, nums: list[int]) -> int:
		if len(nums) <= 1:
			return 0
		first_ones, second_ones = -1, -1
		zero_found: bool = False
		single_zero: bool = False
		maxlen: int = 0

		for i in range(len(nums)):
			if nums[i] == 1:
				if first_ones == -1:
					first_ones = i
				else:
					if single_zero:
						second_ones = i
						single_zero = False
				if i == len(nums) - 1 and first_ones != -1:
					if second_ones != -1 or (not zero_found):
						if i - first_ones  > maxlen:
							maxlen = i - first_ones
					else:
						if i - first_ones + 1 > maxlen:
							maxlen = i - first_ones + 1
			else:
				zero_found = True
				if first_ones != -1:
					if second_ones != -1:
						if i - first_ones - 1 > maxlen:
							maxlen = i - first_ones - 1
						first_ones = second_ones
						second_ones = -1
						if i < len(nums)-1 and nums[i+1] == 1:
							single_zero = True
						else:
							single_zero = False
					else:
						if nums[i-1] == 0:
							first_ones = -1
							single_zero = False
						else:
							if i - first_ones  > maxlen:
								maxlen = i - first_ones 
							if i < len(nums)-1 and nums[i+1] == 1:
								single_zero = True
							else:
								first_ones = -1
			print(i, nums[i], first_ones, second_ones, single_zero, maxlen)

		return maxlen
