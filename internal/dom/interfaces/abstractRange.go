package dom

type AbstractRange struct {
	startContainer *Node
	startOffset    uint64
	endContainer   *Node
	endOffset      uint64
	collapsed      bool
}

type StaticRange struct {
	AbstractRange
}

type StaticRangeInterface interface {
	StaticRangeInit(startContainer *Node,
		startOffset uint64,
		endContainer *Node,
		endOffset uint64) StaticRange
}

func StaticRangeInit(startContainer *Node,
	startOffset uint64,
	endContainer *Node,
	endOffset uint64) StaticRange {
	return StaticRange{
		AbstractRange: AbstractRange{
			startContainer: startContainer,
			startOffset:    startOffset,
			endContainer:   endContainer,
			endOffset:      endOffset},
	}
}
