// import (
//     "fmt"
// )

type Node struct {
    X int
    Y int
    Next *Node
}

// func manhattanDist(n1 *Node, n2 *Node) int {
//     xDelta := n1.X-n2.X
//     if xDelta < 0 {
//         xDelta *= -1
//     }
//     yDelta := n1.Y-n2.Y
//     if yDelta < 0 {
//         yDelta *= -1
//     }
//     return xDelta + yDelta
// }

func inBounds(x int, y int, n int) bool {
    if x >=0 && x<n && y>=0 && y<n {
        return true
    }
    return false
}

// func manhattanSweep(w *Node, grid [][]int) int {
//     // test concentric diamonds around w cell
//     n := len(grid)
//     x := w.X
//     y := w.Y
//     //fmt.Println("x=",x,"y=",y)
//     for offset:=1; offset<=n+n; offset++{
//         // check left/right/above/below corners
//         if  (inBounds(x-offset,y,n) && grid[x-offset][y] == 1) ||
//             (inBounds(x+offset,y,n) && grid[x+offset][y] == 1) ||
//             (inBounds(x,y+offset,n) && grid[x][y+offset] == 1) ||
//             (inBounds(x,y-offset,n) && grid[x][y-offset] == 1) {
//             //fmt.Println("offset=",offset)
//             return offset
//         }

//         // check others in diamond
//         for xOff:=offset-1; xOff>0; xOff-- {
//             if  (inBounds(x-xOff,y-(offset-xOff),n) && grid[x-xOff][y-(offset-xOff)] == 1) ||
//                 (inBounds(x+xOff,y-(offset-xOff),n) && grid[x+xOff][y-(offset-xOff)] == 1) ||
//                 (inBounds(x-xOff,y+(offset-xOff),n) && grid[x-xOff][y+(offset-xOff)] == 1) ||
//                 (inBounds(x+xOff,y+(offset-xOff),n) && grid[x+xOff][y+(offset-xOff)] == 1) {
//                 //fmt.Println("xOff=",xOff,"offset=",offset)
//                 return offset
//             }
//         }
//     }
//     //fmt.Println("default")
//     return n+n
// }

func maxDistance(grid [][]int) int {

    // Official solution: multi-source BFS. Using each land cell as a source, search each simeltaneously for increasing manhattan distance to find water cells. Keep track of all the water cells you find until you find all of them, and denote the manhattan distance of the last water cell visited.

    n := len(grid)
    //  initialize grid2 to be copy of grid
    grid2 := grid
    sentinel := Node{-1,-1,nil}
    landQ := &sentinel

    // iterate through all cells, adding any land cells to queue
    waterExists := false
    for x:=0; x<n; x++ {
        for y:=0; y<n; y++{
            newNode := &Node{X:x, Y:y,Next: nil}
            if grid[x][y] == 1 {
                newNode.Next = landQ
                landQ = newNode
            } else {
                waterExists = true
            }
        }
    }

    // Check if no land or water cells
    if landQ.Next == nil || !waterExists {
        return -1
    }

    // Perform BFS with multi-sources, using each land cell as a source
    distance := 1
    for distance = 1; distance<n+n; distance++{
        if landQ.Next == nil {
            break
        }
        newLandQ := &sentinel
        for landQ.Next != nil{
            // left
            newX := landQ.X+1
            newY := landQ.Y
            if inBounds(newX,newY,n) && grid2[newX][newY] == 0 {
                grid2[newX][newY] = 1
                newNode := &Node{X: newX, Y: newY, Next: newLandQ}
                newLandQ = newNode
            }
            // right
            newX = landQ.X-1
            newY = landQ.Y
            if inBounds(newX,newY,n) && grid2[newX][newY] == 0 {
                grid2[newX][newY] = 1
                newNode := &Node{X: newX, Y: newY, Next: newLandQ}
                newLandQ = newNode
            }
            // above
            newX = landQ.X
            newY = landQ.Y-1
            if inBounds(newX,newY,n) && grid2[newX][newY] == 0 {
                grid2[newX][newY] = 1
                newNode := &Node{X: newX, Y: newY, Next: newLandQ}
                newLandQ = newNode
            }
            // below
            newX = landQ.X
            newY = landQ.Y+1
            if inBounds(newX,newY,n) && grid2[newX][newY] == 0 {
                grid2[newX][newY] = 1
                newNode := &Node{X: newX, Y: newY, Next: newLandQ}
                newLandQ = newNode
            }
            landQ = landQ.Next
            
        }
        landQ = newLandQ
        // check for water nodes left
        waterExists := false
        for x:=0; x<n; x++{
            for y:=0; y<n; y++{
                if grid2[x][y] == 0 {                    
                    waterExists = true
                }
            }
        }
        if !waterExists {
            break
        }
        //fmt.Println(distance,"\n", grid2)
    } 
    return distance
}
