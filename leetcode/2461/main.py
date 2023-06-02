# class Solution:
#     def maximumSubarraySum(self, nums: List[int], k: int) -> int:

from collections import Counter
class Solution:
  def maximumSubarraySum(self, nums: list[int], k: int) -> int:
    # Idea: just make a fixed-width sliding window of k and test all possible substrings
    # Better: make it a sliding window and keep track of the sum, only adding/removing one element at a time
    max_subarray_sum: int = 0
    window_counter: Counter = Counter(nums[:k])
    window_sum: int = sum(nums[:k]) # THIS made all the difference in runtime

    for i in range(len(nums)-k+1):
      if len(window_counter) == k and window_sum > max_subarray_sum:
        max_subarray_sum = window_sum        
      if (i+k) < len(nums):
        window_sum = window_sum - nums[i] + nums[i+k]
        if window_counter[nums[i]] == 1:
          del(window_counter[nums[i]])
        else:
          window_counter[nums[i]] -= 1
        window_counter[nums[i+k]] += 1
    return max_subarray_sum
