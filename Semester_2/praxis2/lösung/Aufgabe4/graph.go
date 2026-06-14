package aufgabe4

type Node struct {
	Label      string
	neighbours []*Node
}

func NewNode(label string) *Node {
	return &Node{Label: label}
}

func (n *Node) AddNeighbourNode(node *Node) {
	n.neighbours = append(n.neighbours, node)
}
