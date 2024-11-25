package main

type NodeIterator struct {
	root                       *Node
	referenceNode              *Node
	pointerBeforeReferenceNode bool
	whatToShow                 WhatToShowType
	filter                     *NodeFilter
}

type NodeIteratorInterface interface {
	nextNode() *Node
	previousNode() *Node
	detach()
}
