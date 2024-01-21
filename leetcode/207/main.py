class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # create adjacency list representation of graph and map of indegrees
        adj: dict[int, list[int]] = {}
        indeg: dict[int,int] = {}
        for dst, src in prerequisites:    
            # add to indegrees
            if dst in indeg:
                indeg[dst] += 1
            else:
                indeg[dst] = 1
            
            # add to adjacency map
            if src in adj:
                adj[src].append(dst)
            else:
                adj[src] = [dst]

        # create queue of no indegree nodes
        q: deque[int]= deque()
        for n in range(numCourses):
            if n not in indeg:
                q.append(n)
        
        # iterate through queue until empty
        while len(q) > 0:
            n = q.popleft()
            if n in adj:
              for src in adj[n]:
                  indeg[src] -= 1
                  if indeg[src] == 0:
                      q.append(src)
                      del(indeg[src])

        # once queue is exhausted, check if anything remains in indegrees
        if len(indeg) == 0:
            return True
        return False
