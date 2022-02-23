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
        self.g1 = Graph()
        self.g2 = Graph(type="directed")
        self.g3 = Graph(nodes={"n1": GraphNode()})
        self.g4 = Graph(type="directed", nodes={
            "n1": GraphNode("n1", children=[]),
            "n2": GraphNode("n2", children=[])
        })
        self.g5 = Graph(type="undirected", nodes={
            "n1": GraphNode("n1", children=[]),
            "n2": GraphNode("n2", children=[])
        })

    def test_Graph__init__(self):
        # test no args passed
        self.assertEqual(self.g1.type, "undirected")
        self.assertEqual(len(self.g1.nodes), 0)

        # test type arg passed
        self.assertEqual(self.g2.type, "directed")
        self.assertEqual(len(self.g2.nodes), 0)

        # test nodes arg passed
        self.assertEqual(self.g3.type, "undirected")
        self.assertEqual(len(self.g3.nodes), 1)

    def test_Graph_add_edge(self):
        # test with directed graph
        self.assertEqual(len(self.g4.nodes["n1"].get_children()), 0)
        self.assertEqual(len(self.g4.nodes["n2"].get_children()), 0)
        self.g4.add_edge("n1", "n2")
        self.assertEqual(len(self.g4.nodes["n1"].get_children()), 1)
        self.assertEqual(len(self.g4.nodes["n2"].get_children()), 0)

        # test with undirected graph
        self.assertEqual(len(self.g5.nodes["n1"].get_children()), 0)
        self.assertEqual(len(self.g5.nodes["n2"].get_children()), 0)
        self.g5.add_edge("n1", "n2")
        self.assertEqual(len(self.g5.nodes["n1"].get_children()), 1)
        self.assertEqual(len(self.g5.nodes["n2"].get_children()), 1)

    def test_Graph_remove_edge(self):
        # test with directed graph
        self.g4.add_edge("n2", "n1")
        self.assertEqual(len(self.g4.nodes["n1"].get_children()), 0)
        self.assertEqual(len(self.g4.nodes["n2"].get_children()), 1)
        self.g4.remove_edge("n2", "n1")
        self.assertEqual(len(self.g4.nodes["n1"].get_children()), 0)
        self.assertEqual(len(self.g4.nodes["n2"].get_children()), 0)

        # test with undirected graph
        self.g5.add_edge("n1","n2")
        self.assertEqual(len(self.g5.nodes["n1"].get_children()), 1)
        self.assertEqual(len(self.g5.nodes["n2"].get_children()), 1)
        self.g5.remove_edge("n1", "n2")
        self.assertEqual(len(self.g5.nodes["n1"].get_children()), 0)
        self.assertEqual(len(self.g5.nodes["n2"].get_children()), 0)

    def test_Graph_add_node(self):
        # test adding single node
        self.assertEqual(len(self.g4.nodes.keys()), 2)
        self.g4.add_node("n3","n3",[])
        self.assertEqual(len(self.g4.nodes.keys()), 3)

    def test_Graph_remove_node(self):
        # test adding single node
        self.assertEqual(len(self.g4.nodes.keys()), 2)
        self.g4.remove_node("n2")
        self.assertEqual(len(self.g4.nodes.keys()), 1)
        self.assertEqual(list(self.g4.nodes.keys())[0], "n1")

    def test_Graph__repr__(self):
        self.assertEqual(repr(self.g5), \
             "graph_type:undirected;\nNodes:\nname:n1, datum:n1, children:[];\nname:n2, datum:n2, children:[];\n")

class TestTree(unittest.TestCase):

    def setUp(self):
        pass # TODO: Implement

if __name__ == '__main__':
    unittest.main()