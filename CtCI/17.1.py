'''
Solution for problem 17.1 in Cracking the Coding Interview

addWithoutPlus takes two numbers and returns the sum of the two numbers, without using any arithmetic operators
'''

def addWithoutPlus(n1, n2):
    # TODO: check if n1 and n2 are 32 bit positive integers
    bmasks = [0x0001, 0x0002, 0x0004, 0x0008,
              0x0010, 0x0020, 0x0040, 0x0080,
              0x0100, 0x0200, 0x0400, 0x0800,
              0x1000, 0x2000, 0x4000, 0x8000]
    out = 0
    carry = 0x0000
    for i, bmask in enumerate(bmasks):
        n1b = n1 & bmask
        n2b = n2 & bmask
        out = out | ( ( n1b ^ n2b ) ^ carry )
        if(i > 14):
            if (n1b & n2b) | ((n1b ^ n2b) & carry):
                # TODO: raise overflow error
                print("Overflow")
                return -1
        else:
            carry = bmasks[i+1] if (n1b & n2b) | ((n1b ^ n2b) & carry) else 0x0000
    return out



# Test addWithoutPlus
assert addWithoutPlus(0,1) == 1
assert addWithoutPlus(3,5) == 8
assert addWithoutPlus(3456,798) == 3456+798
assert addWithoutPlus(0x8000,0x0fff) == 0x8fff
assert addWithoutPlus(0x8000,0x8000) == -1
assert addWithoutPlus(65432156,8974645) == -1