from bisect import bisect_left
class Solution:
    def search(self, nums: list[int], tgt: int) -> int:
        if len(nums) == 0:
            return -1
        if len(nums) == 1:
            return 0 if nums[0] == tgt else -1
        n:int = len(nums)

        # 1. determine if nums was rotated
        def decreaseBinsearch(a: list[int]) -> int:
            # returns the index of the element that is < previous element, or -1
            # if it does not exist.
            lo: int = 1
            hi: int = len(a)-1
            lastval: int = a[-1]
            while hi > lo:
                mid: int = (hi-lo)//2 + lo
                if a[mid] < a[mid-1]:
                    return mid
                elif a[mid] > lastval:
                    # move right
                    lo = mid + 1
                else:
                    hi = mid - 1

            if a[lo] < a[lo-1]:
                return lo
            return -1
        #  set r_idx to the index of the first element in the 2nd subarray
        r_idx: int = decreaseBinsearch(nums)

        # 2. search for tgt in nums
        min: int = 0
        max: int = n-1
        if r_idx != -1:
            # nums rotated, determine which subarray to search
            if nums[-1] == tgt:
                return n-1
            elif nums[-1] > tgt:
                min = r_idx
            else:
                max = r_idx-1
        # perform binsearch on nums[min:max+1]
        bl: int = bisect_left(a=nums, x=tgt, lo=min, hi=max)
        print(bl, n, nums[bl] if bl < n else "-",tgt)
        return bl if nums[bl] == tgt else -1
