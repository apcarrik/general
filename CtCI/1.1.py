'''
Solution for Problem 1.1 in Cracking the Coding Interview.
brute_force solves the problem with n^2 runtime, where n is the length of the input string

'''

# Class CharHash 
class CharHash:
  def __init__(self, strlen):
      self.length = strlen
      self.chartable = [[]] * self.length # TODO: Turn this into a real linked list

  def hlookup(self,ch):
      return hash(ch)%self.length

  def add(self, ch):
      self.chartable[self.hlookup(ch)].append(ch)

  def contains(self, ch):
      for chmatch in self.chartable[self.hlookup(ch)] :
        if (ch == chmatch): return True
      return False

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
