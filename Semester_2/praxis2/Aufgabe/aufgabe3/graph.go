package aufgabe3

type Node struct {
	Label      string
	neighbours []*Node
}

func NewNode(label string) *Node {
	return &Node{Label: label}
}

func (n *Node) AddNeighbours(labels ...string) {
	for _, label := range labels {
		n.neighbours = append(n.neighbours, NewNode(label))
	}
}
