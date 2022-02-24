import unittest
from graph.graph import GraphNode, Graph, Tree

# TODO: Just redesigned Graph, need to update tests
class TestGraphNode(unittest.TestCase):

    def setUp(self):
        self.n1 = GraphNode()
        self.n2 = GraphNode(datum="data2")

    def test_GraphNode__init__(self):
        # test n1 init
        self.assertEqual(self.n1.datum, None)

        # test n2 init
        self.assertEqual(self.n2.datum, "data2")

    def test_GraphNode__repr__(self):
        # test none representation
        self.assertEqual(repr(self.n1), "")

        # test string representation
        self.assertEqual(repr(self.n2), "data2")

class TestGraph(unittest.TestCase):

    def setUp(self):
        self.g1 = Graph()

        self.g2 = Graph(type="directed")

        self.g3 = Graph()
        self.g3.add_node("n1")
        self.g3.add_node("n2")
        self.g3.add_node("n3")
        self.g3.add_edge(0,1)

        self.g4 = Graph(type="directed")
        self.g4.add_node("n1")
        self.g4.add_node("n2")
        self.g4.add_node("n3")
        self.g4.add_edge(0,1)
        self.g4.add_edge(2,0)

    def test_Graph__init__(self):
        # test no args passed
        self.assertEqual(self.g1.type, "undirected")
        self.assertEqual(len(self.g1.nodes), 0)

        # test type arg passed
        self.assertEqual(self.g2.type, "directed")
        self.assertEqual(len(self.g2.nodes), 0)

    def test_Graph_add_edge(self):
        # test undirected graph
        self.assertEqual(repr(self.g3), "{0: {1}, 1: {0}, 2: set()}")

        # test directed graph
        self.assertEqual(repr(self.g4), "{0: {1}, 1: set(), 2: {0}}")

    def test_Graph_remove_edge(self):
        # test undirected graph
        self.g3.remove_edge(0,1)
        self.assertEqual(repr(self.g3), "{0: set(), 1: set(), 2: set()}")

        # test directed graph
        self.g4.remove_edge(0,1)
        self.assertEqual(repr(self.g4), "{0: set(), 1: set(), 2: {0}}")

    def test_Graph_get_edges(self):
        # test undirected graph
        self.assertEqual(repr(self.g3.get_edges(0)), "{1}")
        self.assertEqual(repr(self.g3.get_edges(1)), "{0}")

        # test directed graph
        self.assertEqual(repr(self.g4.get_edges(0)), "{1}")
        self.assertEqual(repr(self.g4.get_edges(1)), "set()")

    def test_Graph_add_node(self):
        # test adding single node
        self.assertEqual(len(self.g1.nodes), 0)
        self.assertEqual(self.g1.add_node("n1"), 0)
        self.assertEqual(len(self.g1.nodes), 1)

        # test adding second nodes
        self.assertEqual(len(self.g1.nodes), 1)
        self.assertEqual(self.g1.add_node("n2"), 1)
        self.assertEqual(len(self.g1.nodes), 2)

    def test_Graph_remove_node(self):
        # test removing single node
        self.g1.add_node("n1")
        self.g1.add_node("n2")
        self.assertEqual(len(self.g1.nodes), 2)
        self.g1.remove_node(1)
        self.assertEqual(len(self.g1.nodes), 1)

    def test_Graph__repr__(self):
        # test empty graph
        self.assertEqual(repr(self.g1), "{}")

        # test nodes with no edges
        self.g1.add_node()
        self.g1.add_node()
        self.assertEqual(repr(self.g1), "{0: set(), 1: set()}")

        # test nodes with edge
        self.assertEqual(repr(self.g3), "{0: {1}, 1: {0}, 2: set()}")

class TestTree(unittest.TestCase):

    def setUp(self):
        self.T1 = Tree()
        self.T2 = Tree()
        self.T2.add_node("n1")
        self.T2.add_node("n2", 0)

    def test_Tree__init__(self):
        self.assertIsNone(self.T1.root)
        self.assertEqual(self.T1.type, "undirected")

    def test_Tree_add_edge(self):
        with self.assertRaises(RuntimeError):
            self.T1.add_edge()

    def test_Tree_remove_edge(self):
        with self.assertRaises(RuntimeError):
            self.T1.remove_edge()

    def test_Tree_add_node(self):
        self.assertEqual(repr(self.T1), "{}")
        self.assertEqual(repr(self.T2), "{0: {1}, 1: {0}}")

    def test_Tree_remove_node(self):
        self.T2.remove_node(1)
        self.assertEqual(repr(self.T2), "{0: set()}")

if __name__ == '__main__':
    unittest.main()