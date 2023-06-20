class Solution:
    def findUnsortedSubarray(self, nums: List[int]) -> int:
        n: int = len(nums)
        if n <= 1:
            return 0
            
        min_seen: list[int] = [nums[-1]]
        for i in range(n-2,-1,-1):
            min_seen.append(min(min_seen[-1],nums[i]))
        i1: int = -1
        for i in range(n):
            if nums[i] != min_seen[n-1-i]:
                i1 = i
                break
        if i1 == -1:
            return 0

        max_seen: list[int] = [nums[0]]
        for i in range(1,n):
            max_seen.append(max(max_seen[-1],nums[i]))
        i2: int = -1
        for i in range(n-1,-1,-1):
            if nums[i] != max_seen[i]:
                i2 = i
                break
        return i2-i1+1
