package dom

type AbstractRange struct {
	startContainer *Node
	startOffset    uint64
	endContainer   *Node
	endOffset      uint64
	collapsed      bool
}
