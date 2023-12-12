class MyHashSet:
	def __init__(self):
		# decide how large array will be
		self.size = 5000
		self.hashTable = [None] * self.size

	def getHashValue(self, key: int) -> int:
		return key % self.size

	def add(self, key: int) -> None:	
		hashVal: int = self.getHashValue(key)
		if self.hashTable[hashVal] is not None:
			for val in self.hashTable[hashVal]:
				if val == key:
					return
			self.hashTable[hashVal].append(key)
		else:
			self.hashTable[hashVal] = [key]

	def contains(self, key: int) -> bool:
		hashVal: int = self.getHashValue(key)
		if self.hashTable[hashVal] is not None: 
			for val in self.hashTable[hashVal]:
				if val == key:
					return True
		return False

	def remove(self, key: int) -> None:		
		hashVal: int = self.getHashValue(key)
		if self.hashTable[hashVal] is not None:
			if len(self.hashTable[hashVal]) == 1 and self.hashTable[hashVal][0] == key:
				self.hashTable[hashVal] = None
			else:
				for idx, val in enumerate(self.hashTable[hashVal]):
					if val == key:
						self.hashTable[hashVal] = self.hashTable[hashVal][:idx] + self.hashTable[hashVal][idx+1:] if idx != len(self.hashTable[hashVal])-1 else self.hashTable[hashVal][:idx]

        


# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)
