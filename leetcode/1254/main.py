# class Solution:
#     def closedIsland(self, grid: List[List[int]]) -> int:
class Solution:
    def dfs(self, x:int, y:int):
        # turn this cell to water
        self.grid[x][y] = 1
        self.unchecked.discard((x,y))

        # check all 4 cells surrounding this (if they exist)
        if x > 0 and self.grid[x-1][y] == 0:
            self.dfs(x-1, y)
        if y < self.m-1 and self.grid[x][y+1] == 0:
            self.dfs(x,y+1)
        if x < self.n-1 and self.grid[x+1][y] == 0:
            self.dfs(x+1, y)
        if y > 0 and self.grid[x][y-1] == 0:
            self.dfs(x,y-1)

    def closedIsland(self, grid: list[list[int]]) -> int:
        self.grid = grid
        self.n: int = len(self.grid)
        if self.n == 0 or len(self.grid[0]) == 0:
            return 0
        self.m: int = len(self.grid[0])

        # create set of all unchecked cells
        self.unchecked: set[tuple[int,int]] = {(i,j) for i in range(self.n) for j in range(self.m)}

        # check border to remove all non-closed islands
        for x,y in [(0,j) for j in range(0,self.m)]\
            + [(i,self.m-1) for i in range(1,self.n)]\
            + [(self.n-1,j) for j in range(self.m-2,-1,-1)]\
            + [(i, 0) for i in range(self.n-2, 0, -1)]:
            if self.grid[x][y] == 0:
                self.dfs(x,y)
            self.unchecked.discard((x,y))

        # go through remaining unchecked cells and perform dfs to find all islands
        count: int = 0
        while len(self.unchecked) > 0:
            x,y = self.unchecked.pop()
            if self.grid[x][y] == 0:
                count += 1
                self.dfs(x,y)
            self.unchecked.discard((x,y))
        return count
