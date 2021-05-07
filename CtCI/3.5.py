'''
Solution for problem 3.5 in Cracking the Coding Interview
sortStack takes an input stack with methods push, pop, peek, and isEmpty, and returns the stack with the elements ordered
    smallest on top. This is done using only one other stack, and no other data structures.
'''

class stack:
    def __init__(self):
        self.top = None

    def push(self, item):
        oldtop = self.top
        self.top = stnode(item)
        self.top.next = oldtop

    def pop(self):
        if (self.top == None):
            print("pop error, self.top == None")
            return None # TODO: throw an error
        else:
            popped = self.top.data
            self.top = self.top.next
            return popped

    def peek(self):
        if (self.top == None):
            print("peek error, self.top == None")
            return None #TODO: throw an error
        else:
            return self.top.data

    def isEmpty(self):
        return (self.top == None)

    def iterateStack(self):
        eles = []
        iter = self.top
        while iter != None:
            eles.append(iter.data)
            iter = iter.next
        return eles

class stnode:
    def __init__(self, item):
        self.data = item
        self.next = None

# Implements stack sorting using only an additional stack and no other data structures
def sortStack(s1):
    s2 = stack()

    # find length of s1
    length =0
    while((not s1.isEmpty())):
        length +=1
        s2.push(s1.pop())
    while((not s2.isEmpty())):
        s1.push(s2.pop())
    
    #Sort, maintaining s1 with elements sorted largest on botoom and s2 sorted with smallest on bottom
    numsorted = 0
    min = None
    max = None
    while(numsorted < length-1 ):
        # going from s1 to s2
        max = s1.pop()
        for _ in range(1, length-numsorted):
            if(s1.peek() > max):
                s2.push(max)
                max = s1.pop()
            else:
                s2.push(s1.pop())
        s1.push(max)
        numsorted +=1
        # going from s2 to s1
        min = s2.pop()
        for _ in range(1, length-numsorted):
            if(s2.peek() < min):
                s1.push(min)
                min = s2.pop()
            else:
                s1.push(s2.pop())
        s2.push(min)
        numsorted +=1
    #pop all of s2 onto s1
    while(not s2.isEmpty()):
        s1.push(s2.pop())
    return s1

def compareStacks(s1, s2):
    #test data mismatch
    while((not s1.isEmpty()) and (not s2.isEmpty())):
        if(s1.pop() != s2.pop()):  #TODO: check that this works for all data types
            return False 
    #test length mismatch
    if (s1.isEmpty() and s2.isEmpty()):
        return True
    else:
        return False
            
def makeStackFromList(lst):
    st = stack()
    for el in reversed(lst):
        st.push(el)
    return st

# Test implementation
testinput = makeStackFromList([4,12,3,9,7])
expectedoutput = makeStackFromList([3,4,7,9,12])
assert compareStacks(sortStack(testinput), expectedoutput)