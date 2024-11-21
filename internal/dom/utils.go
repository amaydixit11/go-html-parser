package dom

// FindChildrenByTag finds all children with a specific tag name
func (n *Node) FindChildrenByTag(tagName string) []*Node {
	var result []*Node
	for _, child := range n.Children {
		if child.TagName == tagName {
			result = append(result, child)
		}
	}
	return result
}

// GetAttribute retrieves the value of an attribute by name
func (n *Node) GetAttribute(name string) string {
	for _, attr := range n.Attributes {
		if attr.Name == name {
			return attr.Value
		}
	}
	return ""
}
