# class Solution:
#     def maximumSubarraySum(self, nums: List[int], k: int) -> int:

class Solution:  
  def maximumSubarraySum(self, nums: list[int], k: int) -> int:
    max_subarray_sum: int = 0
    start: int = 0
    windowSum: int = 0
    digitMap: dict[int: int] = {}

    for end in range(len(nums)):
      digit: int = nums[end]

      if digit in digitMap:
        dup_idx: int = digitMap[digit]
        for i in range(start,dup_idx+1):
          windowSum -= nums[i]
          del(digitMap[nums[i]])	
        digitMap[digit] = end	
        windowSum += digit
        start = dup_idx+1

      else:
        digitMap[digit] = end
        windowSum += digit
        if end - start == k-1:
          if windowSum > max_subarray_sum:
            max_subarray_sum = windowSum
          windowSum -= nums[start]
          del(digitMap[nums[start]])	
          start += 1
      
    return max_subarray_sum
