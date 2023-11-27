class Event:
    def __init__(self, start: int, end: int):
        """Returns an event object with start and end times set."""
        self.start: int = start
        self.end: int = end
    def __repr__(self):
        """Returns a string representation of an event object."""
        return f"start: {self.start}, end: {self.end}"

class MyCalendar:       
    def __init__(self):
        """Returns a MyCalendar object with its empty event list initialized"""
        self.events: List[Event] = []

    def __getPreviousEvent(self, start: int) -> int:
        """Returns the index of the event that comes before or at the same time as start. If none are found, returns -1."""
        if len(self.events) == 0:
            return -1

        low: int = 0
        high: int = len(self.events) - 1
        while high > low:
            mid: int = (high-low)//2 + low
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
        """Books an event on this MyCalendar object's calendar if possible. Returns a boolean value indicating wheither or not the booking was succesful."""
        previous_event_index: int = self.__getPreviousEvent(start)
        if previous_event_index == -1:
            if len(self.events) != 0 and end > self.events[0].start:
                return False
            else:
                self.events = [Event(start=start, end=end)] + self.events
                return True
        elif previous_event_index == len(self.events) - 1:
            if self.events[previous_event_index].end > start:
                return False
            else:
                self.events = self.events + [Event(start=start, end=end)]
                return True
        else:
            if self.events[previous_event_index].end > start or self.events[previous_event_index + 1].start < end:
                return False
            else:
                self.events = self.events[:previous_event_index+1] + [Event(start=start, end=end)] + self.events[previous_event_index+1:]
                return True

    def __repr__(self):
        """Returns a string representation of this MyCalendar object"""
        ret_str: str = "["
        for event in self.events:
            ret_str += f"({event.start},{event.end})"
        return ret_str + "]"

# Your MyCalendar object will be instantiated and called as such:
# obj = MyCalendar()
# param_1 = obj.book(start,end)
