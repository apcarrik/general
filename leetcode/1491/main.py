class Salaries:

	def __init__(self, salaries) -> None:
		self.salaries: List[int] = salaries

	def getModifiedSalaries(self) -> float:
		if len(self.salaries) > 2:
			return (sum(self.salaries) - min(self.salaries) - max(self.salaries)) / (len(self.salaries) - 2)
		return 0
        
class Solution:
    def average(self, salary: List[int]) -> float:
        salaries = Salaries(salary)
        return salaries.getModifiedSalaries()
