'''
Solution for problem 8.3 in Cracking the Coding Interview
'''

# magicIndex TODO: debug function
def magicIndex(A):
    # test if fn is strictly increasing or strictly decreasing
    n = len(A)
    if (A[0] < A[n - 1]): # fn is strictly increasing
        if(A[0] <= 0 and A[n-1] >= n-1):
            leftidx = 0
            rightidx = n-1
            while rightidx-leftidx >=1:
                if(leftidx == A[leftidx] or rightidx == A[rightidx]):
                    return True
                difference = rightidx - leftidx
                if (difference == 1):
                    break
                elif(difference %2 == 0): # even distance
                    midpoint = leftidx + (difference / 2)
                    if(midpoint <= A[midpoint]):
                        leftidx = midpoint
                    else:
                        rightidx = midpoint
                else: # odd distance
                    midpoint = leftidx + (difference / 2)
                    if(midpoint >= A[midpoint]):
                        rightidx = midpoint
                    else:
                        leftidx = midpoint+1
    else: # fn is strictly decreasing
        #TODO: debug this
        if(A[0]>=0 and A[n-1] <= n-1):
            leftidx = 0
            rightidx = n-1
            while rightidx-leftidx >=1:
                if(leftidx != A[leftidx] or rightidx != A[rightidx]):
                    return True
                difference = rightidx - leftidx
                if(difference == 1):
                    break
                elif(difference %2 == 0): # even distance
                    midpoint = leftidx + (difference / 2)
                    if(midpoint <= A[midpoint]):
                        leftidx = midpoint
                    else:
                        rightidx = midpoint
                else: # odd distance
                    midpoint = leftidx + (difference / 2)
                    if(midpoint >= A[midpoint]):
                        rightidx = midpoint
                    else:
                        leftidx = midpoint+1



# TODO: implement test for magicIndex
A = [-3,-2,-1,4,5,7,9]
assert magicIndex(A) == True