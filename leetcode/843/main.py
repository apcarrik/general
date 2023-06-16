# """
# This is Master's API interface.
# You should not implement it, or speculate about its implementation
# """
# class Master:
#     def guess(self, word: str) -> int:
from secrets import randbelow
class Solution:
    def distance(self, w1: str, w2: str) -> int:
        score:int = 0
        for i in range(6):
            if w1[i] == w2[i]:
                score += 1
        return score

    def findSecretWord(self, words: List[str], master: 'Master') -> None:

        # try to pre-compute the most likely guesses, by computing how many other words have distance score > 0 to each word
        ranks: list[tuple(int, int)] = [] # tuple is defined as (count of words with distance >0 with this word, this words index) 
        for i in range(len(words)):
            dist: int = 0
            for j in range(len(words)):
                if i != j:
                    if self.distance(words[i], words[j]) > 0:
                        dist += 1
            ranks.append((-1*dist, randbelow( len(words)) ,i))
            
        # now, rank guesses by their distance score
        heapq.heapify(ranks)

        sameScore: set[int] = set()
        while True:
            # decide which word to guess
            print( "ranks:", ranks)
            guess: int = heapq.heappop(ranks)[2]

            # guess word
            print("guess:", guess, words[guess])
            score: int = master.guess(words[guess]) # note: this class relies on master to throw an error if guess is correct or if max number of guesses have been made
            print("score:", score)
            if score == 6:
                return
            else:
            # now, we need to update ranks to only consider words that have the same score as the current word
                deleted: int = 0
                for idx, tup in enumerate(ranks):
                    if self.distance(words[guess], words[tup[2]]) != score:
                        deleted+=1
                        ranks[idx] = ranks[-1]
                        ranks.pop()
                print("deleted", deleted)
                heapq.heapify(ranks)
