'''
Solution for problem 8.3 in Cracking the Coding Interview
permsNoDupes takes an input string of unique characters (alphabet) and returns an array containing all permutations
of the string.
'''

def permsNoDupes(S):
    #TODO implement recursivley, setting a character and then calling permutation on alphabet minus that character

testinput = "xyz"
testoutput = ["xyz","xzy", "yxz","yzx","zxy","zyx"]
assert permsNoDupes(testinput) == testoutput # Kind of a cheat, since output array order matters.
                                             # Should probably call a sort method on test outputs to be fair.

