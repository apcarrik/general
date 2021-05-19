'''
Solution for problem 8.3 in Cracking the Coding Interview
permsNoDupes takes an input string of unique characters (alphabet) and returns an array containing all permutations
of the string.
'''
def removeChar(rmch,S):
    retstr = ""
    for ch in S:
        if(ch != rmch):
            retstr += ch
    return retstr

def permsNoDupes(S):
    #TODO implement recursivley, setting a character and then calling permutation on alphabet minus that character
    perms = []
    for ch in S:
        Smod = removeChar(ch,S) # removes ch from S, keeping order otherwise
        if(len(Smod) > 0):
            smallperms = permsNoDupes(Smod)
        else:
            smallperms = [""]
        for perm in smallperms:
            perms.append(ch+perm)
    return perms

def inputMatchesExpected(input, expected):
    inlen = len(input)
    exlen = len(expected)
    if(inlen != exlen):
        return False
    else:
        for i in range(0,inlen):
            if(input[i] != expected[i]):
                return False
    return True



testinput = "xyz"
testoutput = ["xyz","xzy","yxz","yzx","zxy","zyx"]
assert inputMatchesExpected(permsNoDupes(testinput), testoutput)  # Kind of a cheat, since output array order matters.
                                                    # Should probably call a sort method on test outputs to be fair.
testinput = "rxyz"
testoutput = ["rxyz","rxzy","ryxz","ryzx","rzxy","rzyx",
              "xryz","xrzy","xyrz","xyzr","xzry","xzyr",
              "yrxz","yrzx","yxrz","yxzr","yzrx","yzxr",
              "zrxy","zryx","zxry","zxyr","zyrx","zyxr"]
assert inputMatchesExpected(permsNoDupes(testinput), testoutput)  # Kind of a cheat, since output array order matters.
                                                    # Should probably call a sort method on test outputs to be fair.

