'''
Solution for problem 5.2 in Cracking the Coding

binaryToString takes a float value that is between 0 and 1, and returns the 32-bit binary fraction of the value
if one exists. If it cannot be represented as a 32-bit binary fraction, it prints "ERROR" and returns NaN.
'''

def binaryToString(re):
    # test if 32 bit precision is not enough for input
    if( re % (2**(-32)) != 0.0):
        print("ERROR")
        return float("NaN")
    # get each binary digit of input
    retstr = ""
    for d in range(1,33):
        if((re / (2**(-d))) >= 1.0):
            retstr += "1"
            re -= 2**(-d)
        else:
            retstr += "0"
    return retstr

# Test function
testinput = 2**(-15) + 2**(-23) + 2**(-26) + 2**(-31) + 2**(-32)
expectedoutput = "00000000000000100000001001000011"
assert binaryToString(testinput) == expectedoutput