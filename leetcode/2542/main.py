class Solution:
  def _getN1Sum(self, n1FreqMap: Dict[int, int], k: int) -> int:
    n1Sum: int = 0
    for n1Key, n1Val in sorted(n1FreqMap.items(), reverse=True):
      if n1Val >= k:
        n1Sum += n1Key * k
        break
      else:
        n1Sum += n1Key * n1Val
        k -= n1Val
    return n1Sum

  def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:

    # Create n1FreqMap
    n1FreqMap: Dict[int, int] = {}
    for i, num1 in enumerate(nums1):
      if num1 in n1FreqMap:
        n1FreqMap[num1] += 1
      else:
        n1FreqMap[num1] = 1


    # Create n2Map (with frequency map of n1 values as value)
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
      n2Min: int = min(n2Map)
      n1Sum: int = self._getN1Sum(n1FreqMap, k)
      localMaxScore: int = n1Sum*n2Min
      if localMaxScore > globalMaxScore:
        globalMaxScore = localMaxScore

      # clean up for next loop iteration 
      for n1Key, n1Val in n2Map[n2Min].items():
        if n1FreqMap[n1Key] == n1Val:
          nums1Len -= n1FreqMap[n1Key]
          del n1FreqMap[n1Key]
        else:
          nums1Len -= n1Val
          n1FreqMap[n1Key] -= n1Val
      del n2Map[n2Min]

    return globalMaxScore
