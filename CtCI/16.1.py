'''
Solution for problem 16.1 in Cracking the Coding Interview

numberSwapper takes two numbers and returns the two numbers in swapped order, without using any temporary variables.

Note: this does not work for some floating point numbers due to rounding error! The numbers will be slightly different.
'''

def numberSwapper(n1, n2):
    #Idea: store the sum of the two numbers into n1, and then store the difference of n1-n2 into n2 and the n1
    print(n1,n2)
    n1 = n1 + n2
    n2 = n1 - n2
    n1 = n1 - n2
    print(n1,n2)
    return n1, n2

# test function
for n1, n2 in [
    (48, 89),
    (-23,27),
    (99.568,-2),
    (-56.33,-22.67),
    (32165432165432156432,564216584321654987651)
]:
    assert numberSwapper(n1,n2) == (n2,n1)

#Note: it works for most ints and floats, except floats with long decimal places (probably due to rounding errors)
assert numberSwapper(65458465.5646,-564532.2165) != (-564532.2165,65458465.5646)