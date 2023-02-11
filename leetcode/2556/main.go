import (
    "strconv"
    //"fmt"
)

type PathNode struct {
    X int
    Y int
    Next *PathNode
}

// type QNode struct {
//     N *PathNode
//     Next *QNode
// }

func validCoordinates(x int, y int, m int, n int) bool {
    if x>=0 && x<m && y>=0 && y<n {
        return true
    }
    return false
}

func isPossibleToCutPath(grid [][]int) bool {
    // Idea: starting from (0,0), find all paths to (m-1,n-1) (use linked list for each path)
    // - If none exist, return true
    // - if only 1 exists, return true
    // - else, there are p paths. 
    //   - Keep frequency map of all 1 cells. If any use all p/p paths, then return true
    // all else, return false

    m := len(grid)
    n := len(grid[0])
    pathSentinel := &PathNode{X:-1, Y:-1, Next:nil}
    paths := []*PathNode{&PathNode{X:0, Y:0, Next:pathSentinel}} 

    //edge case: m & n is 1
    // if m==1 && n==1 {
    //     return false
    // }

    for i:=0; i<m+n; i++{

        // Check if all paths have finished
        allPathsFinished := true
        for j:=0; j<len(paths); j++{
            if paths[j].X != m-1 || paths[j].Y != n-1 {
                allPathsFinished = false
            }
        }
        if allPathsFinished {
            break
        }

        // Iterate to next cell in path, or split path if both options avaliable
        for j:=0; j<len(paths); j++{
            //fmt.Println("j:",j,"paths[j]:",paths[j])
            currX := paths[j].X
            currY := paths[j].Y
            if validCoordinates(currX+1,currY,m,n) && grid[currX+1][currY] == 1 && validCoordinates(currX,currY+1,m,n) && grid[currX][currY+1] == 1 {
                // split path
                //fmt.Println("splitting path")
                newPathNode1 := &PathNode{X:currX+1 , Y:currY , Next: paths[j]}
                newPathNode2 := &PathNode{X:currX , Y:currY+1 , Next: paths[j]}
                paths[j] = newPathNode1
                paths = append(paths,newPathNode2)
            } else if validCoordinates(currX+1,currY,m,n) && grid[currX+1][currY] == 1 {
                // move down
                //fmt.Println("moving down")
                newPathNode := &PathNode{X:currX+1 , Y:currY , Next: paths[j]}
                paths[j] = newPathNode
            } else if validCoordinates(currX,currY+1,m,n) && grid[currX][currY+1] == 1 {
                // move right
                //fmt.Println("moving right")
                newPathNode := &PathNode{X:currX , Y:currY+1 , Next: paths[j]}
                paths[j] = newPathNode
            } else if !(paths[j].X == m-1 && paths[j].Y==n-1) {
                // path ended before finish cell; remove from paths\
                //fmt.Println("removing path ", paths[j])
                if j == 0 {
                    paths = paths[1:]
                } else if j == len(paths)-1{
                    paths = paths[:len(paths)-1]
                } else {
                    paths = append(paths[:j], paths[j+1:]...)
                }
            }
        }

    }
    //fmt.Println(paths)
    // Check how many paths exists
    switch len(paths) {
        case 0:
            return true
        case 1:
            // check that there is at least 1 cell that is not beginning or end
            nonEndCellFound := false
            for paths[0] != pathSentinel{
                //fmt.Println(paths[0],"m:",m,"n:",n)
                if paths[0].X ==0 && paths[0].Y == 0 {
                    nonEndCellFound = nonEndCellFound
                } else if paths[0].X == m-1 && paths[0].Y ==n-1{
                    nonEndCellFound = nonEndCellFound
                } else {
                    //fmt.Println("nonEndCellFound")
                    nonEndCellFound = true
                }
                paths[0] = paths[0].Next
            }
            if nonEndCellFound {
                return true
            } else{
                return false
            }
    }
    // More than 1 path. Need to see if there is critical point that every path uses
    // Make freq map of cells on paths. check if any cell is on every path, if so return true, else false
    var freqMap = make(map[string]int)
    for j:=0; j<len(paths); j++{
        for paths[j] != pathSentinel {
            coordStr := "X" + strconv.Itoa(paths[j].X) + "Y" + strconv.Itoa(paths[j].Y)
            _, exists := freqMap[coordStr]
            if exists {
                freqMap[coordStr] += 1
            } else {
                freqMap[coordStr] = 1
            }
            paths[j] = paths[j].Next
        }
    }
    //fmt.Println(freqMap)

    for k,v := range freqMap {
        if  v == len(paths) && k != "X-1Y-1" && k != "X0Y0" && k != "X"+strconv.Itoa(m-1)+"Y"+strconv.Itoa(n-1) {
            //fmt.Println("returning true; k=", k, "v=",v)
            return true
        }
    }
    return false
}
