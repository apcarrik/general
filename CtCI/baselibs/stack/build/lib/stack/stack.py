# A stack uses a LIFO ordering. It uses the following operations:
#   - pop() : Remove the top item from the stack.
#   - push(item): Add an item to the top of the stack.
#   - peek(): Return the top of the stack.
#   - isEmpty(): Return true if and only if the stack is empty.

class StackNode:
    def __init__(self, datum, prev=None):
        self.datum = datum
        self.prev = prev

    def set_prev(self, prev):
        self.prev = prev

    def __repr__(self):
        return self.datum

class Stack:
    def __init__(self, current=None):
        self.current = current
        self.len = 0 if current is None else 1

    def __iter__(self):
        node = self.current
        while node is not None:
            yield node
            node = node.prev

    def push(self, datum):
        if self.current is None:
            self.current = StackNode(datum)
        else:
            newnode = StackNode(datum, self.current)
            self.current = newnode
        self.len += 1

    def pop(self):
        if self.current is None:
            raise RuntimeError("Can't pop from an empty stack.")
        return_datum = self.current.datum
        self.current = self.current.prev
        self.len -= 1
        return return_datum

    def peek(self):
        if self.current is None:
            raise RuntimeError("Can't peek an empty stack.")
        return self.current.datum

    def isEmpty(self):
        return True if self.len == 0 else False

    def __len__(self):
        return self.len

    def __repr__(self):
        data = [n.datum for n in self]
        return " <- ".join(data)


