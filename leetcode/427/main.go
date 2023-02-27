/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

// X index is outer array, Y index is inner array

func constructRecursive(grid [][]int, xMin, xMax, yMin, yMax int) *Node {
    // Check if grid is 1x1; if so, return leaf node with value
    xDiff := xMax - xMin
    yDiff := yMax - yMin
    if xDiff == 0 {
        return &Node{
            Val: grid[xMin][yMin] != 0,
            IsLeaf: true,
        }
    }

    // Get all 4 quadrant values
    topLeft := constructRecursive(grid , xMin, xMin + xDiff/2, yMin, yMin + yDiff/2)
    topRight := constructRecursive(grid , xMin, xMin + xDiff/2, yMin + yDiff/2 + 1, yMax)
    bottomLeft := constructRecursive(grid , xMin + xDiff/2 + 1, xMax, yMin, yMin + yDiff/2)
    bottomRight := constructRecursive(grid , xMin + xDiff/2 + 1, xMax, yMin + yDiff/2 + 1, yMax)

    // Check if all 4 quadrants are leafs with the same values; if so, coalesce into one leaf node
    quads := []*Node{topLeft,topRight,bottomLeft,bottomRight}
    allQuadsAreLeavesAndSameVal := true
    leavesVal := false
    for i,quad := range quads {
        // check if all are leaves, and all have same value
        if !(*quad).IsLeaf {
            allQuadsAreLeavesAndSameVal = false
            break
        } else {
            if i == 0 {
                leavesVal = (*quad).Val
            } else if (*quad).Val != leavesVal {
                allQuadsAreLeavesAndSameVal = false
                break                
            }
        }
    }

    if allQuadsAreLeavesAndSameVal {
        return &Node{
            Val: leavesVal,
            IsLeaf: true,
        }
    }

    // Otherwise, return new non-leaf node with quadrants
    return &Node{
        IsLeaf: false,
        TopLeft: topLeft,
        TopRight: topRight,
        BottomLeft: bottomLeft,
        BottomRight: bottomRight,
    }

}

func construct(grid [][]int) *Node {
    // Idea: Recursivley call construct on quadrant of current bounds, keeping track of x&y min/max. Return once quadrant is only 1 large. Otherwise, get results from recursive steps and check if all values are the same. If so, coalesce current node into leaf node.
    val := constructRecursive(grid, 0, len(grid) - 1, 0, len(grid[0]) - 1)   
    return val
}
