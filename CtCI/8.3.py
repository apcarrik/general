'''
Solution for problem 8.3 in Cracking the Coding Interview
'''

# magicIndex TODO: finish filling out function description
def magicIndex(A):
    # test if fn is strictly increasing or strictly decreasing
    n = len(A)
    if (A[0] < A[n - 1]): # fn is strictly increasing
        # TODO: check if fn could intersect A[i] = i line

    else: # fn is strictly decreasing
        #TODO: finish this



# TODO: implement test for magicIndex
A = [-3,-2,-1,4,5,7,9]
assert magicIndex(A) == True