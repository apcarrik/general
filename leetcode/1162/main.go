// import (
//     "fmt"
// )

type Node struct {
    X int
    Y int
    Next *Node
}

func manhattanDist(n1 *Node, n2 *Node) int {
    // TOOD: calculate manhattan distance
    xDelta := n1.X-n2.X
    if xDelta < 0 {
        xDelta *= -1
    }
    yDelta := n1.Y-n2.Y
    if yDelta < 0 {
        yDelta *= -1
    }
    return xDelta + yDelta
}

func maxDistance(grid [][]int) int {
    // Idea: make a list of all water cell coordinates. Make another list with all land cells, and check for each water cell what the manhattan distance is to each land cell, and find the minimum of this. Then, find the maximum of the minimum distance to land of each water cell.

    sentinel := Node{-1,-1,nil}
    water := &sentinel
    land := &sentinel
    n := len(grid)

    // iterate through all cells, classifying each as water or land
    for x:=0; x<n; x++ {
        for y:=0; y<n; y++{
            newNode := &Node{x,y,nil}
            if grid[x][y] == 0 {
                newNode.Next = water
                water = newNode
            } else {
                newNode.Next = land
                land = newNode
            }
        }
    }

    // TODO: optimization - check if no land cells or no water cells

    // iterate through each water cell, checking each land cell's distance
    globalMaxDist := -1
    for water.Next != nil {
        localMinDist := n+n
        landIter := land
        for landIter.Next != nil {
            manhattan := manhattanDist(water, landIter)
            //fmt.Println("water:[",water.X,",",water.Y,"], land:[",landIter.X,",",landIter.Y,"], manhattan:",manhattan)
            if manhattan < localMinDist {
                localMinDist = manhattan
                //fmt.Println("localMinDist=",localMinDist)
            }
            landIter = landIter.Next
        }
        if localMinDist < n+n && localMinDist > globalMaxDist {
            globalMaxDist = localMinDist
            //fmt.Println("globalMaxDist=",globalMaxDist)
        }
        water = water.Next
    }

    return globalMaxDist // TODO: Test
}
