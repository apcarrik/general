class Solution:
    def findUnsortedSubarray(self, nums: List[int]) -> int:
        nums2: list[int] = sorted(nums)
        n: int = len(nums)

        i1: int = -1
        for i in range(n):
            if nums[i] != nums2[i]:
                i1 = i
                break
        if i1 == -1:
            return 0

        i2: int = -1
        for i in range(n-1,-1,-1):
            if nums[i] != nums2[i]:
                i2 = i
                break
        return i2-i1+1
