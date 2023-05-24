class Solution:
  def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:

    # Create n2Map
    n2Map: Dict[int, List[int]] = {}
    for i, num2 in enumerate(nums2):
      if num2 in n2Map:
        n2Map[num2].append(nums1[i])
      else:
        n2Map[num2] = [nums1[i]]

    # Sort nums1
    nums1.sort(reverse=True)

    # Find max score using smallest n2 while there are at least k elements in nums1
    globalMaxScore: int = 0
    while len(nums1) >= k:
      n2: int = min(n2Map)
      localMaxScore: int = sum(nums1[:k])*n2
      if localMaxScore > globalMaxScore:
        globalMaxScore = localMaxScore

      # clean up for next loop iteration
      for n1Val in n2Map[n2]:
        nums1.remove(n1Val)
      del n2Map[n2]

    return globalMaxScore
