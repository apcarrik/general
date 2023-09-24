// func closedIsland(grid [][]int) int {
    
// }

type Island struct {
	Squares [][2]int
	Closed bool	
}

func closedIsland(grid [][]int) int {
	xmax := len(grid[0])-1
	ymax := len(grid)-1
	islands := map[[2]int]*Island{}
	for y:=0; y<len(grid); y++{
		for x:=0; x<len(grid[0]); x++ {
			if grid[y][x] == 0 {	
				coords := [2]int{y,x}	
				// create new island	
				islands[coords] = &Island{
					Squares: [][2]int{coords},
					Closed: true,
				}

				// check if square is on boundary
				if coords[0] == 0 || coords[0] == ymax || coords[1] == 0 || coords[1] == xmax {					
          islands[coords].Closed = false
				}
				
				// check square above this if applicable
				if y > 0 {
					topCoords := [2]int{y-1,x}
					if _,ok := islands[topCoords]; ok {
						// coalesce this island to above island
						if islands[topCoords].Closed == false || islands[coords].Closed == false {
            	islands[topCoords].Closed = false
							islands[coords].Closed = false
							fmt.Println(islands[coords],islands[topCoords])
						}
						for _,square := range islands[coords].Squares {
					 		islands[topCoords].Squares = append(islands[topCoords].Squares,square)
            }
						islands[coords] = islands[topCoords]
					}
				}
				// check square to left of this if applicable
				if x > 0 {
					leftCoords := [2]int{y,x-1}
					if _,ok := islands[leftCoords]; ok && islands[coords] != islands[leftCoords]  {
						// coalesce this island to above island		
						if islands[leftCoords].Closed == false || islands[coords].Closed == false {
            	islands[leftCoords].Closed = false
							islands[coords].Closed = false
						}
						for _,square := range islands[coords].Squares {
							islands[leftCoords].Squares = append(islands[leftCoords].Squares, square)
							islands[square] = islands[leftCoords]
            }
						islands[coords] = islands[leftCoords]
					}
				}
			}
		}
	}

	// go through islands and return the number of closed islands
	uniqueIslands := map[*Island]bool{}
	islandCount := 0
	for _,island := range islands {
		if island.Closed {
			if _,ok := uniqueIslands[island]; !ok {
				uniqueIslands[island] = true
				islandCount++
			}
		}
	}
	for island := range uniqueIslands{
		fmt.Println(island)
	}
	return islandCount
}
