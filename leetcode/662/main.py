# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# class Solution:
#     def widthOfBinaryTree(self, root: Optional[TreeNode]) -> int:

class Solution:
    class IndexedNode:
        def __init__(self, node, index):
            self.node = node
            self.index = index

    def widthOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        tree_max_width = 0
        next_layer = [self.IndexedNode(node=root, index=0)]
        while len(next_layer) > 0 : 
            next_next_layer = []
            layer_min_idx = float('inf')
            layer_max_idx = 0
            for idx_node in next_layer:
                layer_min_idx = min(layer_min_idx, idx_node.index)
                layer_max_idx = max(layer_max_idx, idx_node.index)
                if idx_node.node.left is not None:
                    next_next_layer.append(self.IndexedNode(
                        node= idx_node.node.left,
                        index= (idx_node.index*2) )
                    )
                if idx_node.node.right is not None:
                    next_next_layer.append(self.IndexedNode(
                        node= idx_node.node.right,
                        index= (idx_node.index*2)+1 )
                    )
            tree_max_width = max(layer_max_idx - layer_min_idx+1, tree_max_width)
            next_layer = next_next_layer
        return tree_max_width
