class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        if len(points) == 0 or len(points[0]) == 0:
            return 0

        cols: int = len(points)
        if cols == 1:
            return max(points[0])

        rows: int = len(points[0])
        if rows == 1:
            return sum([c[0] for c in points])

        # if here, there are at least 2 rows and at least 2 columns
        for c in range(cols-2,-1,-1):
            for r1 in range(rows):
                maxpts: int = points[c][r1] + points[c+1][0] - abs(r1)
                for r2 in range(1,rows):
                    cost: int = points[c][r1] + points[c+1][r2] - abs(r1-r2)
                    if cost > maxpts:
                        maxpts = cost
                points[c][r1] = maxpts
        return max(points[0])
