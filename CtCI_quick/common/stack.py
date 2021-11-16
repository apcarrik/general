class snode:
    def __init__(self):
        self.val = None
        self.next = None

    def SetVal(self,val):
        self.val = val

    def SetNext(self,next):
        self.next = next

    def GetVal(self):
        return self.val

    def GetNext(self):
        return self.next

class stack:
    def __init__(self):
        self.head = None

    def push(self,val):
        new = snode()
        new.SetVal(val)
        if(self.head == None):
            self.head = new
        else:
            new.SetNext(self.head)
            self.head = new

    def pop(self):
        if(self.head == None):
            raise RuntimeError("Stack is empty.")
        tmp = self.head
        self.head = self.head.GetNext()
        return tmp.GetVal()

def test_stack():
    s = stack()
    for v in [1,2,3]:
        s.push(v)
    assert s.pop() == 3
    assert s.pop() == 2
    s.push(55)
    assert s.pop() == 55
    assert s.pop() == 1
    exception = None
    try:
        s.pop()
    except RuntimeError as e:
        exception = e
    assert exception

test_stack()
