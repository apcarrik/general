# class Solution:
#     def sortedSquares(self, nums: List[int]) -> List[int]:

class Solution:
  def sortedSquares(self, nums: list[int]) -> list[int]:
    return sorted(list(map(lambda num: num ** 2, nums)))
