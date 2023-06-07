class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort()
        current_start: int = intervals[0][0]
        current_end: int = intervals[0][1]
        non_overlapping: set(tuple(int,int)) = set()
        for interval in intervals:
            if interval[0] >= current_end:
                non_overlapping.add(tuple([current_start,current_end]))
                current_start = interval[0]
                current_end = interval[1]      
            elif interval[1] < current_end:
                current_start = interval[0]
                current_end = interval[1]
        non_overlapping.add(tuple([current_start,current_end]))
        return len(intervals) - len(non_overlapping)
