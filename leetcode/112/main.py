# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if root is not None:
            if root.left is None and root.right is None:
                if root.val == targetSum:
                    return True
            else:
                if self.hasPathSum(root = root.left, targetSum = targetSum - root.val):
                    return True
                if self.hasPathSum(root = root.right, targetSum = targetSum - root.val):
                    return True
            return False
