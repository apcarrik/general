// func minPathSum(grid [][]int) int {
    
// }


func minPathSum(grid [][]int) int {
	m := len(grid) 
	n := len(grid[0])

	//initialize mpsGrid
	mpsGrid := [][]int{}
	for ri,row := range grid {
		mpsGrid = append(mpsGrid, []int{})
		for j:=0; j<len(row); j++ {
			mpsGrid[ri] = append(mpsGrid[ri], 0)
		}
	}

	// fill out mpsGrid one row at a time and one column at a time
	for ri, row:= range grid {
		for ci, _:= range row {			
			if ri==0 && ci == 0 {
				mpsGrid[ri][ci] = grid[ri][ci]
			} else if ri == 0 {
				mpsGrid[ri][ci] = grid[ri][ci] + mpsGrid[ri][ci-1]
			} else if ci == 0 {
				mpsGrid[ri][ci] = grid[ri][ci] + mpsGrid[ri-1][ci]
			} else {
				top := mpsGrid[ri-1][ci]
				left := mpsGrid[ri][ci-1]
				if top < left {
					mpsGrid[ri][ci] = grid[ri][ci] + top
				} else {
					mpsGrid[ri][ci] = grid[ri][ci] + left
				}
			}
		}
	}
	// return bottom right most element
	return mpsGrid[m-1][n-1]
	
}
