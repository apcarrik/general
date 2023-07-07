/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* Notes
    - Inorder traversal prefers going left as far as possible, then to the parent, then right as far as possible, then backwards
    - Postorder traversal goes left as far as possible, then right node and left as far as possible, then right as far as possible, then parent node, and then backwards
    - With postorder, the root is always the last node in the list
    - With preorder, the last node in the list is the rightmost
    - With both orderings, the first node in the list is the leftmost
    - With inorder, all items to left of root are left descendants of root, and all items to right of root are right descendants
        - you know which is root by the last element in postorder
    - With postorder, all items that are left descendants of root are grouped together in array, and all right descendants are grouped to the right of the left group, and then the root is the last element of array.
    Idea: Divide and conquor. Find the root by the last element of postorder. Then, find the root in the inorder array, which should be towards the middle of the array. All elements to the left of the root in the inorder array are left descendants of root, and same with the right being right descendants. Call build tree recursively on the left descendants and right descendants, passing in the correct indorder and postorder array slices and return the results.

*/
func findIndexOf(x int, arr []int) int {
    idx := len(arr)/2
    for i:=0; i<=len(arr); i++{    
        if idx >= len(arr) || idx < 0 {
            break
        } 
        if arr[idx] == x {
            return idx
        }
        if i % 2 == 0 {
            idx += i
        } else {
            idx -= i
        }
    }
    return -1
    
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    // edge case
    if len(inorder) == 0 {
        return nil
    }

    rootVal := postorder[len(postorder)-1]

    // find rightmost element to root in inorder array (it should be in middle, so do bloom search from middle element)
    inorderRootIdx := findIndexOf(rootVal, inorder)
    //fmt.Println("--- inorder:",inorder,"postorder:",postorder, "rootVal:", rootVal, "inorderRootIdx", inorderRootIdx)

    // find left group
    var leftTree *TreeNode
    if inorderRootIdx > 0 {
        leftTree = buildTree(inorder[:inorderRootIdx], postorder[:inorderRootIdx])
    }

    // find right group
    var rightTree *TreeNode
    if inorderRootIdx < len(inorder) - 1 {
        rightTree = buildTree(inorder[inorderRootIdx + 1:], postorder[inorderRootIdx:len(postorder)-1])
    }
    //fmt.Println("leftTree:",leftTree, "rightTree:",rightTree)
    return &TreeNode{
        Val: rootVal,
        Left: leftTree,
        Right: rightTree,
    }
}
