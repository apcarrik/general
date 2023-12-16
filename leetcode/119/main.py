# class Solution:
#     def getRow(self, rowIndex: int) -> List[int]:

class Solution:
	def getRow(self, rowIndex: int) -> List[int]:
		prevRow: List[int] = [1]
		if rowIndex == 0:
			return prevRow
		for i in range(1,rowIndex+1):
			newRow: List[int] = [1]
			for j in range(1,i):
				newRow.append(prevRow[j-1]+prevRow[j])
			newRow.append(1)
			prevRow = newRow
		return prevRow
