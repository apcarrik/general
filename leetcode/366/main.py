# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

"""
Notes:
- we want to return nodes per wave of them being leaves after all other leaves have been removed
- we can use dfs or bfs to find all leaves
- note that a node may lose its children in different waves
- if we do dfs for every level, it will be O(n*log(n)) becasue we will have repeated work when we traverse the tree for each level, and there are log(n) levels

Complexity: n nodes
    runtime: O(n log(n)) - we preform dfs, which takes O(n) time, for all layers of the tree, of which there are log(n), so overall runtime is O(n*log(n))
    memory: O(n) - we have to return a list including the values for all n nodes, so its O(n). DFS will have only log(n) items on stack at any time.
"""
class Solution:
    def findLeaves(self, root: Optional[TreeNode]) -> List[List[int]]:
        # DFS approach
        out: list[list[int]] = []
        self.layer: list[int] = []

        def dfs(node: TreeNode) -> bool:
            if node.left is None and node.right is None:
                self.layer.append(node.val)
                return True
            else:
                if node.left is not None:
                    if dfs(node.left):
                        node.left = None
                if node.right is not None:
                    if dfs(node.right):
                        node.right = None
            return False
        
        while not dfs(root):
            out.append(self.layer)
            self.layer = []

        out.append([root.val])
        return out
