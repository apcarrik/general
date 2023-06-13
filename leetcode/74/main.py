class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return -1
        low:int = 0
        high:int = len(matrix)*len(matrix[0]) -1
        while high-low > 1:
            mid: int = (high-low)//2 + low
            if self.matrixLookup(mid, matrix) == target:
                return True
            elif self.matrixLookup(mid, matrix) > target:
                high = mid - 1
            else:
                low = mid + 1
        if self.matrixLookup(high, matrix) == target or self.matrixLookup(low, matrix) == target:
            return True
        return False

    def matrixLookup(self, idx: int, matrix: list[list[int]]) -> int:
        # given index of element in matrix, returns the value of the element at that index
        m: int = len(matrix[0]) # col idx
        return matrix[idx // m][idx % m]
        
