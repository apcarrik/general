'''
Solution for problem 7.1 in Cracking the Coding Interview
Implements a deck of cards for blackjack.
'''
import numpy as np

# class card implements a single playing card
class card:
    def __init__(self, suit, value):
        self.suit = suit
        self.value = value

# class cardDeck implements a standard 52-card playing card deck
class cardDeck:
    def __init__(self): # Start with full draw pile, empty discard pile
        self.drawpile = self.newdeck()
        self.discardpile = [None] * len(self.drawpile)

    # newdeck returns a new deck in cannonical order
    def newdeck(self):
        suits = ["Clubs", "Diamonds", "Hearts", "Spades"]
        values = ["Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"]
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
        np.random.shuffle(targetpile)
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
            return None

        # Draw card
        if(targetpile[0] == None):
            return None
        drawcard = targetpile[0]
        for i in range(1, len(targetpile)):
            targetpile[i-1] = targetpile[i]
            if(targetpile[i] == None):
                break
        targetpile[51] = None
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

# test_card runs tests on card class
def test_card():
    newcard = card("Diamonds","Jack")
    assert newcard.suit == "Diamonds"
    assert newcard.value == "Jack"

    newcard = card("asdf","fdsa")
    assert newcard.suit != "Diamonds"
    assert newcard.value != "Jack"
    assert newcard.suit == "asdf"
    assert newcard.value == "fdsa"

# test_cardDeck runs tests on carddeck class
def test_cardDeck():
    # Test new deck has full draw pile and empty discard pile
    ndeck = cardDeck()
    assert len(ndeck.drawpile) == 52
    assert ndeck.drawpile[0] != None
    assert ndeck.drawpile[51] != None
    assert len(ndeck.discardpile) == 52
    assert ndeck.discardpile[0] == None
    assert ndeck.discardpile[51] == None

    # Test shuffle - ensure that not every card is in the same position in the deck after shuffling
    comparedeck = [None]*len(ndeck.drawpile)
    i=0
    for tcard in ndeck.drawpile:
        comparedeck[i] = tcard
        i+=1
    ndeck.shuffle("draw")
    allsame = True
    i=0
    for tcard in ndeck.drawpile:
        if(comparedeck[i] != tcard):
            allsame = False
        i+=1
    assert allsame == False

    # Test drawfrom - ensure that, after a card is drawn from the deck, no card in the deck has the same suit/value
    decksize = len(ndeck.drawpile)
    drawncard = ndeck.drawfrom("draw")
    assert drawncard != None
    # assert decksize == len(ndeck.drawpile) + 1
    for tcard in ndeck.drawpile:
        if(tcard != None):
            assert (drawncard.suit != tcard.suit or drawncard.value != tcard.value)

    # Test placeonto - ensure that, after a card is placed onto the deck, that that card is in the deck
    decksize = len(ndeck.drawpile)
    assert ndeck.placeonto("draw",drawncard) == True
    # assert decksize == len(ndeck.drawpile) - 1
    cardfound = False
    for tcard in ndeck.drawpile:
        if(tcard != None):
            if (drawncard.suit == tcard.suit and drawncard.value == tcard.value):
                cardfound = True
    assert cardfound == True

# Run unit tests
test_card()
test_cardDeck()

# TODO: Implement blackjack