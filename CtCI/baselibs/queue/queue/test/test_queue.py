import unittest
from queue.queue import QueueNode, Queue

class TestQueueNode(unittest.TestCase):

    def setUp(self):
        self.n1 = QueueNode("d1")
        self.n2 = QueueNode("d2",self.n1)
        self.n3 = QueueNode("d3")

    def test_QueueNode__init__(self):
        self.assertEqual(self.n1.next, None)
        self.assertEqual(self.n2.next, self.n1)

    def test_QueueNode_set_next(self):
        self.n1.set_next(self.n3)
        self.assertEqual(self.n1.next, self.n3)

    def test_QueueNode__repr__(self):
        self.assertEqual(repr(self.n1), 'd1')

class TestQueue(unittest.TestCase):

    def setUp(self):
        self.data = ["d1", "d2"]
        self.l1 = Queue()
        self.l2 = Queue()
        self.l3 = Queue()
        self.l2.add(self.data[0])
        for datum in self.data:
            self.l3.add(datum)

    def test_Queue__init__(self):
        self.assertEqual(self.l1.head, None)
        self.assertEqual(self.l1.len, 0)

    def test_Queue__iter__(self):
        for n in self.l1:
            self.assertTrue(False)
        for l in [self.l2, self.l3]:
            for n in l:
                self.assertTrue(repr(n) in self.data)

    def test_Queue_add(self):
        len1 = len(self.l1)
        self.l1.add("foo")
        self.assertTrue(len(self.l1)-len1 == 1)
        self.assertTrue(repr(self.l1.remove())=="foo")

    def test_Queue_remove(self):
        self.assertTrue(repr(self.l3.remove())==self.data[0])

    def test_Queue__len__(self):
        self.assertEquals(len(self.l1),0)
        self.assertEquals(len(self.l2),1)
        self.assertEquals(len(self.l3),2)

    def test_Queue__repr__(self):
        self.assertEquals(repr(self.l1),'')
        self.assertEquals(repr(self.l2),'d1')
        self.assertEquals(repr(self.l3),'d1 -> d2')

if __name__=='__main__':
    unittest.main()