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
            newnode = qnode(data)
            oldtail.next = newnode
            self.tail = newnode
    def dequeue(self):
        if(self.head == None and self.tail == None): #Queue empty
            print("dequeue error")
            return None #TODO change this to throw error
        else:
            oldhead = self.head
            self.head = oldhead.next
            if(self.head == None):
                self.tail = None
            return oldhead.data
    def isempty(self):
        return True if(self.head == None and self.tail == None) else False
    def printqueue(self): # For debugging
        print("Queue contents:")
        tmp = self.head
        while(tmp != None):
            print(tmp.data, " , ", tmp.next)
            tmp = tmp.next


def checkBalanced(bt, root):
    # Idea: perform BFS starting at root node, and keep track of longest and shortest path (high/low water marks)
    lowwater = -1
    highwater = 0
    nodeq = queue()
    depthq = queue()
    pnodeq = queue()
    nodeq.enqueue(root)
    depthq.enqueue(1)
    pnodeq.enqueue(None)
    watchdog = 0
    while(nodeq.isempty() == False and watchdog < 100):
        watchdog +=1
        currentnode = nodeq.dequeue()
        depth = depthq.dequeue()
        parent = pnodeq.dequeue()
        if(bt[currentnode] == [parent]):
            if(lowwater > depth or lowwater == -1):
                lowwater = depth
            if(highwater < depth):
                highwater = depth
        else:
            for child in bt[currentnode]:
                if(parent == None or child != parent):
                    nodeq.enqueue(child)
                    depthq.enqueue(depth+1)
                    pnodeq.enqueue(currentnode)           
    if(highwater-lowwater > 1):
        return False
    else:
        return True

def makebt(balanced):
    # creates a binary tree, either balanced or not
    if(balanced):
        return [[],[2,3],[1,4,5],[1,6,7],[2,8],[2],[3,9],[3],[4],[6]]
    else:
        return [[],[2,3],[1,4,5],[1],[2,6],[2],[4]]


assert checkBalanced(makebt(True),1) == True
assert checkBalanced(makebt(False),1) == False
