'''
Solution for problem 10.2 in Cracking the Coding Interview
'''

# groupAnagrams takes an array of strings and sorts with anagrams next to one another
# assumes anagrams can have arbitrary spaces (spaces do not count as characters)
def groupAnagrams(strs):
    # Idea: create a hash table based on character counts, return all with the same character count together
    # Idea: hash function could be: assign a unique prime to each character, and make the hash function be the sum of
    # each character's associated prime times the count of that character
    return strs


# Test cases
testinput = ["Dormitory","Twelve plus one",
"School master","Here come dots",
"Conversation","Real fun",
"Listen","Joyful Fourth",
"Astronomer","Silent",
"The eyes","They see",
"A gentleman","Dirty room",
"Funeral","Elegant man",
"The Morse Code","Moon starer",
"Eleven plus two","The classroom",
"Slot machines","Voices rant on",
"Fourth of July","Cash lost in me"]
expectedoutput = ["Dormitory","Dirty room",
"School master","The classroom",
"Conversation","Voices rant on",
"Listen","Silent",
"Astronomer","Moon starer",
"The eyes","They see",
"A gentleman","Elegant man",
"Funeral","Real fun",
"The Morse Code","Here come dots",
"Eleven plus two","Twelve plus one",
"Slot machines","Cash lost in me",
"Fourth of July","Joyful Fourth"]# may need to change this order based on hash function
assert groupAnagrams(testinput) == expectedoutput #may need to make function to test strings are the same