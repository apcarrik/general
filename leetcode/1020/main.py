from collections import deque
class Solution:
    def numEnclaves(self, grid: list[list[int]]) -> int:
        n: int = len(grid) # max height
        if n==0 or len(grid[0]) == 0:
            return 0
        m: int = len(grid[0]) # max width

        unexplored: set[tuple[int,int]] = set([(i,j) for i in range(n) for j in range(m) ])
        # iterate through all border pieces, following all connected land
        for b in [(0,j) for j in range(m)] \
        + [(i,m-1) for i in range(1,n)] \
        + [(n-1, j) for j in range(m-2, -1, -1)] \
        + [(i, 0) for i in range(n-2, 0, -1)]:
            # print("checking, b=", b)
            if b in unexplored:
                if grid[b[0]][b[1]] == 1:
                    # BFS to explore land
                    q: deque[tuple[int,int]] = deque([b])
                    while len(q) > 0:
                        c = q.popleft()
                        # print("in BFS, c=", c)
                        # check all four directions around c
                        if c[0] > 0 and (c[0]-1,c[1]) in unexplored and grid[c[0]-1][c[1]] == 1: # top
                            q.append((c[0]-1,c[1]))
                            unexplored.discard((c[0]-1,c[1]))

                        if c[1] < m-1 and (c[0],c[1]+1) in unexplored and grid[c[0]][c[1]+1] == 1: # right
                            q.append((c[0],c[1]+1))
                            unexplored.discard((c[0],c[1]+1))

                        if c[0] < n-1 and (c[0]+1,c[1]) in unexplored and grid[c[0]+1][c[1]] == 1: # bottom
                            q.append((c[0]+1,c[1]))
                            unexplored.discard((c[0]+1,c[1]))

                        if c[1] > 0 and (c[0],c[1]-1) in unexplored and grid[c[0]][c[1]-1] == 1: # left
                            q.append((c[0],c[1]-1))
                            unexplored.discard((c[0],c[1]-1))
                        # print(" unexplored=", unexplored)
            if b in unexplored: unexplored.discard(b)

        # iterate through all remaining pieces in unexplored, counting island land pieces
        # print("after border, unexplored:",unexplored)
        count: int = 0
        while len(unexplored) > 0:
            i,j = unexplored.pop()
            if grid[i][j] == 1:
                count += 1
        # print(count)
        return count
