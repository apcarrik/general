/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /* Idea: traverse through tree using BFS, adding any left leaves as you go. Use a Queue to implement BFS.

 */

type QNode struct {
    TN *TreeNode
    Next *QNode
}

func sumOfLeftLeaves(root *TreeNode) int {

    nodeQ := &QNode{TN: root}
    nodeQTail := nodeQ
    leftSum := 0
    for nodeQ != nil {
        node := nodeQ.TN
        if node != nil {
            if node.Left != nil {
                if node.Left.Left == nil && node.Left.Right == nil {
                    leftSum += node.Left.Val
                } else {
                    // add left node to nodeQ
                    newNode := &QNode{
                        TN: node.Left,
                    }
                    nodeQTail.Next = newNode
                    nodeQTail = nodeQTail.Next
                }
            }
            if node.Right != nil {
                if node.Right.Left != nil || node.Right.Right != nil {
                    // add right node to nodeQ
                    newNode := &QNode{
                        TN: node.Right,
                    }
                    nodeQTail.Next = newNode
                    nodeQTail = nodeQTail.Next
                }
            }
        }
        nodeQ = nodeQ.Next
    }

    return leftSum
    
}
