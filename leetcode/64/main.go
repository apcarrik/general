// func minPathSum(grid [][]int) int {
    
// }


func minPathSum(grid [][]int) int {
	m := len(grid) 
	n := len(grid[0])

	// fill out grid one row at a time and one column at a time
	for ri, row:= range grid {
		for ci, _:= range row {			
			if ri==0 && ci == 0 {
				grid[ri][ci] = grid[ri][ci]
			} else if ri == 0 {
				grid[ri][ci] = grid[ri][ci] + grid[ri][ci-1]
			} else if ci == 0 {
				grid[ri][ci] = grid[ri][ci] + grid[ri-1][ci]
			} else {
				top := grid[ri-1][ci]
				left := grid[ri][ci-1]
				if top < left {
					grid[ri][ci] = grid[ri][ci] + top
				} else {
					grid[ri][ci] = grid[ri][ci] + left
				}
			}
		}
	}
	// return bottom right most element
	return grid[m-1][n-1]
	
}
