package main

type BoundaryPointsType uint8

const (
	START_TO_START BoundaryPointsType = 0
	START_TO_END   BoundaryPointsType = 1
	END_TO_END     BoundaryPointsType = 2
	END_TO_START   BoundaryPointsType = 3
)

type Range struct {
	AbstractRange
	commonAncestorContainer *Node
}

type RangeInterface interface {
	setStart(node *Node, offset uint64)
	setEnd(node *Node, offset uint64)
	setStartBefore(node *Node)
	setStartAfter(node *Node)
	setEndBefore(node *Node)
	setEndAfter(node *Node)
	collapse(toStart bool)
	selectNode(node *Node)
	selectNodeContents(node *Node)
	compareBoundaryPoints(how BoundaryPointsType, sourceRange *Range) uint16
	deleteContents()
	extractContents() *DocumentFragment
	cloneContents() *DocumentFragment
	insertNode(node *Node)
	surroundContents(newParent *Node)
	cloneRange() *Range
	detach()
	isPointInRange(node *Node, offset uint64) bool
	comparePoint(node *Node, offset uint64) uint16
	intersectsNode(node *Node) bool
	stringify() string
}
