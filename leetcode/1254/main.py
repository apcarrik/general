# class Solution:
#     def closedIsland(self, grid: List[List[int]]) -> int:
class Solution:
    def dfs(self, x:int, y:int) -> bool:
        # Return true only if this and all connected land doesn't touch border            
        if x < 0 or y < 0 or x >= self.n or y >= self.m:
            return False

        if self.grid[x][y] == 1:
            return True

        self.grid[x][y] = 1
        top: bool = self.dfs(x-1,y)
        right: bool = self.dfs(x,y+1)
        down: bool = self.dfs(x+1,y)
        left: bool = self.dfs(x,y-1)

        return top and right and down and left

    def closedIsland(self, grid: list[list[int]]) -> int:
        self.grid = grid
        self.n: int = len(self.grid)
        if self.n == 0 or len(self.grid[0]) == 0:
            return 0
        self.m: int = len(self.grid[0])

        # go through all cells, calling dfs if they are land, and incrementing count if dfs returns true, indicating no boundaries encountered
        count: int = 0
        for x in range(self.n):
            for y in range(self.m):            
                if self.grid[x][y] == 0 and self.dfs(x,y):
                    count += 1
        return count
