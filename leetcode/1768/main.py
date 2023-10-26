class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:        
        min_len = min(len(word1),len(word2))
        new_str = ""
        for i in range(min_len):
            new_str += word1[i] + word2[i]
        if len(word1) > min_len:
            new_str += word1[min_len:]
        elif len(word2) > min_len:
            new_str += word2[min_len:]
        return new_str
