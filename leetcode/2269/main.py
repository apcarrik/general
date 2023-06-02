# class Solution:
#     def divisorSubstrings(self, num: int, k: int) -> int:


class Solution:
	def divisorSubstrings(self, num: int, k: int) -> int:
		# convert num to string
		num_str: str = str(num)

		# check all k-length substrings
		k_beauty: int = 0
		for start in range(len(num_str)-k+1):
			if int(num_str[start:start+k]) != 0 and num % int(num_str[start:start+k]) == 0:
				k_beauty+=1

		return k_beauty
