# class Solution:
#     def sortArrayByParity(self, nums: List[int]) -> List[int]:
class Solution:
	def sortArrayByParity(self, nums: List[int]) -> List[int]:
		odd_ptr: int = 0
		even_ptr: int = 0
		while odd_ptr < len(nums) and even_ptr < len(nums):
			if nums[odd_ptr] %2 == 0: #odd_ptr is on even num
				odd_ptr+=1
				if even_ptr < odd_ptr:
					even_ptr = odd_ptr
			else: # odd_ptr is on odd num
				if nums[even_ptr] % 2 == 0: # even_ptr is on even num
					# swap elements at odd and even ptrs
					tmp:int = nums[odd_ptr]
					nums[odd_ptr] = nums[even_ptr]
					nums[even_ptr] = tmp
					odd_ptr+=1
					even_ptr+=1
				else: #even_ptr is on odd num
					# increment even ptr
					even_ptr+=1
		return nums
