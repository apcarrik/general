'''
Solution for problem 10.1 in Cracking the Coding Interview

sortedMerge takes input sorted arrays A and B and returns a combined sorted array of the elements of both A and B.
Note that array A has enough space to hold all of the elements of B, and should be the returned array.
'''

# sortedMerge returns A with elements of B added in order
def sortedMerge(A, B):
    # idea: start with the largest elements, filling in the indexes at the end of A. Basically like one iteration of
    # mergesort.
    idx = len(A)-1
    aiter = len(A)-1-len(B)
    biter = len(B)-1
    while biter >= 0:
        if(A[aiter] < B[biter]):
            A[idx] = B[biter]
            biter -= 1
        else:
            A[idx] = A[aiter]
            aiter -= 1
        idx -= 1
    return A

# Test sortedMerge
testinputA = [-3, 0, 1, 55, 56, None, None, None]
testinputB = [-2, 23, 77]
expectedoutput = [-3, -2, 0, 1, 23, 55, 56, 77]
assert sortedMerge(testinputA, testinputB) == expectedoutput