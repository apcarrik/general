# A graph is a collection of nodes with edges between them. There are many sub-types of graphs, and
# this file will implement undirected graphs, directed graphs, and trees including general trees,
# binary trees, binary (min) heap, and tries. The underlying graph will be stored in adjacency
# list representation.

import pprint

class GraphNode:
    def __init__(self, datum=None):
        self.datum = datum

    def __repr__(self):
        return self.datum if self.datum is not None else ""

class Graph:
    allowed_graph_types = ("undirected", "directed")

    def __init__(self, type='undirected'):
        if type not in Graph.allowed_graph_types:
            raise RuntimeError(f"Graph type {type} is not supported.")
        self.type = type
        self.nodes = {} # dictionary of node_id(int):node_ptr(GraphNode)
        self.adj_list = {} # dictionary of node_id(int):adjacent_nodes(set(node_id(int)))
        self.num_nodes = 0

    def add_edge(self, n1_id, n2_id):
        self.adj_list[n1_id].add(n2_id)
        if self.type == "undirected":
            self.adj_list[n2_id].add(n1_id)

    def remove_edge(self, n1_id, n2_id):
        self.adj_list[n1_id].remove(n2_id)
        if self.type == "undirected":
            self.adj_list[n2_id].remove(n1_id)

    def get_edges(self, n_id):
        return self.adj_list[n_id]

    def add_node(self, datum=None):
        new_node = GraphNode(datum)
        n_id = self.num_nodes
        self.num_nodes += 1
        self.nodes[n_id] = new_node
        self.adj_list[n_id] = set()
        return n_id

    def remove_node(self, rem_id): # removes all edges that include this node
        for n_id in self.adj_list.keys():
            self.adj_list[n_id].discard(rem_id)
        self.adj_list.pop(rem_id, None)
        self.nodes.pop(rem_id)

    def __repr__(self):
        pp = pprint.PrettyPrinter(indent=4)
        return pp.pformat(self.adj_list)


class Tree(Graph):
    def __init__(self, root=None):
        super().__init__(type="undirected")
        self.root = root

    def add_edge(self, *args):
        raise RuntimeError("Method 'add_edge' is not defined for Tree")

    def remove_edge(self, *args):
        raise RuntimeError("Method 'remove_edge' is not defined for Tree")

    def add_node(self, datum=None, parent_id=None):
        n_id = super().add_node(datum)
        if self.root is None:
            self.root=n_id
        else:
            # add edge to parent
            super().add_edge(n_id,parent_id)
        return n_id

    def remove_node(self, rem_id):
        super().remove_node(rem_id)
        if rem_id == self.root:
            self.root = None

    # TODO: implement binary tree

    # TODO: implement binary min heap

    # TODO: implement trie