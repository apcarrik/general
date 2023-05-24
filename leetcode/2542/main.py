from heapq import nlargest

class Solution:
  def _getN1s(n1FreqMap: Dict[int, int], k: int) -> int:
    n1s = 0
    # TODO
    return n1s

  def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:

    # Create n1FreqMap
    n1FreqMap: Dict[int, int] = {}
    for i, num1 in enumerate(nums1):
      if num1 in n1FreqMap:
        n1FreqMap[num1] += 1
      else:
        n1FreqMap[num1] = 1


    # Create n2Map
    n2Map: Dict[int, Dict[int,int]] = {}
    for i, num2 in enumerate(nums2):
      if num2 in n2Map:
        if nums1[i] in n2Map[num2]:
          n2Map[num2][nums1[i]] += 1
        else:
          n2Map[num2][nums1[i]] = 1
      else:
        n2Map[num2] = {nums1[i]: 1}

    # Find max score using smallest n2 while there are at least k elements in nums1
    globalMaxScore: int = 0
    nums1Len = len(nums1)
    while nums1Len >= k:
      n2: int = min(n2Map)
      n1s: int = self._getN1s(n1FreqMap, k)
      localMaxScore: int = sum(n1s)*n2
      if localMaxScore > globalMaxScore:
        globalMaxScore = localMaxScore

      # clean up for next loop iteration 
      # TODO: update this
      for n1Val in n2Map[n2]:
        nums1.remove(n1Val)
      del n2Map[n2]

    return globalMaxScore
