package dom

func (n *Node) Traverse(level int) {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}

	switch n.Type {
	case ElementNode:
		println(prefix + "<" + n.TagName + ">")
	case TextNode:
		println(prefix + n.Content)
	case CommentNode:
		println(prefix + "<!-- " + n.Content + " -->")
	case DoctypeNode:
		println(prefix + "<!DOCTYPE " + n.Content + ">")
	}

	for _, child := range n.Children {
		child.Traverse(level + 1)
	}

	if n.Type == ElementNode {
		println(prefix + "</" + n.TagName + ">")
	}
}
