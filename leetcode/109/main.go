/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* Idea:
    - store list as array, then recursivley construct sub-trees and merge them together

*/

func buildBST(listPtr *[]int, low, high int) *TreeNode {
    list := *listPtr
    //fmt.Println("high:",high,"low",low)
    if high-low < 2 {
        switch high-low {
            case 1:
                return &TreeNode{
                    Val: list[high],
                    Left: &TreeNode{
                        Val: list[low],
                    },
                }
            case 0:
                return &TreeNode{
                    Val: list[low],
                }
        }
    }
    mid := (high-low)/2 + low
    //fmt.Println("mid",mid)
    //fmt.Println("calling buildBST with left:[",low,mid-1,"],right:[",mid+1,high,"]")
    leftBST := buildBST(listPtr, low, mid-1)
    rightBST := buildBST(listPtr, mid+1, high)
    return &TreeNode{
        Val: list[mid],
        Left: leftBST,
        Right: rightBST,
    }
}

func sortedListToBST(head *ListNode) *TreeNode {
    // store list as slice
    list := []int{}
    for node := head; node != nil; node = node.Next {
        list = append(list, node.Val)
    }
    if len(list) == 0 {
        return nil
    }

    // recursivley build bst from bottom up
    return buildBST(&list, 0, len(list)-1)
}
