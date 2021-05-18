'''
Solution for problem 8.3 in Cracking the Coding Interview
'''

# magicIndex
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
                else:
                    midpoint = leftidx + (difference // 2)
                    if(midpoint <= A[midpoint]):
                        rightidx = midpoint
                    else:
                        leftidx = midpoint
    else: # fn is strictly decreasing
        if(A[0]>=0 and A[n-1] <= n-1):
            leftidx = 0
            rightidx = n-1
            while rightidx-leftidx >=1:
                if(leftidx == A[leftidx] or rightidx == A[rightidx]):
                    return True
                difference = rightidx - leftidx
                if(difference == 1):
                    break
                else:
                    midpoint = leftidx + (difference // 2)
                    if(midpoint <= A[midpoint]):
                        leftidx = midpoint
                    else:
                        rightidx = midpoint

    return False

# test magicIndex
A = [-3,-2,-1,2,4,6,7,9]
assert magicIndex(A) == True

A = [-3,-2,-1]
assert magicIndex(A) == False

A = [9,7,2,0,-1,-2,-3,-4]
assert magicIndex(A) == True

A = [-1,-2,-3]
assert magicIndex(A) == False