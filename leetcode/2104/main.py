class Solution:
    def subArrayRanges(self, nums: List[int]) -> int:
        sum: int = 0
        for i in range(0,len(nums)-1):
            min: int = nums[i]
            max: int = nums[i]
            for j in range(i+1,len(nums)):
                if nums[j] < min:
                    min = nums[j] 
                elif nums[j] > max:
                    max = nums[j] 
                sum += max-min
        return sum
