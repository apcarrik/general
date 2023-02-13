import (
    "fmt"
)

type Cell struct {
    X int
    Y int
    Usage int
}

type PathNode struct {
    Path []*Cell
    Next *PathNode
    Prev *PathNode
}

func shallowCopyPath(oldPath []*Cell) []*Cell{
    newPath := []*Cell{}
    for i:=0; i<len(oldPath); i++{
        newPath = append(newPath,oldPath[i])
    }
    return newPath
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
    
    // create map of 1 cells indexed by coordinates
    cellMap := map[int]map[int]*Cell{}
    for x:=0; x<m; x++{
        for y:=0; y<n; y++{
            if grid[x][y] == 1 {
                _, exists := cellMap[x]
                if !exists {
                    cellMap[x] = map[int]*Cell{}
                }
                    cellMap[x][y] = &Cell{X:x,Y:y}
            }
        }
    }
    fmt.Println("m:",m,"n:",n,"cellMap",cellMap)
    // create path data structures
    // pathSentinel := &PathNode{Path: nil, Next:nil, Prev:nil}
    pathsHead := &PathNode{Path: []*Cell{cellMap[0][0]}, Next: nil, Prev:nil}
    var finishedPathsHead *PathNode

    //edge case: m & n is 1
    // if m==1 && n==1 {
    //     return false
    // }
    numSplits := 0
    for i:=0; i<(m+n)*2; i++{
        // Iterate to next cell in path, or split path if both options avaliable
        fmt.Println("i:",i,"pathsHead",pathsHead)

        if pathsHead == nil {
            fmt.Println("All paths finished, moving on.")
            break
        }
        pathsCtr :=0
        for pathIter:=pathsHead; pathIter != nil; pathIter = pathIter.Next{
            pathsCtr += 1
            lastCell := pathIter.Path[len(pathIter.Path)-1]
            fmt.Println("pathIter",pathIter,"lastCell:",lastCell.X,",",lastCell.Y)
            if lastCell == cellMap[m-1][n-1]{
                // path ended on finish cell; place path in finishedPaths and remove from paths
                fmt.Println("placing path on finishedPaths")
                if pathIter == pathsHead {
                    pathsHead = pathsHead.Next
                    if pathIter.Next != nil {                        
                        pathIter.Next.Prev = pathsHead
                    }
                } else if pathIter.Next == nil {
                    pathIter.Prev.Next = pathIter.Next
                } else {
                    pathIter.Prev.Next = pathIter.Next
                    pathIter.Next.Prev = pathIter.Prev
                }

                pathIter.Prev = nil
                pathIter.Next = nil
                if finishedPathsHead == nil {
                    finishedPathsHead = pathIter
                }  else {
                    finishedPathsHead.Next = pathIter
                    pathIter.Prev = finishedPathsHead
                }
            } else {
                _, onesCellRightExists := cellMap[lastCell.X][lastCell.Y+1]
                _, onesCellDownExists := cellMap[lastCell.X+1][lastCell.Y]

                if onesCellRightExists && onesCellDownExists {
                    // split path
                    fmt.Println("splitting path")
                    numSplits++
                    prevPathNode := pathIter.Prev
                    newPath := shallowCopyPath(pathIter.Path)
                    newPath = append(newPath,cellMap[lastCell.X][lastCell.Y+1])
                    newPathNode := &PathNode{Path: newPath, Next: pathIter, Prev:prevPathNode}
                    pathIter.Path = append(pathIter.Path,cellMap[lastCell.X+1][lastCell.Y])     
                    pathIter.Prev = newPathNode      
                    if pathIter == pathsHead {
                        pathsHead = pathIter.Prev
                    }
                    fmt.Println(" pathIter.Prev",pathIter.Prev," pathIter.Prev.X,Y:",pathIter.Prev.Path[0].X, pathIter.Prev.Path[0].Y)
                    fmt.Println(" pathIter",pathIter," pathIter.Prev.X,Y:",pathIter.Prev.Path[0].X, pathIter.Prev.Path[0].Y)         
                } else if onesCellRightExists {
                    // move right
                    fmt.Println("moving right")
                    pathIter.Path = append(pathIter.Path,cellMap[lastCell.X][lastCell.Y+1])
                } else if onesCellDownExists {
                    // move down
                    fmt.Println("moving down")
                    pathIter.Path = append(pathIter.Path,cellMap[lastCell.X+1][lastCell.Y])
                } else {
                    // path ended before finish cell; remove path from paths
                    fmt.Println("removing path ", pathIter)
                    if pathIter.Next == nil && pathIter == pathsHead {
                        pathsHead = pathsHead.Next
                    } else if pathIter == pathsHead {
                        pathsHead = pathsHead.Next
                        pathIter.Next.Prev = pathsHead
                    } else if pathIter.Next == nil {
                        pathIter.Prev.Next = pathIter.Next
                    } else {
                        pathIter.Prev.Next = pathIter.Next
                        pathIter.Next.Prev = pathIter.Prev
                    }
                }
            }
        }
        fmt.Println("pathsCtr:",pathsCtr)

    }
    fmt.Println("numSplits:",numSplits)
    fmt.Println("-----\npathsHead",pathsHead)
    // Check how many paths exists 
    fmt.Println("finishedPathsHead",finishedPathsHead)
    switch  {
        case finishedPathsHead == nil:
            return true
        case finishedPathsHead.Next == nil:
            // check that there is at least 1 cell that is not beginning or end
            if len(finishedPathsHead.Path) > 2  {
                return true
            } else{
                return false
            }
    }
    // More than 1 path. Need to see if there is critical point that every path uses
    // Make freq map of cells on paths. check if any cell is on every path, if so return true, else false


    numPaths := 0
    for pathIter:=finishedPathsHead; pathIter != nil; pathIter = pathIter.Next{
        fmt.Println("pathIter",pathIter)
        numPaths += 1
        for i:=0; i< len(pathIter.Path); i++{
            pathIter.Path[i].Usage += 1
        }
    
    }
    fmt.Println("numPaths",numPaths)
    // Drop beginning/end path nodes
    delete(cellMap[0],0)
    delete(cellMap[m-1],n-1)

    for _,nArr := range cellMap {
        for _,cell := range nArr {
            fmt.Println("cell",cell)
            if cell.Usage == numPaths{ // if a 1 cell is used by all paths, it is critical cell and can disconnect matrix by being flipped
                return true
            }
        }
    }
    return false
}
