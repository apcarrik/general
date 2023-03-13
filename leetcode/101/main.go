/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetricRecursive(left, right *TreeNode) bool {
    switch {
        case left == nil && right == nil:
            return true
        case left != nil && right != nil:
            if left.Val != right.Val{
                return false
            }
            return isSymmetricRecursive(left.Left, right.Right) && isSymmetricRecursive(left.Right, right.Left)
    }
    return false
}

func isSymmetric(root *TreeNode) bool {
    return isSymmetricRecursive(root.Left, root.Right)
}
