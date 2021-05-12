'''
Solution for problem 5.1 in Cracking the Coding Interview
checkBalanced takes an input binary tree and checks if it is balanced.
'''

def bitInsertion(N,M,i,j):
    if (i >j):
        tmp = j
        j = i
        i = tmp

    # Create M'
    Mmask = (1 << (j - i + 1)) - 1
    Mprime = (Mmask & M) << i

    # Create N'
    tmp1 = Mmask << i
    tmp2 = ~tmp1
    Nprime = N & tmp2

    # Out
    return Nprime + Mprime

#TODO test this