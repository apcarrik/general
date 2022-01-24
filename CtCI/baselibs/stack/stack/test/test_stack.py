import unittest
from stack.stack import StackNode, Stack

class TestStackNode(unittest.TestCase):

    def setUp(self):
        self.n1 = StackNode("Datum1")
        self.n2 = StackNode("Datum2", self.n1)
        self.n3 = StackNode("Datum3")

    def test_StackNode__init__(self):
        self.assertEqual(self.n1.datum, "Datum1")
        self.assertEqual(self.n1.prev, None)
        self.assertEqual(self.n2.prev, self.n1)

    def test_StackNode_set_prev(self):
        self.assertEqual(self.n3.prev, None)
        self.n3.set_prev(self.n2)
        self.assertEqual(self.n3.prev, self.n2)

    def test_StackNode__repr__(self):
        self.assertEqual(repr(self.n1), "Datum1")

class TestStack(unittest.TestCase):

    def setUp(self):
        self.data = ["d1", "d2"]
        self.s1 = Stack()
        self.s2 = Stack(StackNode(self.data[0]))
        self.s3 = Stack()
        for datum in self.data:
            self.s3.push(datum)

    def test_Stack__init__(self):
        self.assertEqual(self.s1.current, None)
        self.assertEqual(self.s1.len, 0)
        self.assertEqual(repr(self.s2.current), self.data[0])
        self.assertEqual(self.s2.len, 1)

    def test_Stack__iter__(self):
        for n in self.s1:
            self.assertTrue(False)
        for s in [self.s2]: #, self.s3]:
            for n in s:
                self.assertTrue(repr(n) in self.data)

    def test_Stack_push(self):
        len1 = len(self.s1)
        self.s1.push("foo")
        self.assertTrue(len(self.s1)-len1 == 1)
        self.assertTrue(self.s1.peek()=="foo")
        len2 = len(self.s2)
        self.s2.push("bar")
        self.assertTrue(len(self.s2)-len2 == 1)
        self.assertTrue(self.s2.peek()=="bar")


    def test_Stack_pop(self):
        self.assertTrue(self.s3.pop()==self.data[1])
        self.assertTrue(self.s3.pop()==self.data[0])

    def test_Stack_peek(self):
        self.assertTrue(self.s3.peek()==self.data[1])
        self.assertTrue(self.s3.peek()==self.data[1])

    def test_Stack_isEmpty(self):
        self.assertTrue(self.s1.isEmpty())
        self.assertFalse(self.s2.isEmpty())

    def test_Stack__len__(self):
        self.assertEquals(len(self.s1),0)
        self.assertEquals(len(self.s2),1)
        self.assertEquals(len(self.s3),2)

    def test_Stack__repr__(self):
        self.assertEquals(repr(self.s1),'')
        self.assertEquals(repr(self.s2),'d1')
        self.assertEquals(repr(self.s3),'d2 <- d1')

if __name__=='__main__':
    unittest.main()