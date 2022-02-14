import unittest
from graph.graph import GraphNode, Graph, Tree

class TestGraphNode(unittest.TestCase):

    def setUp(self):
        self.n1 = GraphNode()
        self.n2 = GraphNode(datum="data2")
        self.n3 = GraphNode(children=[self.n1, self.n2])

    def test_GraphNode__init__(self):
        # test n1 init
        self.assertEqual(self.n1.datum, None)
        self.assertEqual(len(self.n1.children), 0)
        
        # test n2 init
        self.assertEqual(self.n2.datum, "data2")
        self.assertEqual(len(self.n2.children), 0)
        
        # test n3 init
        self.assertEqual(len(self.n3.children), 2)

    def test_GraphNode_add_child(self):
        # test adding one child
        self.assertEqual(len(self.n1.children), 0)
        self.n1.add_child(self.n2)
        self.assertEqual(len(self.n1.children), 1)

        # test adding itself as child
        self.assertEqual(len(self.n3.children), 2)
        self.n3.add_child(self.n3)
        self.assertEqual(len(self.n3.children), 3)
        
        # test adding the same child twice
        self.assertEqual(len(self.n3.children), 3)
        self.n3.add_child(self.n1)
        self.assertEqual(len(self.n3.children), 3)
        
    def test_GraphNode_remove_child(self):
        # test removing one child
        self.assertEqual(len(self.n3.children), 2)
        self.n3.remove_child(self.n1)
        self.assertEqual(len(self.n3.children), 1)

        # test removing the same child twice
        with self.assertRaises(ValueError):
            self.n3.remove_child(self.n1)

    def test_GraphNode_get_children(self):
        # test getting no children
        self.assertEqual(GraphNode(children=[]).get_children(), [])

        # test getting 1 children
        self.assertEqual(GraphNode(children=[self.n1]).get_children(), [self.n1])

        # test getting 2 children
        self.assertEqual(len(self.n3.get_children()), 2)

    def test_GraphNode__repr__(self):
        # test none representation
        self.assertEqual(repr(self.n1), "")

        # test string representation
        self.assertEqual(repr(self.n2), "data2")

class TestGraph(unittest.TestCase):

    def setUp(self):
        pass # TODO: Implement

class TestTree(unittest.TestCase):

    def setUp(self):
        pass # TODO: Implement

if __name__ == '__main__':
    unittest.main()