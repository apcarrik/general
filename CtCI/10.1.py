'''
Solution for problem 10.1 in Cracking the Coding Interview

sortedMerge takes input sorted arrays A and B and returns a combined sorted array of the elements of both A and B.
Note that array A has enough space to hold all of the elements of B, and should be the returned array.
'''

# sortedMerge returns A with elements of B added in order
def sortedMerge(A, B):
    # idea: start with the largest elements, filling in the indexes at the end of A. Basically like one iteration of
    # mergesort.

    return None