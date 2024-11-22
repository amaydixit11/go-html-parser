package dom

type NodeType uint16

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

type DocumentPosition uint16

const (
	DOCUMENT_POSITION_DISCONNECTED            DocumentPosition = 0x01
	DOCUMENT_POSITION_PRECEDING               DocumentPosition = 0x02
	DOCUMENT_POSITION_FOLLOWING               DocumentPosition = 0x04
	DOCUMENT_POSITION_CONTAINS                DocumentPosition = 0x08
	DOCUMENT_POSITION_CONTAINED_BY            DocumentPosition = 0x10
	DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC DocumentPosition = 0x20
)

// Node interface represents the Node behaviors in DOM
type Node struct {
	nodeType        NodeType
	nodeName        string
	baseURI         string
	isConnected     bool
	ownerDocument   *Document
	parentNode      *Node
	parentElement   *Element
	childNodes      *NodeList
	firstChild      *Node
	lastChild       *Node
	previousSibling *Node
	nextSibling     *Node
	nodeValue       string
	textContent     string
}
type NodeInterface interface {
	EventTarget
	GetNodeType() NodeType
	GetNodeName() string
	GetBaseURI() string
	IsConnected() bool
	GetOwnerDocument() *Document
	GetRootNode(options *GetRootNodeOptions) *Node
	GetParentNode() *Node
	GetParentElement() *Element
	HasChildNodes() bool
	GetChildNodes() []*Node
	GetFirstChild() *Node
	GetLastChild() *Node
	GetPreviousSibling() *Node
	GetNextSibling() *Node
	GetNodeValue() string
	SetNodeValue(value string)
	GetTextContent() string
	SetTextContent(content string)
	Normalize()
	CloneNode(deep bool) *Node
	IsEqualNode(otherNode *Node) bool
	IsSameNode(otherNode *Node) bool
	CompareDocumentPosition(otherNode *Node) DocumentPosition
	Contains(otherNode *Node) bool
	LookupPrefix(namespace string) string
	LookupNamespaceURI(prefix string) string
	IsDefaultNamespace(namespace string) bool
	InsertBefore(newNode *Node, referenceNode *Node) *Node
	AppendChild(newNode *Node) *Node
	ReplaceChild(newNode *Node, oldNode *Node) *Node
	RemoveChild(childNode *Node) *Node
}

type GetRootNodeOptions struct {
	Composed bool
}

type NodeList struct {
	nodes []*Node
}

type NodeListInterface interface {
	GetItem(index uint64) *Node
	GetLength() uint64
	Iterate() <-chan *Node
}
