package dom

type NodeType int

const (
	ELEMENT_NODE                NodeType = 1
	ATTRIBUTE_NODE              NodeType = 2
	TEXT_NODE                   NodeType = 3
	CDATA_SECTION_NODE          NodeType = 4
	ENTITY_REFERENCE_NODE       NodeType = 5
	ENTITY_NODE                 NodeType = 6
	PROCESSING_INSTRUCTION_NODE NodeType = 7
	COMMENT_NODE                NodeType = 8
	DOCUMENT_NODE               NodeType = 9
	DOCUMENT_TYPE_NODE          NodeType = 10
	DOCUMENT_FRAGMENT_NODE      NodeType = 11
	NOTATION_NODE               NodeType = 12
)

// Node interface represents the Node behaviors in DOM.
// type Node interface {
// 	GetNodeType() int
// 	GetNodeName() string
// 	GetBaseURI() string
// 	IsConnected() bool
// 	GetOwnerDocument() *Document
// 	GetRootNode(options *GetRootNodeOptions) *Node
// 	GetParentNode() *Node
// 	GetParentElement() *Element
// 	HasChildNodes() bool
// 	GetChildNodes() []Node
// 	GetFirstChild() *Node
// 	GetLastChild() *Node
// 	GetPreviousSibling() *Node
// 	GetNextSibling() *Node
// 	GetNodeValue() string
// 	SetNodeValue(value string)
// 	GetTextContent() string
// 	SetTextContent(content string)
// 	Normalize()
// 	CloneNode(deep bool) *Node
// 	IsEqualNode(otherNode *Node) bool
// 	IsSameNode(otherNode *Node) bool
// 	CompareDocumentPosition(otherNode *Node) int
// 	Contains(otherNode *Node) bool
// 	LookupPrefix(namespace string) string
// 	LookupNamespaceURI(prefix string) string
// 	IsDefaultNamespace(namespace string) bool
// 	InsertBefore(newNode *Node, referenceNode *Node) *Node
// 	AppendChild(newNode *Node) *Node
// 	ReplaceChild(newNode *Node, oldNode *Node) *Node
// 	RemoveChild(childNode *Node) *Node
// }

type Node struct {
	Type       NodeType
	TagName    string
	Content    string
	Attributes []Attribute
	Children   []*Node
	Parent     *Node
}

// Add a child node to the current node
func (n *Node) AddChild(child *Node) {
	child.Parent = n
	n.Children = append(n.Children, child)
}

// Remove a child node
func (n *Node) RemoveChild(child *Node) {
	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			break
		}
	}
}
