/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRecursive(root *TreeNode, sum int) int {
    if root.Left == nil && root.Right == nil {
        return (sum*10) + root.Val
    }
    localSum := 0
    if root.Left != nil {
        localSum += sumRecursive(root.Left, (sum*10) + root.Val)
    }
    if root.Right != nil {
        localSum += sumRecursive(root.Right, (sum*10) + root.Val)
    }
    return localSum
}

func sumNumbers(root *TreeNode) int {
    // Idea: traverse down using recursion, returning the sum of root-to-leaf paths of all leaves below
    return sumRecursive(root, 0)
}
