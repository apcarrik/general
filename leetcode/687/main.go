/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  import "fmt"


func findArrowsRec(root *TreeNode, maxLenIn int) (arrowLen, maxLen int) {
    maxLen = maxLenIn
    arrowLen = 0
    if root == nil {
        return
    }

    left, maxLenLeft := findArrowsRec(root.Left, maxLen)
    right, maxLenRight := findArrowsRec(root.Right, maxLen)
    arrowLeft, arrowRight := 0, 0
    if root.Left != nil && root.Val == root.Left.Val {
        arrowLeft = left + 1
    }
    if root.Right != nil && root.Val == root.Right.Val {
        arrowRight = right + 1
    }

    // Update arrowLen from recursive calls
    if arrowLeft > arrowRight {
        arrowLen = arrowLeft
    } else {
        arrowLen = arrowRight
    }

    // Update maxLen from recursive calls and arrows
    if maxLenLeft > maxLen {
        maxLen = maxLenLeft
    }
    if maxLenRight > maxLen {
        maxLen = maxLenRight
    }
    if arrowLeft+arrowRight > maxLen {
        maxLen = arrowLeft+arrowRight
    }
    // fmt.Println("Root:",root,", arrowLen:",arrowLen,", maxLen:", maxLen, ", left:", left, ", maxLenLeft:", maxLenLeft,", right:", right, ", maxLenRight:", maxLenRight)
    return
}

func longestUnivaluePath(root *TreeNode) int {
    // Official Solution: Recursivley search for longest path on either child if this node has the same value, and if both children have same value add their path lengths together

    _, maxLen := findArrowsRec(root,0)
    return maxLen
    
}
