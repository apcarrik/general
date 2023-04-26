import os, binascii

class Codec:

	_short_len: int = 9

	def __init__(self):
		self.long_to_short: Dict[str,str] = {}
		self.short_to_long: Dict[str,str] = {}

	def _create_short_url(self) -> string:
		shortURL: str = binascii.b2a_hex(os.urandom(self._short_len))		
		return shortURL

	def encode(self, long_url: string) -> string:
		if long_url not in self.long_to_short:
			short_url = self._create_short_url()
			self.long_to_short[long_url] = short_url
			self.short_to_long[short_url] = long_url
		return self.long_to_short[long_url]

	def decode(self, short_url : string) -> string:
		return self.short_to_long[short_url]

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))
