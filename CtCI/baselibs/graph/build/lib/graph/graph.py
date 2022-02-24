# A graph is a collection of nodes with edges between them. There are many sub-types of graphs, and
# this file will implement undirected graphs, directed graphs, and trees including general trees,
# binary trees, binary (min) heap, and tries. The underlying graph will be stored in adjacency
# list representation.

from queue.queue import QueueNode, Queue # Need to fix not finding these classes
import pprint

class GraphNode:
    def __init__(self, datum=None):
        self.datum = datum

    def __repr__(self):
        return repr(self.datum)

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

    def add_node(self, datum=None):
        new_node = GraphNode(datum)
        n_id = self.num_nodes
        self.num_nodes += 1
        self.nodes[n_id] = new_node
        return n_id

    def remove_node(self, rem_id): # removes all edges that include this node
        for n_id in self.adj_list.keys():
            self.adj_list[n_id].discard(rem_id)
        self.adj_list.pop(rem_id)

    def __repr__(self):
        pp = pprint.PrettyPrinter(indent=4)
        return pp.pformat(self.adj_list)

# TODO: just redesigned graph, need to revisit tree
class TreeNode(GraphNode):
    def __init__(self, name, datum=None, children=[]):
        super().__init__(self, datum=datum, children=children)
        self.name = name

class Tree(Graph):
    # TODO: impelment basic tree
    def __init__(self, root=None):
        super().__init__(type="undirected")
        self.root = root

    def add_edge(self):
        raise RuntimeError("Method 'add_edge' is not defined for Tree")

    def remove_edge(self):
        raise RuntimeError("Method 'remove_edge' is not defined for Tree")

    def add_node(self, name, datum=None, parent_name=None):
        new_node = TreeNode(name, datum, [])
        if self.root is None:
            self.root=new_node
        else:
            parent_node = self._bfs(parent_name)
            if parent_node is None:
                raise RuntimeError(f"Parent node '{parent_name}' could not be found in tree")
            parent_node.add_child(new_node) # TODO: I think I want to add edges for graph representation
        return name

    def remove_node(self, name):
        pass # TODO

    # Implements breadth-first search to find the node pointer corresponding to the node name
    def _bfs(self, node_name):
        search_queue = Queue()
        search_queue.add(self.root)
        while len(search_queue) > 0:
            n = search_queue.remove()
            if n.name == node_name:
                return n
            for c in n.get_children():
                search_queue.add(c)
        return None




    def remove_node(self, name):
        for n in self.nodes["name"].get_children():
            self._remove_node_from_pointer(n)
        super().remove_node(name=name)


    # TODO: implement binary tree

    # TODO: implement binary min heap

    # TODO: implement trie