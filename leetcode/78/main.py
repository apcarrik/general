class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        superset: list[list[int]] = [[]]
        for num in nums:
            superset += [set+[num] for set in superset]
        return superset
