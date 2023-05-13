/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func searchRecursive(node, parent *TreeNode, val int, depth int) (*TreeNode, int) {
    if node == nil {
        return nil, depth
    } else if node.Val == val {
        return parent, depth
    }

    leftParent, leftDepth := searchRecursive(node.Left, node, val, depth+1)
    rightParent, rightDepth := searchRecursive(node.Right, node, val, depth+1)

    if leftParent != nil {
        return leftParent, leftDepth
    } 
    return rightParent, rightDepth
}

func isCousins(root *TreeNode, x int, y int) bool {
    // Idea: recursivley search for x and y, keeping track of parents and depth
    xParent, xDepth := searchRecursive(root, nil, x, 0)
    yParent, yDepth := searchRecursive(root, nil, y, 0)
    if xDepth == yDepth && xParent != yParent {
        return true
    }
    return false
}
