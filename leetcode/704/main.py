class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return -1
        low:int = 0
        high: int = len(nums) - 1
        while high - low > 1:
            mid: int = (high - low) // 2 + low
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                high = mid - 1
            else:
                low = mid + 1
        if nums[high] == target:
            return high
        if nums[low] == target:
            return low
        return -1
