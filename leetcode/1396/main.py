# class UndergroundSystem:

#     def __init__(self):
        

#     def checkIn(self, id: int, stationName: str, t: int) -> None:
        

#     def checkOut(self, id: int, stationName: str, t: int) -> None:
        

#     def getAverageTime(self, startStation: str, endStation: str) -> float:
        
class StationAverage:
		
	def __init__(self, averageTime: float, numTrips: int = 1):
		self.averageTime: float = averageTime
		self.numTrips: int = numTrips

class Trip:

	def __init__(self, startStation: str, startTime: int):
		self.startStation: str = startStation
		self.startTime: int = startTime

class UndergroundSystem:


	def __init__(self):
		self._stations: dict[str: dict[str: StationAverage]] = {}
		self._trips: dict[int: Trip] = {}

	def checkIn(self, id: int, startStation: str, t: int) -> None:
		if startStation not in self._stations:
			self._stations[startStation] = {}
		if id not in self._trips:
			self._trips[id] = Trip(startStation = startStation, startTime = t)

	def checkOut(self, id: int, endStation: str, t: int) -> None:
		tripTime: int = t - self._trips[id].startTime
		startStation: str = self._trips[id].startStation

		# update average
		if endStation in self._stations[startStation]:
			averageTime: float = self._stations[startStation][endStation].averageTime
			numTrips: int = self._stations[startStation][endStation].numTrips
			self._stations[startStation][endStation].averageTime = \
                ((averageTime * numTrips) + tripTime) / (numTrips + 1)
			self._stations[startStation][endStation].numTrips += 1
		else:
			self._stations[startStation][endStation] = \
                StationAverage(averageTime=tripTime, numTrips=1)

		# remove trip from _trips
		del(self._trips[id])

	def getAverageTime(self, startStation: str, endStation: str) -> float:
		return self._stations[startStation][endStation].averageTime

# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)
