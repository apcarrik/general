'''
Solution for Problem 1.1 in Cracking the Coding Interview.
brute_force solves the problem with O(n^2) runtime, where n is the length of the input string
hash_method solves the problem with O(n) runtime using a hashmap with a linked list
'''
# Class LnkdListNode
class LnkdListNode:
    def __init__(self, ch):
        self.data = ch
        self.next = None

# Class LnkdList
class Lnkdlist:
    def __init__(self):
        self.head = None

    def add(self, ch):
        newnode = LnkdListNode(ch)
        if (self.head):
            nextnode = self.head
            while nextnode.next:
                nextnode = nextnode.next
            nextnode.next = newnode
        else:
            self.head = newnode
    
    def contains(self, ch):
        if (self.head):
            querynode = self.head
            while (querynode):
                if(querynode.data == ch): return True
                querynode = querynode.next
        return False

# Class CharHash 
class CharHash:
  def __init__(self, strlen):
      self.length = strlen
      self.chartable = [Lnkdlist] * self.length

  def hlookup(self,ch):
      return hash(ch)%self.length

  def add(self, ch):
      self.chartable[self.hlookup(ch)].add(ch)

  def contains(self, ch):
      self.chartable[self.hlookup(ch)].contains(ch)

# Function: brute_force 
def brute_force(str):
    uniques = []
    for ch in str:
        for un in uniques:
            if (ch == un): return False
        uniques.append(ch)
    return True


# Function: hash_method
def hash_method(str):
    chash = CharHash(len(str))
    for ch in str:
        if chash.contains(ch): return False
        chash.add(ch)
    return True

# Test Implemenataions

truestr = "abcdefghijklmnopqrstuvwxyz"
falsestr = "abcdefghijklmnopqrstuvwxyza"

assert brute_force(truestr) == True
assert brute_force(falsestr) == False
