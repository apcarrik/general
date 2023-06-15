class Solution:
    def subArrayRanges(self, nums: List[int]) -> int:
        sum: int = 0
        for i in range(0,len(nums)-1):
            min: list[int] = []
            max: list[int] = []
            heapq.heapify(min)
            heapq.heapify(max)
            heapq.heappush(min, nums[i])
            heapq.heappush(max, -1*nums[i])
            for j in range(i+1,len(nums)):
                heapq.heappush(min, nums[j])
                heapq.heappush(max, -1*nums[j])
                sum += (-1*max[0])-min[0]
        return sum
