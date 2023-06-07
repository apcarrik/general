class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        n: int = len(nums)
        nth_exists: bool = False

        # cycle sort of nums
        i: int = 0
        while i <len(nums):
            num:int = nums[i]
            if num == n:
                nth_exists = True
                nums[i] = -1
            elif num != -1 and num != i:
                other: int = nums[num]
                if other == n:
                    nth_exists = True
                    nums[num] = num
                    nums[i] = -1
                else:
                    nums[num] = num
                    nums[i] = other
                    i -= 1
            i += 1
            
        
        # do linear search for -1 and return index it is at, otherwise return n
        if nth_exists:
            for i,num in enumerate(nums):
                if num == -1:
                    return i
        return n
