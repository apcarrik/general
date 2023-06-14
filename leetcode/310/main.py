class Solution:
    def findMinHeightTrees(self, n: int, edges: list[list[int]]) -> list[int]:
        if n == 1:
            return [0]
        edges_map: dict[int, set[int]] = {}
        for src,dst in edges:
            # update edges map
            if src in edges_map:
                edges_map[src].add(dst)
            else:
                edges_map[src] = set([dst])
            if dst in edges_map:
                edges_map[dst].add(src)
            else:
                edges_map[dst] = set([src])

        leaves: list[int] = [i for i in range(n) if len(edges_map[i]) == 1]

        while n > 2:
            n -= len(leaves)
            next_leaves: list[int] = []
            for src in leaves:
                dst: int = edges_map[src].pop()
                edges_map[dst].remove(src)
                if len(edges_map[dst]) == 1:
                    next_leaves.append(dst)
            leaves = next_leaves
        return leaves
