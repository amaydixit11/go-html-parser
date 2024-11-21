package dom

import "strings"

func (n *Node) Serialize() string {
	var builder strings.Builder
	n.serializeHelper(&builder)
	return builder.String()
}

func (n *Node) serializeHelper(builder *strings.Builder) {
	switch n.Type {
	case ElementNode:
		builder.WriteString("<" + n.TagName)
		for _, attr := range n.Attributes {
			builder.WriteString(" " + attr.Name + `="` + attr.Value + `"`)
		}
		builder.WriteString(">")
		for _, child := range n.Children {
			child.serializeHelper(builder)
		}
		builder.WriteString("</" + n.TagName + ">")
	case TextNode:
		builder.WriteString(n.Content)
	case CommentNode:
		builder.WriteString("<!-- " + n.Content + " -->")
	case DoctypeNode:
		builder.WriteString("<!DOCTYPE " + n.Content + ">")
	}
}
