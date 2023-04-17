class Solution:
	def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
		max_candies = max(candies)
		for i,c in enumerate(candies):
			candies[i] = c + extraCandies >= max_candies
		return candies
