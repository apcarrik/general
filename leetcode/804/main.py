# class Solution:
#     def uniqueMorseRepresentations(self, words: List[str]) -> int:
class Solution:
    def _transform(self, word: str) -> str:
        transformation: str = ""
        morseDict: dict[str: str] = {"a": ".-",
        "b": "-...",
        "c": "-.-.",
        "d": "-..",
        "e": ".",
        "f": "..-.",
        "g": "--.",
        "h": "....",
        "i": "..",
        "j": ".---",
        "k": "-.-",
        "l": ".-..",
        "m": "--",
        "n": "-.",
        "o": "---",
        "p": ".--.",
        "q": "--.-",
        "r": ".-.",
        "s": "...",
        "t": "-",
        "u": "..-",
        "v": "...-",
        "w": ".--",
        "x": "-..-",
        "y": "-.--",
        "z": "--.."}

        for letter in word:
            transformation += morseDict[letter]
        return transformation

    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        transformations: set[str] = set()
        for word in words:
            transformations.add(self._transform(word))
        return len(transformations)
