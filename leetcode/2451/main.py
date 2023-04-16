class Solution:
  def oddString(self, words: List[str]) -> str:
    difference_integer_array = []
    word1 = words[0]
    word2 = words[1]
    
    for wi in range(2,len(words)):
        word3 = words[wi]
        for ci in range(len(word1)-1):
            word1_c1 = ord(word1[ci])
            word1_c2 = ord(word1[ci+1])
            word2_c1 = ord(word2[ci])
            word2_c2 = ord(word2[ci+1])
            word3_c1 = ord(word3[ci])
            word3_c2 = ord(word3[ci+1])

            word1_diff = word1_c2 - word1_c1
            word2_diff = word2_c2 - word2_c1
            word3_diff = word3_c2 - word3_c1
            
            if word1_diff != word2_diff and word1_diff != word3_diff:
                return words[0]
            elif word2_diff != word1_diff and word2_diff != word3_diff:
                return words[1]
            elif word3_diff != word1_diff and word3_diff != word2_diff:
                return words[wi]

    return ""
