
class Solution:

    def longestIncreasingPath(self, matrix: list[list[int]]) -> int:
        matrix = matrix
        rows: int = len(matrix)
        if rows == 0 or len(matrix[0]) == 0:
            return 0
        cols: int = len(matrix[0])

        longest: list[list[int]] = [[0 for _ in range(cols)] for _ in range(rows)]

        largest: list[tuple[int,int,int]] = \
          sorted([(matrix[r][c], r, c) for c in range(cols) for r in range(rows)], reverse=True)

        for _, r, c in largest:
            # check all 4 directions to find longest path for cell
            maxlen: int = 0
            if r > 0 and matrix[r-1][c] > matrix[r][c]:
                newlen: int = longest[r-1][c] + 1
                if newlen > maxlen:
                    maxlen = newlen
            if c > 0 and matrix[r][c-1] > matrix[r][c]:
                newlen: int = longest[r][c-1] + 1
                if newlen > maxlen:
                    maxlen = newlen
            if r < rows-1 and matrix[r+1][c] > matrix[r][c]:
                newlen: int = longest[r+1][c] + 1
                if newlen > maxlen:
                    maxlen = newlen
            if c < cols-1 and matrix[r][c+1] > matrix[r][c]:
                newlen: int = longest[r][c+1] + 1
                if newlen > maxlen:
                    maxlen = newlen 
            longest[r][c] = maxlen
        return max(map(max, longest))+1
