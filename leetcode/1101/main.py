class Solution:
    def earliestAcq(self, logs: List[List[int]], n: int) -> int:
        idlookup: dict[int, int] = {i:i for i in range(n)} # nodeid:componentid
        componentlookup: dict[int,set[int]] = {i:set([i]) for i in range(n)} #componentid:set(nodeids)
        logs.sort() # assume logs is in chronological order, otherwise call logs.sort()

        for l in logs:
            ts, xi, yi = l
            if idlookup[xi] != idlookup[yi]:
                newcomponent: int = idlookup[xi]
                oldcomponent: int = idlookup[yi]
                if len(componentlookup[idlookup[xi]]) < len(componentlookup[idlookup[yi]]): # swapping to have oldcomponent be smaller component
                    newcomponent, oldcomponent = oldcomponent, newcomponent
                for nodeid in componentlookup[oldcomponent]:
                    idlookup[nodeid] = newcomponent
                componentlookup[newcomponent] |= componentlookup[oldcomponent]
                del[componentlookup[oldcomponent]]
                if len(componentlookup) == 1:
                    return ts
        return -1
