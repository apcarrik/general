import unittest
from linkedlist.linkedlist import SLLNode, SLL, DLLNode, DLL


class TestSLLNode(unittest.TestCase):

    def setUp(self):
        self.n1 = SLLNode("d1")
        self.n2 = SLLNode("d2",self.n1)
        self.n3 = SLLNode("d3")

    def test_SLLNode__init__(self):
        self.assertEqual(self.n1.next, None)
        self.assertEqual(self.n2.next, self.n1)

    def test_SLLNode_set_next(self):
        self.n1.set_next(self.n3)
        self.assertEqual(self.n1.next, self.n3)

    def test_SLLNode__repr__(self):
        self.assertEqual(repr(self.n1), 'd1')

class TestSLL(unittest.TestCase):

    def setUp(self):
        self.data = ["d1", "d2"]
        self.l1 = SLL()
        self.l2 = SLL()
        self.l3 = SLL()
        self.l2.append(self.data[0])
        for datum in self.data:
            self.l3.append(datum)

    def test_SLL__init__(self):
        self.assertEqual(self.l1.head, None)
        self.assertEqual(self.l1.len, 0)

    def test_SLL__iter__(self):
        for n in self.l1:
            self.assertTrue(False)
        for l in [self.l2, self.l3]:
            for n in l:
                self.assertTrue(repr(n) in self.data)

    def test_SLL_append(self):
        len1 = len(self.l1)
        self.l1.append("foo")
        self.assertTrue(len(self.l1)-len1 == 1)
        self.assertTrue(repr(self.l1.pop())=="foo")

    def test_SLL_pop(self):
        self.assertTrue(repr(self.l3.pop())==self.data[0])

    def test_SLL__len__(self):
        self.assertEquals(len(self.l1),0)
        self.assertEquals(len(self.l2),1)
        self.assertEquals(len(self.l3),2)

    def test_SLL__repr__(self):
        self.assertEquals(repr(self.l1),'')
        self.assertEquals(repr(self.l2),'d1')
        self.assertEquals(repr(self.l3),'d1 -> d2')


# TODO: class TestDLLNode(unittest.TestCase):
# TODO: class TestDLL(unittest.TestCase):


if __name__=='__main__':
    unittest.main()