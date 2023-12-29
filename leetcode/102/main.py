# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from collections import deque
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        ret: list[list[int]] = []
        q: Deque[tuple[int,TreeNode]] = deque([(0,root)])
        while len(q) > 0:
            (layer, node) = q.popleft()
            # add this to ret
            if len(ret)-1 < layer:
                ret.append([node.val])
            else:
                ret[layer].append(node.val)

            # add children to q
            if node.left is not None:
                q.append((layer+1, node.left))
            if node.right is not None:
                q.append((layer+1, node.right))
        return ret
