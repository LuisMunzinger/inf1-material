package aufgabe5

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) IsLeaf() bool {
	return n.Left.IsEmpty() && n.Right.IsEmpty()
}

func (n *Node) Add(value int) {
	if n.IsEmpty() {
		n.Left = NewNode()
		n.Right = NewNode()
		n.Value = value
		return
	}
	if value < n.Value {
		n.Left.Add(value)
		return
	}
	n.Right.Add(value)
}
