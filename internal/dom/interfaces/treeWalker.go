package main

type TreeWalker struct {
	root        *Node
	whatToShow  WhatToShowType
	filter      *NodeFilter
	currentNode *Node
}

type TreeWalkerInterface interface {
	parentNode() *Node
	firstChild() *Node
	lastChild() *Node
	previousSibling() *Node
	nextSibling() *Node
	previousNode() *Node
	nextNode() *Node
}
