package dom

type NodeList struct {
	nodes []*Node
}

type NodeListInterface interface {
	GetItem(index uint64) *Node
	GetLength() uint64
	Iterate() <-chan *Node
}
