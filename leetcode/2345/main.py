"""
Start time: 10:10-10:40am, 1:21pm

Notes: 
- The tallest mountain wont be blocked unless another mountain of same height is at same location.
- We probably want to sort by height and iterate from tallest to shortest
- we need a way to determine if a mountains peak lies within another mountains area
  - we could first approximate mountain with box, and if peak is within box, check if it is within mountain
  - or, we could split mountain in half to make it easier
  - maybe we just store start, end of line, knowing were looking for point below or on line
  - if there are m mountains, then there are 2m lines to check, and were checking for all m mountains so time complexity is O(m^2) and space would be O(m)
- Splitting mountain in half made it harder, actually
  - also, we kind of need to have a more efficient way of searching for mountains the peak may be in
    - we could keep the mountains sorted by x value and use binary search
    - if we know this mountain is leq in height to all other mountains in bounds, then we know that if this mountains bounds are within or equal to larger mountains bounds, the peak is not visible.
        - so, if this mountain has bounds b=[l,r], for all other mountains, check for mountain m's bounds m[ml, mr] that if ml <= l and mr >= r, then b is not visible. If ml=l and mr=r, neither mountain is visible.
    Idea: order peaks by y value and iterate from largest to smallest. For each peak, check list bounds if it is within the bounds (x range) of another mountain. If so, then check if the point lies within the mountain. If both of these are not true, then increment visible counter. If they are both true, then check if this peak matches the previous mountain's peak, and if so decrement visible counter. Then, add the bounds of this peak to the bounds list in the correct order
    Note that you'll need to implement a leq version of binary search to search for this mountains peak within bounds.
    Binsearch will be custom logic where you search for first bounds where first element is <= 

example: peaks = [[2,2],[6,3],[5,4]], bounds = [], visible = 0
    p=[2,2], b=[0,4]. bounds is empty so visible += 1 = 1. add b to bounds=[[0,4]]
    p=[6,3], b=[3,9]. binsearch bounds for 3, get idx 0 => bounds[0] = [0,4]. p[0]=6>bounds[0][1]=4, so check next in bounds. no more in bounds, so visible += 1 and add p to bounds => bounds = [[0,4],[3,9]]
    p=[5,4], binsearch bounds for 5, get idx=0 => bounds[0]=[0,4]. p[0]=5

I'm just going to start solution without binary search, and maybe add it later if it takes too long without it.

Problem: if we have more than 2 of the same peak, we decrement visible more than necessary
    Solution: keep frequency map of peaks seen, and if count is = 1, decrement visible, but if count > 1: don't decrement visible
    Ok, decrementing visible is not working, because you have no way of knowing if previous duplicate peak incremented visible or not
        Solution: when about to increment visible, check that next peak is not the same (it would have to be next if it exists in peaks because it was ordered)

Another problem with keeping an ordered list of bounds is that inserting in the correct spot is O(n), and since we're doing this for all n peaks it causes overall runtime to be O(n^2)
    could we just search the entire list of mountains?
    means we'd need to pre-compute the list of mountain bounds

if a mountain has height y, bounds l and r, then we'd need to search for all other mountains m with my, ml, and mr such that ml <= r and mr >= l, and my >= y


ACTUALLY, I think we need to start with the smallest mountains first. That way we can bound the search area to within the current mountain for any leq size mountains within this mountains boundaries.

Idea: order peaks by y value ascending. Now, make an list len(peaks) long called visible that's intialized to 1 for all elements. The index of this will correspond with index of peak. Iterate through each mountain m with peaks index mi and peak at (mx,my) and bounds (ml,mr). Do a leq binary search in peaks for ml. For all mountains starting at that index of peaks until we hit the first peak past mr:
    other mountain o with peaks index oi has peak at (ox,oy) and bounds (ol, or). if (ox,oy) == (mx,my), ensure visible[mi] and visible[oi] are = 0. elif ol>ml and or<mr, ensure visible[oi] = 0.
    Return the sum of visible

Complexity:
    Runtime: still O(n^2) in the worst case, if every mountain is subsumed by a larger mountain (Russian nesting dolls setup). But, for the general case it should be better.
    Memory: still O(n)

Yeah, this still gets Time Limit Exceeded

Need to look at published solutions.

Ok, following this solution: https://leetcode.com/problems/finding-the-number-of-visible-mountains/solutions/2311603/python-3-monotonic-stack-sort-explanation/?envType=study-plan-v2&envId=google-spring-23-high-frequency

It looks like the suggested method is to sort peaks by x value and then iterate through each one, checking if it has any previously found points within it or it is within the most recent point. we keep track of previous points with a monotonic increasing stack. For each peak, pop items off the stack if they are within this peak, until one is found not in it or stack is empty. then, check if this peak is within the top of the stack, and if not, append to stack.

Also, use a frequency map to discard any that have frequency greater than 1 (but keep them in consideration for occluding other mountains, so only filter them at the end).
"""

from collections import Counter
class Solution:
    def visibleMountains(self, peaks: List[List[int]]) -> int:
        fmap: dict[tuple[int,int],int] = Counter([(l[0],l[1]) for l in peaks])
        peaks = sorted(fmap.keys())
        if len(peaks) == 0:
            return 0
        def within(t1: tuple[int,int], t2: tuple[int,int])->bool:
            x1,y1 = t1
            x2,y2 = t2
            b1,b2 = y1-x1, x1+y1
            # l2,r2 = x2-y2, x2+y2
            return y2 <= x2 + b1 and y2 <= -x2 + b2

        ms: list[tuple[int,int]] = [peaks[0]]

        for pi in range(1,len(peaks)):            
            while len(ms) > 0 and within(peaks[pi],ms[-1]):
                ms.pop() 
            if len(ms) == 0 or not within(ms[-1],peaks[pi]):
                ms.append(peaks[pi])
        return len([t for t in ms if fmap[t] == 1])



from bisect import bisect_left
class SlowSolution:
    def visibleMountains(self, peaks: List[List[int]]) -> int:

        #iterate through peaks sorted by y value, checking if any other mountains are within this one and setting their visible value accordingly
        peaksx: list[list[int]] = sorted(peaks.copy())
        peaks = sorted([p[::-1] for p in peaks])
        # print("peaksx=",peaksx)
        # print("peaks=",peaks)
        visible: list[int] = [1 for _ in range(len(peaks))]
        for i, coords in enumerate(peaks):
            y,x = coords
            b: tuple[int,int] = (x-y,x+y)

            # use bisect to find potential subsumed matches
            oi: int = bisect.bisect_left(peaksx, b[0], key=lambda t:t[0])
            if oi < 0: 
                oi = 0
            # print(coords, oi, b, peaksx[oi] if oi < len(peaksx) else "")
            firstdup: bool = True
            while oi < len(peaksx) and peaksx[oi][0]<=b[1]:
                ox: int = peaksx[oi][0]
                oy: int = peaksx[oi][1]
                ob: tuple[int,int] = (ox-oy,ox+oy)
                # print(b, ob)
                if ob[0] >= b[0] and ob[1] <= b[1]:
                    if ox == x and oy == y:
                        if firstdup:
                            firstdup = False
                            oi += 1 
                            continue
                        else:
                            # print("found duplicate")
                            visible[i] = 0 
                    # print("Found mountain in this one")
                    visible[oi] = 0
                oi += 1 
        return sum(visible)
    
