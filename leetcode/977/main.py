# class Solution:
#     def sortedSquares(self, nums: List[int]) -> List[int]:

class Solution:
  def sortedSquares(self, nums: list[int]) -> list[int]:
    # Find element closest to 0
    smallest_abs_val: float = math.inf
    smallest_abs_val_idx: int= -1
    for i,num in enumerate(nums):
      if float(abs(num)) <= smallest_abs_val:
        smallest_abs_val = float(abs(num))
        smallest_abs_val_idx = i
      else:
        break

    # Now, iterate pointers left and right
    left: int = smallest_abs_val_idx
    right: int = smallest_abs_val_idx+1
    squared: list[int] = []
    while left >= 0 or right < len(nums):
      if left < 0:
        squared.append(nums[right]**2)
        right += 1
      elif right >= len(nums):
        squared.append(nums[left]**2)
        left -= 1
      else:
        if abs(nums[left]) < abs(nums[right]):
          squared.append(nums[left]**2)
          left -= 1				
        else:
          squared.append(nums[right]**2)
          right += 1
    return squared
