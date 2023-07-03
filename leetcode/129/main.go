/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    // Idea: traverse down using recursion, returning the sum of root-to-leaf paths of all leaves below
    if root.Left == nil && root.Right == nil {
        return  root.Val
    }
    localSum := 0
    if root.Left != nil {
        root.Left.Val = root.Val * 10 + root.Left.Val
        localSum += sumNumbers(root.Left)
    }
    if root.Right != nil {
        root.Right.Val = root.Val * 10 + root.Right.Val
        localSum += sumNumbers(root.Right)
    }
    return localSum
}
