/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    // Idea: BFS
    // uses priority queues to keep track of unexplored nodes and their depths
    var unexplored_nodes []*TreeNode
    var unexplored_node_depths []int

    unexplored_nodes = append(unexplored_nodes, root)
    unexplored_node_depths = append(unexplored_node_depths, 1)
    greatest_depth := 0

    for len(unexplored_nodes) > 0 {
        node := unexplored_nodes[0]
        unexplored_nodes = unexplored_nodes[1:]
        node_depth := unexplored_node_depths[0]
        unexplored_node_depths = unexplored_node_depths[1:]

        if node_depth > greatest_depth {
            greatest_depth = node_depth
        }

        if node.Left != nil {
            unexplored_nodes = append(unexplored_nodes,node.Left)
            unexplored_node_depths = append(unexplored_node_depths,node_depth+1)
        }

        if node.Right != nil {
            unexplored_nodes = append(unexplored_nodes,node.Right)
            unexplored_node_depths = append(unexplored_node_depths,node_depth+1)
        }
    }

    return greatest_depth
    
}
