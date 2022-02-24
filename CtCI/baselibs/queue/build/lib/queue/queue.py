# A queue uses a FIFO ordering. It contains the following operations:
#   - add(item) : Add an item to the end of the list.
#   - remove(): Remove the first item in the list.
#   - peek(): Return the top of the stack.
#   - isEmpty(): Return true if and only if the queue is empty.

from linkedlist.linkedlist import SLLNode, SLL

class QueueNode(SLLNode):
    def __init__(self, datum, nxt=None):
        super().__init__(datum, nxt=nxt)

class Queue(SLL):
    def __init__(self):
        super().__init__()

    def add(self, item):
        super().append(item)

    def remove(self):
        try:
            return super().pop().datum
        except RuntimeError as e:
            raise RuntimeError(str.replace(str(e), "linked list", "queue"))

    def append(self):
        raise RuntimeError("Method 'append' is not defined for queue")

    def pop(self):
        raise RuntimeError("Method 'pop' is not defined for queue")
