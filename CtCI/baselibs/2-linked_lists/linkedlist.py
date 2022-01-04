# Singly Linked List
class sll_node:
    def __init__(self, data):
        self.data = None
        self.next = None

    def set(self,data):
        self.data=data

    def set_next(self,next):
        self.next=next

class sll:
    def __init__(self):
        self.head = None
        self.length = 0

    def append(self, data):
        new_node = sll_node(data)
        if self.head == None:
            self.head = new_node
        else:
            node_iter = self.head
            while node_iter.next != None:
                node_iter = node_iter.next
            node_iter.next = new_node
        self.length += 1

    def pop(self):
        if self.head == None:
            raise RuntimeError("linked list empty")
        old_head = self.head
        self.head = self.head.next
        self.length -=1
        return old_head

    def search_for_index_of_data(self,data):
        node_iter = self.head
        while node_iter != None:
            if node_iter.data == data:
                return True
        return False

    def return_ith_element(self,i):
        if i < 0:
            raise RuntimeError("i cannot be less than 0")
        node_iter = self.head
        for j in range(0,i+1):
            if node_iter == None:
                raise RuntimeError(f"linked list does not have {i} elements")
            if j == i:
                return node_iter.data
            node_iter = node_iter.next


# Doubly Linked List

class dll_node:
    def __init__(self, data, prev):
        self.data = data
        self.prev = prev
        self.next = next


class dll:
    def __init__(self)
        self.head = None
        self.tail = None
        self.length = 0

    def append(self, data):
        new_node = dll_node(data)
        if self.head == None:
            self.head = new_node
            self.tail = new_node
        else:
            old_tail = self.tail
            old_tail.next = new_node
            new_node.prev = old_tail
            self.tail = new_node
        self.length += 1

    def prepend(self, data):
        new_node = dll_node(data)
        if self.head == None:
            self.head = new_node
            self.tail = new_node
        else:
            old_head = self.head
            new_node.next = old_head
            old_head.prev = new_node
            self.head = new_node
        self.length += 1

    def pop(self):  # take first element in list (head)
        old_head = self.head
        self.head = self.head.next
        self.head.prev = None
        self.length -= 1
        return old_head


def snip(self):  # take last element in list (tail)
    old_tail = self.tail
    self.tail = self.tail.prev
    self.tail.next = None
    self.length -= 1
    return old_tail


def search_for_index_of_data(self, data):
    iterator = self.head
    i = 0
    while iterator != None:
        if iterator.data == data:
            return i
        iterator = iterator.next
        i +=1
    return False


def return_ith_element(self, i):
    if i > self.len or i < 0:
        raise RuntimeError(f"i must be between 0 and {self.length}")
    iterator = self.head
    for j in range(0, i + 1):
        if i == j:
            return iterator.data
        iterator = iterator.next
    raise RuntimeError(f"could not find index {i} element")

