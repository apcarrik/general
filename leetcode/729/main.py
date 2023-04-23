# class MyCalendar:

#     def __init__(self):
#         # print(sys.version)

#     def book(self, start: int, end: int) -> bool:
#         return False

class Event:
    def __init__(self, start: int, end: int):
        self.start = start
        self.end = end
    def __repr__(self):
        return f"start: {self.start}, end: {self.end}"

class MyCalendar:   
    
    def __init__(self):
        self.events = [] # list of Event objects

    def __getPreviousEvent(self, start: int) -> int:
        # performs binary search for event that has start time closest to but not greater than start and returns the index of that event. If none are found, returns -1
        if len(self.events) == 0:
            return -1
        low = 0
        high = len(self.events) - 1
        while high > low:
            mid = (high-low)//2 + low
            if self.events[mid].start == start:
                return mid
            elif self.events[mid].start > start:
                high = mid - 1
            else:
                low = mid + 1

        if low == 0 and self.events[0].start > start:
            return -1
        if low > 0 and self.events[low].start > start and self.events[low-1].start < start:
            low -= 1
        return low   

    def book(self, start: int, end: int) -> bool:
        previous_event = self.__getPreviousEvent(start)
        # print(start,end,previous_event,self.events[previous_event]if previous_event >= 0 and previous_event < len(self.events) else -1, self)
        if previous_event == -1:
            if len(self.events) != 0 and end > self.events[0].start:
                return False
            else:
                self.events= [Event(start=start, end=end)] + self.events
                return True
        elif previous_event == len(self.events) - 1:
            if self.events[previous_event].end > start:
                return False
            else:
                self.events = self.events + [Event(start=start, end=end)]
                return True
        else:
            if self.events[previous_event].end > start or self.events[previous_event + 1].start < end:
                return False
            else:
                self.events = self.events[:previous_event+1] + [Event(start=start, end=end)] + self.events[previous_event+1:]
                return True
    def __repr__(self):
        ret_str = "["
        for event in self.events:
            ret_str += f"({event.start},{event.end})"
        return ret_str + "]"

# Your MyCalendar object will be instantiated and called as such:
# obj = MyCalendar()
# param_1 = obj.book(start,end)
