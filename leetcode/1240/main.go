/* Notes:
    - There are a finite number of ways to fill a rectangle with integer squares
    - however, the number of ways will grow very large, esp. considering the may places to place them
    - Idea: start with the largest possible square and then place the next largest square possible, prioritizing corners first. Iterate through starting with the next largest possible square, ensuring that all later squares are smaller than or equal to the current
        - better idea: when placing square, find all possible locations and prefer one that has the fewest open sides

    Updated algorithm:
        Starting with the largest possible square, place as many squares of this size as possible, each optimally in the grid. Then, attempt to place as many as the next smallest size square as possible, on down until you have a 1x1 square. Keep track of the number of squares in the grid, starting with a grid of all 1x1 squares and subtracting however many 1x1 squares your current square takes up when placed (minus 1 for the square itself).
        Placing a square in the grid optimally means placing it in a way that minimizes the number of sides adjacent to unoccupied cells. So, you'll need to check all possible placements and find the one with the minimum number of 'open' sides (note the bounds of grid are closed sides).


*/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func resetSpaces(spacesPtr *[][]bool) {
    spaces := *spacesPtr
    n := len(spaces)
    m := len(spaces[0])
    for i:=0; i<n; i++{
        for j:=0; j<m; j++{
            spaces[i][j] = false
        }
    }
}

func canPlaceSquare(squareSize int, topLeft [2]int, spacesPtr *[][]bool) bool {
    spaces := *spacesPtr
    n := len(spaces)
    m := len(spaces[0])
    if topLeft[0]+squareSize > n || topLeft[1]+squareSize > m {
        return false
    }
    for y := topLeft[0]; y < topLeft[0] + squareSize; y++ {
        for x := topLeft[1]; x < topLeft[1] + squareSize; x++ {
            if spaces[y][x] == true {
                return false
            }
        }
    }
    return true
}

func updateSpaces(squareSize int, topLeft [2]int, spacesPtr *[][]bool) {
    spaces := *spacesPtr
    for y := topLeft[0]; y < topLeft[0] + squareSize; y++ {
        for x := topLeft[1]; x < topLeft[1] + squareSize; x++ {
            spaces[y][x] = true
        }
    }
    return
}

func placeSquare(squareSize int, spacesPtr *[][]bool) bool {
    spaces := *spacesPtr
    n := len(spaces)
    m := len(spaces[0])

    // TODO: update this with the new logic to find all possible places to place square and choose the one with the smallest number of open sides after placement
    
    // Check four corners for placement
    if spaces[0][0] == false && canPlaceSquare(squareSize, [2]int{0,0}, &spaces) {
        updateSpaces(squareSize, [2]int{0,0}, &spaces)
        return true
    }
    if spaces[n-1][0] == false && canPlaceSquare(squareSize, [2]int{n-squareSize,0}, &spaces) {
        updateSpaces(squareSize, [2]int{n-squareSize,0}, &spaces)
        return true
    } 
    if spaces[n-1][m-1] == false && canPlaceSquare(squareSize, [2]int{n-squareSize,m-squareSize}, &spaces) {
        updateSpaces(squareSize, [2]int{n-squareSize,m-squareSize}, &spaces)
        return true
    } 
    if spaces[0][m-1] == false && canPlaceSquare(squareSize, [2]int{0,m-squareSize}, &spaces) {
        updateSpaces(squareSize, [2]int{0,m-squareSize}, &spaces)
        return true
    }

    // now, check all grid for available spots
    for y:=0; y<n; y++{
        for x:=0; x<m; x++{
            if spaces[y][x] == false && canPlaceSquare(squareSize, [2]int{y,x}, &spaces) {
                updateSpaces(squareSize, [2]int{y,x}, &spaces)
                return true
            }
        }
    }

    return false
}

func tilingRectangle(n int, m int) int {

    // Create 2d slice of occupied spaces
    spaces := [][]bool{}
    for i:=0; i<n; i++{
        spaces = append(spaces,[]bool{})
        for j:=0; j<m; j++{
            spaces[i] = append(spaces[i], false)
        }
    }

    // iterate through largest square down to 1, filling in rectangle with as large of squares as possible
    minSquaresUsed := n*m
    for largestSquare := min(n,m); largestSquare > 1; largestSquare -- {
        squaresUsed := n*m
        for squareSize := largestSquare; squareSize > 1; squareSize-- {
            // attempt to place this square
            for placeSquare(squareSize, &spaces){
                squaresUsed -= (squareSize*squareSize-1)
                fmt.Println("placed square of size ", squareSize, "squares used", squaresUsed)
            }

            // if no way to place this square, go to next smallest

        }
        fmt.Println("squares used:", squaresUsed)
        
        // check if number of rectangles used is less than minSquaresUsed
        if squaresUsed < minSquaresUsed {
            minSquaresUsed = squaresUsed
        }

        // reset spaces
        resetSpaces(&spaces)
    }

    return minSquaresUsed
}
