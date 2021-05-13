'''
Solution for problem 7.1 in Cracking the Coding Interview
Implements a deck of cards for blackjack.
'''

# class card implements a single playing card
class card:
    def __init__(self, suit, value):
        self.suit = suit
        self.value = value

# class cardDeck implements a standard 52-card playing card deck
class cardDeck:
    def __init__(self): # Start with full draw pile, empty discard pile
        self.drawpile = newdeck(self)
        self.discardpile = [None] * len(self.drawpile)

    # newdeck returns a new deck in cannonical order
    def newdeck(self):
        suits = ["Clubs", "Diamonds", "Hearts", "Spades"]
        values = ["Ace", "2", "3", "4", "5", "6", "7", "8", "9", "Jack", "Queen", "King"]
        deck = [None]*(len(suits) * len(values))
        i=0
        for suit in suits:
            for val in values:
                deck[i] = card(suit, val)
                i += 1
        return deck

    # shuffle implements shuffling the specified deck. Returns true if no errors, False otherwise
    def shuffle(self, pile):
        # Validate pile type
        if(pile == "draw"):
            targetpile = self.drawpile
        elif(pile == "discard"):
            targetpile = self.discardpile
        else:
            print("Shuffle error: Pile type not defined")
            return False

        # Shuffle deck
        #TODO: Implement shuffle on shuffle pile
        return True

    # drawfrom draws a card from the specified deck. Returns the drawn card if there are no errors, False otherwise.
    def drawfrom(self, pile):
        # Validate pile type
        if (pile == "draw"):
            targetpile = self.drawpile
        elif (pile == "discard"):
            targetpile = self.discardpile
        else:
            print("Draw error: Pile type not defined")
            return False

        # Draw card
        if(targetpile[0] == None):
            return None
        drawcard = targetpile[0]
        for i in range(1, len(targetpile)):
            targetpile[i-1] = targetpile[i]
            if(targetpile[i] == None):
                break
        return drawcard

    # placeonto places specified card onto specified deck. Returns True if no errors, False otherwise
    def placeonto(self, pile, card):
        # Validate pile type
        if (pile == "draw"):
            targetpile = self.drawpile
        elif (pile == "discard"):
            targetpile = self.discardpile
        else:
            print("Place onto error: Pile type not defined")
            return False

        # Place card on top of deck
        if(targetpile[0] == None):
            targetpile[0] = card
        else:
            previouscard = targetpile[0]
            targetpile[0] = card
            for i in range (1, len(pile)):
                tmpcard = targetpile[i]
                targetpile[i] = previouscard
                previouscard = tmpcard
                if(previouscard == None):
                    break
        return True

