class obj:
    def __init__(self, val):
        self.val = val if val else None
        self.next = None
        self.min = None

class fancystack:

    def __init__(self):
        self.head = None
    def min(self):
        if self.head == None:
            raise ValueError()
        return self.head.min
    def push(self,val):
        newobj = obj(val)
        if self.head == None:
            newobj.min = val
        else:
            if(val < self.head.min):
                newobj.min =val
            else:
                newobj.min = self.head.min
        newobj.next = self.head
        self.head = newobj
    def pop(self):
        if not self.head:
            raise ValueError()
        else:
            retval = self.head
            self.head = self.head.next
            return retval.val

def test_fancystack():
    f = fancystack()
    for val in [1.0,2.0,3.3]:
        f.push(val)
    assert f.min() == 1.0
    f.push(0.4)
    assert f.min() == 0.4
    assert f.pop() == 0.4
    assert f.min() == 1.0
    assert f.pop() == 3.3
    assert f.min() == 1.0

test_fancystack()
