'''
Solution for problem 2.2 in Cracking the Coding Interview.
kth_to_last returns the kth to last node of a linked list.
'''

class sllnode:
    def __init__(self, data):
        self.data = data
        self.next = None


def kth_to_last(h, k):
    # iterate thourgh linked list to find length
    iter = h
    length = 1
    while (iter.next != None):
        iter=iter.next
        length += 1

    # know length of lineked list, now can find kth element from end
    tgtnode = length - k
    iter = h
    for h in range(0,tgtnode):
        iter = iter.next
    return iter

def kth_to_last_efficient(h, k):
    #Idea: have two pointers k nodes apart, and the closer pointer will be on n-kth element when further pointer is at last element
    iter1 = h
    iter2 = h
    c = 1
    while(c < k):
        iter1 = iter1.next
        c += 1
    while(iter1.next != None):
        iter1 = iter1.next
        iter2 = iter2.next
    return iter2

def populatell(nodes):
    c = 1
    h = sllnode(c)
    iter = h
    for _ in range(2,nodes):
        c += 1
        iter.next = sllnode(c)
        iter = iter.next
    return h


#Test kth_to_last
h  = populatell(111)
assert kth_to_last(h, 39).data == 72
assert kth_to_last_efficient(h, 39).data == 72