'''
Solution for problem 16.5 in Cracking the Coding Interview

factorialZeros takes a number n and returns the number of trailing zeros in n!.
'''
def getExponentOfFive(n):
    exponent = 0
    while(n % 5 == 0):
        exponent += 1
        n /= 5
    return exponent

def getNumZeros(n):
    num_zeros = 0
    iterations = n // 5 # assumes n will always be multiple of 5
    for i in range(1,iterations+1):
        num_zeros += (getExponentOfFive(i) + 1)
    return num_zeros

def factorialZeros(n):
    if(n<=5):
        if(n<5): return 0
        else: return 1
    prev_n_mod_5 = n - (n % 5)
    if (n % 5 == 0): #TODO: optimize this by just checking last digit == 0 or 5 (would probably need to check least signifigant 3 bits == 000 or 101)
        prev_n_mod_5 -= 5
        # lookup prev n mod 5
        prev_n_mod_5_num_zeros = getNumZeros(prev_n_mod_5)
        # calculate additional zeros
        new_num_zeros = prev_n_mod_5_num_zeros + getExponentOfFive(n)
    else:
        new_num_zeros = factorialZeros(prev_n_mod_5)
    return new_num_zeros

# Test factorialZeros
assert factorialZeros(4) == 0
assert factorialZeros(5) == 1
assert factorialZeros(9) == 1
assert factorialZeros(10) == 2
assert factorialZeros(14) == 2
assert factorialZeros(15) == 3
assert factorialZeros(19) == 3
assert factorialZeros(20) == 4
assert factorialZeros(24) == 4
assert factorialZeros(25) == 6
assert factorialZeros(30) == 7
assert factorialZeros(1000) == 249
assert factorialZeros(2500) == 624
assert factorialZeros(5000) == 1249
assert factorialZeros(10000) == 2499
assert factorialZeros(50000000) == 12499999
assert factorialZeros(55555555) == 13888884
