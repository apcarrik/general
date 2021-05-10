'''
Solution for problem 4.4 in Cracking the Coding Interview
checkBalanced takes an input binary tree and checks if it is balanced.
'''
#Using Adjacency list representation of binary tree



def checkBalanced(bt):
    # Idea: perform BFS starting at root node, and keep track of longest and shortest path
    #TODO: implement
    return False # Or True

def makebt(balanced):
    # creates a binary tree, either balanced or not
    if(balanced):
        return [[2,3],[1,4,5],[1,6,7],[2,8],[2],[3,9],[3],[4],[6]]
    else:
        return [[2,3],[1,4,5],[1],[2,6],[2],[4]]


assert checkBalanced(makebt(True)) == True
assert checkBalanced(makebt(False)) == False
