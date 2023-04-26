""" Notes:
    - 9 -> 9
    - 99 -> 18 -> 9
    - 999 -> 27 -> 9
    - 9999 -> 36 -> 9
    - 999999999999 -> 108 -> 9
    - 9999999999999 -> 117 -> 9
    - 8 -> 8
    - 88 -> 16 -> 7
    - 888 -> 24 -> 6
    - 8888888 -> 56 -> 11 -> 2
    - 2147483647 -> 46 -> 10 -> 1
    - 999999999 -> 81 -> 9 

"""

class Solution:
    def addDigits(self, num: int) -> int:
        while num > 9:
            digit_sum: int = 0
            while num > 9:
                digit_sum += num % 10
                num //= 10
            digit_sum += num % 10
            num = digit_sum
            
        return num
        
