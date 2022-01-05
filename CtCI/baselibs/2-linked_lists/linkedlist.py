# Singly Linked List
class SLLNode:
    def __init__(self, datum, next=None):
        self.datum = datum
        self.next = next

    def set_next(self, next):
        self.next = next

    def __repr__(self):
        return self.datum


class SLL:
    def __init__(self):
        self.head = None
        self.len = 0

    def __iter__(self):
        node = self.head
        while node is not None:
            yield node
            node = node.next

    def append(self, datum):
        new_node = SLLNode(datum)
        if self.head is None:
            self.head = new_node
        else:
            for current_node in self:
                pass
            current_node.set_next(new_node)
        self.len += 1

    def pop(self):
        if self.head == None:
            raise RuntimeError("linked list empty")
        old_head = self.head
        self.head = self.head.next
        self.len -=1
        return old_head

    def __len__(self):
        return self.len

    def __repr__(self):
        node_data = [n.datum for n in self]
        return " -> ".join(node_data)


# Doubly Linked List
class DLLNode:
    def __init__(self, datum, prev=None, nxt=None):
        self.datum = datum
        self.prev = prev
        self.next = nxt

    def set_next(self,nxt):
        self.next = nxt

    def set_prev(self, prev):
        self.prev = prev

    def __repr__(self):
        return self.datum


class DLL:
    def __init__(self):
        self.head = None
        self.tail = None
        self.len = 0

    def __iter__(self):
        node = self.head
        while node is not None:
            yield node
            node = node.next

    def append(self, datum):
        new_node = DLLNode(datum)
        if self.head is None:
            self.head = self.tail = new_node
        else:
            new_node.set_prev(self.tail)
            self.tail.set_next(new_node)
            self.tail = new_node
        self.len += 1

    def pop(self):  # take first element in list (head)
        if self.head is None:
            raise RuntimeError("DLL is empty.")
        old_head = self.head
        self.head = self.head.next
        self.head.set_prev(None)
        self.len -= 1
        return self.head.datum

    def __len__(self):
        return self.len

    def __repr__(self):
        node_data = [n.datum for n in self]
        return " -> ".join(node_data)