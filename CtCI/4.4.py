'''
Solution for problem 4.4 in Cracking the Coding Interview
checkBalanced takes an input binary tree and checks if it is balanced.
'''
#Using Adjacency list representation of binary tree
class qnode:
    def __init__(self , data):
        self.data = data
        self.next = None

class queue:
    def __init__(self):
        self.head = None
        self.tail = None
    def enqueue(self, data):
        newnode = qnode(data)
        if(self.head == None and self.tail == None): #Queue empty
            self.head = newnode
            self.tail = newnode
        else:
            oldtail = self.tail
            oldtail.nexty = qnode(data)
            self.tail = newnode
    def dequeue(self):
        if(self.head == None and self.tail == None): #Queue empty
            return None #TODO change this to throw error
        else:
            oldhead = self.head
            self.head = oldhead.next
            if(self.head == None):
                self.tail = None
            return oldhead.data
    def isempty(self):
        return True if(self.head == None and self.tail == None) else False


def checkBalanced(bt, root):
    # Idea: perform BFS starting at root node, and keep track of longest and shortest path (high/low water marks)
    #TODO: implement
    return False # Or True

def makebt(balanced):
    # creates a binary tree, either balanced or not
    if(balanced):
        return [[2,3],[1,4,5],[1,6,7],[2,8],[2],[3,9],[3],[4],[6]]
    else:
        return [[2,3],[1,4,5],[1],[2,6],[2],[4]]


assert checkBalanced(makebt(True),1) == True
assert checkBalanced(makebt(False),1) == False
