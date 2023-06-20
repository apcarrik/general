# class Solution:
#     def isThereAPath(self, grid: List[List[int]]) -> bool:
import heapq
class Solution:
    def isThereAPath(self, grid: list[list[int]]) -> bool:
        rows: int = len(grid)
        if rows == 0 or len(grid[0]) == 0:
            return False
        cols: int = len(grid[0])
        if (rows+cols)%2 == 0 or rows < 2 or cols < 2: # 2x3 or 3x2 minimum size
            return False # we would have an odd number of cells on path
            
        def addMove(old: tuple[int,int], moves:list[tuple[int,tuple[int,int],tuple[int,int]]], increase: bool):
            new: tuple[int,int] = (old[0]+1,old[1]-1)
            if increase:
                if grid[old[0]][old[1]] == 0:
                    if grid[new[0]][new[1]] == 1:
                        heapq.heappush(moves,(-4,old,new))
                    else:
                        heapq.heappush(moves,(-2,old,new))
                else:
                    if grid[new[0]][new[1]] == 1:
                        heapq.heappush(moves,(-3,old,new))
                    else:
                        heapq.heappush(moves,(-1,old,new))
            else:
                if grid[old[0]][old[1]] == 1:
                    if grid[new[0]][new[1]] == 0:
                        heapq.heappush(moves,(-4,old,new))
                    else:
                        heapq.heappush(moves,(-2,old,new))
                else:
                    if grid[new[0]][new[1]] == 0:
                        heapq.heappush(moves,(-3,old,new))
                    else:
                        heapq.heappush(moves,(-1,old,new))
            
        path_sum: int = sum(grid[0]) + sum([grid[i][cols-1] for i in range(1,rows)])
        desired: int = (rows+cols-1)//2
        path: list[list[int]] = [[0 for _ in range(cols)] for _ in range(rows)]
        path[0] = [1 for _ in range(cols)]
        for i in range(rows):
            path[i][cols-1] = 1

        # calculate first move value
        first_val: int = 0
        if path_sum == desired:
            return True
        elif path_sum < desired:
            if grid[0][cols-1] == 0:
                if grid[1][cols-2] == 1:
                    first_val = -4
                else:
                    first_val = -2
            else:
                if grid[1][cols-2] == 1:
                    first_val = -3
                else:
                    first_val = -1
        else:
            if grid[0][cols-1] == 1:
                if grid[1][cols-2] == 0:
                    first_val = -4
                else:
                    first_val = -2
            else:
                if grid[1][cols-2] == 0:
                    first_val = -3
                else:
                    first_val = -1

        moves: list[tuple[int,tuple[int,int],tuple[int,int]]] = [(first_val,(0,cols-1),(1,cols-2))]
        while len(moves) > 0:
            # pop move and execute
            move:tuple[int,tuple[int,int],tuple[int,int]] = heapq.heappop(moves)
            val, old, new = move
            if path_sum < desired:
                if val == -4:
                    path_sum += 1
                elif val == -1:
                    path_sum -= 1
            else:
                if val == -4:
                    path_sum -= 1
                elif val == -1:
                    path_sum += 1
            path[old[0]][old[1]] = 0
            path[new[0]][new[1]] = 1

            # check exit condition
            if path_sum == desired:
                return True

            # find next moves
            if old[0] < rows-1 and old[1] > 1 and path[old[0]][old[1]-2] == 1:
                # we can move left shoulder
                addMove((old[0],old[1]-1),moves, path_sum < desired)
            if old[0] < rows-2 and old[1] > 0 and path[old[0]+2][old[1]] == 1:
                # we can move right shoulder
                addMove((old[0]+1,old[1]),moves, path_sum < desired)
        return False
