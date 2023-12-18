# class Solution:
#     def bestHand(self, ranks: List[int], suits: List[str]) -> str:

class Solution:
	def bestHand(self, ranks: List[int], suits: List[str]) -> str:
		# check for Flush
		if len(set(suits)) == 1:
			return "Flush"

		# check for Three of a Kind
		freqMap: dict[int: int] = {}
		for rank in ranks:
			if rank not in freqMap:
				freqMap[rank] = 1
			else:
				freqMap[rank]+=1
				if freqMap[rank] >= 3:
					return "Three of a Kind"

		# check for Pair
		if len(set(ranks)) < 5:
			return "Pair"

		# check for High Card
		return "High Card"
