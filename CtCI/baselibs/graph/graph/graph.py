# A graph is a collection of nodes with edges between them. There are many sub-types of graphs, and
# this file will implement undirected graphs, directed graphs, and trees including general trees,
# binary trees, binary (min) heap, and tries. The underlying graph will be stored in adjacency
# list representation.

class GraphNode:
    def __init__(self, datum=None, children=[]):
        self.datum = datum
        self.children = children

    def add_child(self, child):
        if not isinstance(child, GraphNode):
            raise RuntimeError(f"Unsupported child type: {type(child)}")
        if child not in self.children:
            self.children.append(child)

    def remove_child(self, child):
        self.children.remove(child)

    def get_children(self):
        return self.children

    def __repr__(self):
        return self.datum if self.datum is not None else ""

class Graph:
    allowed_graph_types = ("undirected", "directed")

    def __init__(self, type='undirected', nodes={}):
        self.nodes = nodes
        if type not in Graph.allowed_graph_types:
            raise RuntimeError(f"Graph type {type} is not supported.")
        self.type = type

    def add_edge(self, n1, n2):
        self.nodes[n1].add_child(self.nodes[n2])
        if self.type == "undirected":
            self.nodes[n2].add_child(self.nodes[n1])

    def remove_edge(self, n1, n2):
        self.nodes[n1].remove_child(self.nodes[n2])
        if self.type == "undirected":
            self.nodes[n2].remove_child(self.nodes[n1])

    def add_node(self, name=None, datum=None, children=[]):
        new_node = GraphNode(datum, [self.nodes[child] for child in children])
        if name is None:
            name = len(self.nodes.keys())+1
        self.nodes[name] = new_node
        if self.type == "undirected":
            for child_name in children:
                self.nodes[child_name].add_child(new_node)
        return name

    def remove_node(self, name):
        adjacency_list = self.nodes[name].get_children()
        for (child_name, child_pointer) in self.nodes.items():
            if child_pointer in adjacency_list:
                self.nodes[name].remove_child(child_pointer)
                if self.type == "undirected":
                    self.nodes[child_name].remove_child(self.nodes[name])
        self.nodes.pop(name)

    def __repr__(self):
        adjacency_list = f"graph_type:{self.type};\nNodes:\n"
        for (name, node) in list(self.nodes.items()):
            adjacency_list += f"name:{name}, datum:{repr(node)}, children:{node.get_children()};\n"
        return adjacency_list

# TODO: impelment basic tree
class Tree:
    pass
    # TODO: implement binary tree

    # TODO: implement binary min heap

    # TODO: implement trie