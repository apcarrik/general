class Solution:
  def circularArrayLoop(self, nums: list[int]) -> bool:
    visited: set[int] = set()
    n: int = len(nums)
    for i in range(len(nums)):
      if i in visited:
        continue
      slow: int = i
      fast: int = (i + nums[i]) % n
      while slow != fast:
        fast = (fast + nums[fast]) % n
        fast = (fast + nums[fast]) % n
        slow = (slow + nums[slow]) % n

      # now, fast=slow, so we just iterate through loop, adding new elements to visited
      first_direction: bool = False if nums[fast] < 0 else True
      consistent: bool = True
      loop_length: int = 0
      while fast not in visited:
        visited.add(fast)
        if (nums[fast] < 0 and first_direction == True) or  (nums[fast] >= 0 and first_direction == False):
          consistent = False
        fast = (fast + nums[fast]) % n
        loop_length += 1
      if consistent and loop_length > 1:
        return True

      # now, loop has been added to visited. need to add the tail as well
      fast = i
      while fast not in visited:
        visited.add(fast)
        fast =  (fast + nums[fast]) % n
    return False
