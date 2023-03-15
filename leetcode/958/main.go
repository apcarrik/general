/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/* Notes:
    - the array notation for a tree makes it easy to identify where it goes wrong, i.e. if a null exists in the array at all means that there is another element to the right of it.
    - Idea: traverse like you were making the array representation of the binary tree. Keep a queue of nodes to visit and add children of nodes from left to right so it is in order. Also keep an array of all nodes visitied, indicating if that node was null or not. In the end, ensure that no non-null nodes come after null nodes


*/

func isCompleteTree(root *TreeNode) bool {
    nullNodes := []bool{}    
    for q := []*TreeNode{root}; len(q) > 0; q = q[1:]{
        node := q[0]
        if node.Left == nil {
            nullNodes = append(nullNodes, true)        
        } else {
            // add to q and nullNodes
            nullNodes = append(nullNodes, false)
            q = append(q, node.Left)
        }
        if node.Right == nil {
            nullNodes = append(nullNodes, true)       
        } else  {
            // add to q and nullNodes
            nullNodes = append(nullNodes, false)
            q = append(q,node.Right)
        }
    }

    // Go through nullNodes, ensuring no non-null nodes come after null nodes
    nullSeen := false
    for _,nodeIsNull := range nullNodes {
        if !nodeIsNull && nullSeen == true {
            return false
        }
        if nodeIsNull {
            nullSeen = true
        }
    }

    return true
    
}
