
class HNode:
    def __init__(self, char, count, left=None, right=None):
        self.char = char
        self.count = count
        self.left = left
        self.right = right

def traverseHuffmanTree(encodings, current, parent, encoding):
    if current.left == None and current.right == None:
        encodings[current.char] = encoding
    else:
        traverseHuffmanTree(encodings, current.left, current, encoding+"0")
        traverseHuffmanTree(encodings, current.right, current, encoding+"1")
    return encodings



def buildHuffmanCodeTree(charCounts):

    while(len(charCounts) > 1):
        # Sort character counts
        charCounts.sort(key=lambda x: x.count, reverse=True)

        # Pull out smallest 2 characters
        c1 = charCounts.pop()
        c2 = charCounts.pop()

        # Create new HNode of c1 and c2 concatenated, then point concatenated node to other two
        c3 = HNode(c1.char+c2.char, c1.count+c2.count, c1, c2)

        # Add new node to character counts
        charCounts.append(c3)

    # Assign encoding to characters
    encodings = {}
    encodings = traverseHuffmanTree(encodings, charCounts[0], None, "")
    return encodings


def encode():
    pass

def decode():
    pass

# -------------
charCounts = [
    HNode("a",5),
    HNode("b",4),
    HNode("c",3),
    HNode("d",2),
    HNode("e",2)
]
htree = buildHuffmanCodeTree(charCounts)
print(htree)