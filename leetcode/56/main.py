class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) <= 1:
            return intervals

        intervals.sort()
        current_start: int = intervals[0][0]
        current_end: int = intervals[0][1]
        output: list[list[int,int]] = []
        for (interval_start, interval_end) in intervals:
            if interval_start <= current_end:
                if interval_end > current_end:
                    current_end = interval_end
            else:
                output.append([current_start, current_end])
                current_start = interval_start
                current_end = interval_end
        output.append([current_start, current_end])
        return output
