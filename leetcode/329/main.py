# class Solution:
#     def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
class Solution:
    def dfs(self, r: int, c: int) -> int:
        # check if this cell's longest has been set
        if self.longest[r][c] != -1:
          return self.longest[r][c]

        # check all 4 directions
        maxlen: int = 0
        if r > 0 and self.matrix[r-1][c] > self.matrix[r][c]:
            newlen: int = self.dfs(r=r-1, c=c) + 1
            if newlen > maxlen:
                maxlen = newlen
        if c > 0 and self.matrix[r][c-1] > self.matrix[r][c]:
            newlen: int = self.dfs(r=r, c=c-1) + 1
            if newlen > maxlen:
                maxlen = newlen
        if r < self.rows-1 and self.matrix[r+1][c] > self.matrix[r][c]:
            newlen: int = self.dfs(r=r+1, c=c) + 1
            if newlen > maxlen:
                maxlen = newlen
        if c < self.cols-1 and self.matrix[r][c+1] > self.matrix[r][c]:
            newlen: int = self.dfs(r=r, c=c+1) + 1
            if newlen > maxlen:
                maxlen = newlen                
        return maxlen

    def longestIncreasingPath(self, matrix: list[list[int]]) -> int:
        self.matrix = matrix
        self.rows: int = len(self.matrix)
        if self.rows == 0 or len(self.matrix[0]) == 0:
            return 0
        self.cols: int = len(self.matrix[0])

        self.longest: list[list[int]] = [[-1 for _ in range(self.cols)] for _ in range(self.rows)]

        largest: list[tuple[int,int,int]] = \
          sorted([(self.matrix[r][c], r, c) for c in range(self.cols) for r in range(self.rows)], reverse=True)

        for tup in largest:
              _, r, c = tup
              # find longest path for cell
              self.longest[r][c] = self.dfs(r=r, c=c)
        return max(map(max, self.longest))+1
