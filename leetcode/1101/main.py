class Solution:
    def earliestAcq(self, logs: List[List[int]], n: int) -> int:
        nlookup: dict[int, int] = {i:i for i in range(n)} # nodeid:componentid
        clookup: dict[int,set[int]] = {i:set([i]) for i in range(n)} #componentid:set(nodeids)
        logs.sort()

        for ts, xi, yi in logs:
            if nlookup[xi] != nlookup[yi]:
                new: int = nlookup[xi]
                old: int = nlookup[yi]
                if len(clookup[nlookup[xi]]) < len(clookup[nlookup[yi]]): # swapping to have old be smaller component
                    new, old = old, new
                for nodeid in clookup[old]:
                    nlookup[nodeid] = new
                clookup[new] |= clookup[old]
                del[clookup[old]]
                if len(clookup) == 1:
                    return ts
        return -1
