'''
Solution for problem 5.1 in Cracking the Coding Interview
bitInsertion takes two 32-bit numbers M and N, and two integers < 32, i and j, and returns a value equal to
N with M insterted into the bits from j to i.
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

inputN = int('10000000000',2)
inputM = int('10011',2)
inputi = 2
inputj = 6
expectedoutput = int('10001001100',2)
assert bitInsertion(inputN, inputM, inputi, inputj) == expectedoutput