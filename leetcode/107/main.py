# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []

        ret: list[list[int]] = []
        q: deque[tuple[int,TreeNode]] = deque([(0,root)])
        while len(q) > 0:
            (layer, node) = q.popleft()
            # add node to ret
            if len(ret) - 1 < layer:
                ret.append([node.val])
            else:
                ret[layer].append(node.val)
            # add node's children to q
            if node.left is not None:
                q.append((layer+1, node.left))
            if node.right is not None:
                q.append((layer+1, node.right))
                
        ret.reverse()
        return ret
