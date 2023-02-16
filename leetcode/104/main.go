/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    // Idea: recusivley check the depth of all children. Return 1+ the greater of the two depths
    if root == nil {
        return 0
    }
    leftDepth, rightDepth := 0, 0
    if root.Left != nil {
        leftDepth = maxDepth(root.Left)
    }
    if root.Right != nil {
        rightDepth = maxDepth(root.Right)
    }
    if leftDepth > rightDepth {
        return 1+leftDepth
    }
    return 1+rightDepth
}
