package aufgabe2

type Node struct {
	Value  int
	Next   *Node
	Length int
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
		n.Length = 1
		return
	}
	n.Next.PushBack(value)
	n.Length++
}
