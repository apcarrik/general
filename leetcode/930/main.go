# class Solution:
#     def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:

class Solution:
    def triangleSum(self, n: int) -> int:
        return int(n*(n-1)/2)

    def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:
        subarray_count = 0
        if goal == 0:
            start = -1
            for i in range(len(nums)):
                if nums[i] == 0:
                    if start == -1:
                        start=i
                else:
                    if start != -1:
                        subarray_count += self.triangleSum(i-start+1)
                    start = -1
            if start != -1:
                subarray_count += self.triangleSum(len(nums)-start+1)

        else:
            rolling_sum = [nums[0]]
            rolling_sum_map = {nums[0]:0}
            for i in range(1,len(nums)):
                old_rolling_sum = rolling_sum[len(rolling_sum)-1]
                new_rolling_sum = old_rolling_sum + nums[i]
                rolling_sum.append(new_rolling_sum)
                if old_rolling_sum != new_rolling_sum:
                    rolling_sum_map[new_rolling_sum] = i
            rolling_sum_map[rolling_sum[-1]+1]=len(nums)

            for ni, num in enumerate(nums):
                base_sum = rolling_sum[ni]-num
                if rolling_sum_map.get(goal+base_sum+1) is not None:
                    subarray_count += rolling_sum_map[goal+base_sum+1] - rolling_sum_map[goal+base_sum]

        return subarray_count
