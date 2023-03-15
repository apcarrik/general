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
    - Idea: traverse like you were making the array representation of the binary tree. Keep a queue of nodes to visit and add children of nodes from left to right so it is in order. Also keep an array of each node's layer height. If you come across a null node, add it to array with value equal to the negative of the layer height. Ensure that all negative values are within 1 to ensure the heights of all null children are within 1.


*/

type layerNode struct {
    layer int
    node *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
    layers := []int{}    
    for q := []*layerNode{&layerNode{layer: 0, node: root}}; len(q) > 0; q = q[1:]{
        layer := q[0].layer + 1
        node := q[0].node
        if node.Left == nil {
            layers = append(layers, 0-layer)        
        } else {
            // add to q and layers
            layers = append(layers, layer)
            q = append(q, &layerNode{layer: layer, node: node.Left})
        }
        if node.Right == nil {
            layers = append(layers, 0-layer)       
        } else  {
            // add to q and layers
            layers = append(layers, layer)
            q = append(q, &layerNode{layer: layer, node: node.Right})
        }
    }

    // Go through layers, making sure all negative layers are within 1 of each other
    minLayer := 0
    maxLayer := -100
    negativeSeen := false
    for _,layer := range layers {
        if layer >= 0 && negativeSeen == true {
            return false
        }
        if layer < 0 {
            negativeSeen = true
            if layer < minLayer {
                minLayer = layer
            }
            if layer > maxLayer {
                maxLayer = layer
            }
        }
    }

    return maxLayer - minLayer <= 1
    
}
