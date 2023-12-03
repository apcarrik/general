class Solution:
    def addDigits(self, num: int) -> int:    
        if num == 0: 
            return 0
        mod_nine: int = num % 9
        if mod_nine == 0:
            return 9      
        return mod_nine
