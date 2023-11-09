# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# class Solution:
#     def longestZigZag(self, root: Optional[TreeNode]) -> int:


class Solution:

    def longestZigZagRecursive(self, node: Optional[TreeNode], previous_left: bool, current_zigzag_len: int, longest_zigzag_len: int) -> int:
        longest_zigzag_len = max(current_zigzag_len, longest_zigzag_len)
        if node.left is not None:
            if previous_left:
                longest_zigzag_len = max( 
                    self.longestZigZagRecursive(
                        node = node.left, 
                        previous_left = True, 
                        current_zigzag_len = 1, 
                        longest_zigzag_len = longest_zigzag_len) ,
                    longest_zigzag_len)
            else:
                longest_zigzag_len = max( 
                    self.longestZigZagRecursive(
                        node = node.left, 
                        previous_left = True, 
                        current_zigzag_len = current_zigzag_len+1, 
                        longest_zigzag_len = longest_zigzag_len) ,
                    longest_zigzag_len)    
        if node.right is not None:
            if previous_left:
                longest_zigzag_len = max( 
                    self.longestZigZagRecursive(
                        node = node.right, 
                        previous_left = False, 
                        current_zigzag_len = current_zigzag_len+1, 
                        longest_zigzag_len = longest_zigzag_len) ,
                    longest_zigzag_len)
            else:
                longest_zigzag_len = max( 
                    self.longestZigZagRecursive(
                        node = node.right, 
                        previous_left = False, 
                        current_zigzag_len = 1, 
                        longest_zigzag_len = longest_zigzag_len) ,
                    longest_zigzag_len) 
        return longest_zigzag_len

    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        return max(
            self.longestZigZagRecursive(
                node = root.left, 
                previous_left = True, 
                current_zigzag_len = 1, 
                longest_zigzag_len = 1) if root.left is not None else 0,
            self.longestZigZagRecursive(
                node = root.right, 
                previous_left = False, 
                current_zigzag_len = 1, 
                longest_zigzag_len = 1) if root.right is not None else 0 )
