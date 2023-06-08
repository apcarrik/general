# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        ret: list[list[int]] = []
        if root is not None:
            if root.left is None and root.right is None:
                # leaf - send back list with just leaf, previous callers will populate their node on the return value
                if root.val == targetSum:
                    ret.append([root.val])
            else:
                # nonleaf
                left_list: list[list[int]] = self.pathSum(root=root.left, targetSum=targetSum-root.val)
                if len(left_list) > 0:
                    for list in left_list:
                        list.insert(0,root.val)

                right_list: list[list[int]] = self.pathSum(root=root.right, targetSum=targetSum-root.val)
                if len(right_list) > 0:
                    for list in right_list:
                        list.insert(0,root.val)

                # merge left and right lists
                ret = left_list + right_list
        
        return ret
