'''
Solution for problem 1.3 in Cracking the Coding Interview
URLify_Brute implements brute force solution that copies strings a lot. Could be optimized for speed/memory.
URLify_instring implements solution that takes a bytearray and replaces spaces with '%20' within the input bytearray. Much more space efficient.
'''

# URLify_Brute takes a string and integer length and returns a string with all space characters replaced with '%20'
def URLify_Brute(strg, leng):
    words = []
    tmpword = ""
    for ch in strg:
        if(ch == " "):
            if tmpword:
                words.append(tmpword)
            tmpword = ""
        else:
            tmpword+=ch
    
    #Overwrite string with new string
    new_str = []
    for word in words:
        if(new_str):
            new_str+="%20"
            new_str+=word
        else:
            new_str = word
    return new_str

# URLify_instring takes a bytearray and integer length, and returns the bytearray with all space characters replaced with '%20'
def URLify_instring(barr, leng):
    lastspaceptr = len(barr)-1
    charptr = len(barr)-2
    # Move charptr to last char in bytearray
    while (barr[charptr] == 0x20): #0x20 is ascii space character
        charptr -= 1
    # Start moving charcteris in words to end, and eplace intermieary " " with "%2d"
    while (lastspaceptr != charptr):
        while barr[charptr] != 0x20:
            barr[lastspaceptr] = barr[charptr]
            barr[charptr] = 0x20
            lastspaceptr -= 1
            charptr -= 1
        # Now, charptr = " ", so we need ot inject "%20" just before last space ptr
        barr[lastspaceptr] = 0x30 #0x30 is ascii '0'' character
        barr[lastspaceptr - 1] = 0x32 #0x32 is ascii '2' character
        barr[lastspaceptr - 2] = 0x25 #0x25 is ascii '%' character
        lastspaceptr -= 3
        charptr -= 1
    return barr

#Testing
teststr = "Mr John Smith    "
testlen = 13
outputstr = "Mr%20John%20Smith"
assert (URLify_Brute(teststr,testlen) == outputstr)

testbytearr = bytearray(teststr, 'ascii')
outputbytearr = bytearray('Mr%20John%20Smith', 'ascii')
assert URLify_instring(testbytearr, testlen)  == outputbytearr