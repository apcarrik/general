# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# class Solution:
#     def minimumFlips(self, root: Optional[TreeNode], result: bool) -> int:

# Inspired by blue_sky5's solution: https://leetcode.com/problems/minimum-flips-in-binary-tree-to-get-result/solutions/2187968/python-pure-dfs/?envType=study-plan-v2&envId=google-spring-23-high-frequency
class Solution:
    def minimumFlips(self, root: TreeNode, result: bool) -> int:
        def dfs(root: TreeNode) -> tuple[int,int]:
            # returns tuple with number of flips needed to evaluate True and False, respectivley
            if root.val == 0: # False
                return (1,0)
            elif root.val == 1: # True
                return (0,1)

            elif root.val == 5: # NOT
                c: TreeNode = root.left if root.left is not None else root.right
                ct, cf = dfs(c)
                return (cf, ct)


            lt, lf = dfs(root.left)
            rt, rf = dfs(root.right)
            
            if root.val == 2: # OR
                changet: int = min(lt+rt, lt+rf, lf+rt)
                changef: int = lf+rf
                return (changet, changef)

            elif root.val == 3: # AND
                changet: int = lt+rt
                changef: int = min(lf,rf)
                return (changet, changef)

            elif root.val == 4: # XOR
                changet: int = min(lt+rf, lf+rt)
                changef: int = min(lt+rt, lf+rf)
                return (changet, changef)

        ct, cf = dfs(root)
        return ct if result is True else cf
