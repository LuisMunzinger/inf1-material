package aufgabe1

type Node struct {
	Value int
	Next  *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.Next == nil
}

func (n *Node) PushBack(value int) {
	if n.IsEmpty() {
		n.Next = NewNode()
		n.Value = value
		return
	}
	n.Next.PushBack(value)
}
