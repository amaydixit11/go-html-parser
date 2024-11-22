package dom

type Document struct {
	Node
	implementation  DOMImplementation
	URL             string
	documentURI     string
	compatMode      string
	characterSet    string
	charset         string
	inputEncoding   string
	contentType     string
	doctype         DocumentType
	documentElement *Element
}
type DocumentInterface interface {
	NodeInterface
	GetImplementation() *DOMImplementation
	GetURL() string
	GetDocumentURI() string
	GetCompatMode() string
	GetCharacterSet() string
	GetCharset() string
	GetInputEncoding() string
	GetContentType() string
	GetDoctype() *DocumentType
	GetDocumentElement() *Element
	GetElementsByTagName(qualifiedName string) *HTMLCollection
	GetElementsByTagNameNS(namespace string, localName string) *HTMLCollection
	GetElementsByClassName(className string) *HTMLCollection
	CreateElement(localName string, options *ElementCreationOptions) *Element
	CreateElementNS(namespace string, qualifiedName string, options *ElementCreationOptions) *Element
	CreateDocumentFragment() *DocumentFragment
	CreateTextNode(data string) *Text
	CreateCDATASection(data string) *CDATASection
	CreateComment(data string) *Comment
	CreateProcessingInstruction(target string, data string) *ProcessingInstruction
	ImportNode(node *Node, deep bool) *Node
	AdoptNode(node *Node) *Node
	CreateAttribute(localName string) *Attr
	CreateAttributeNS(namespace string, qualifiedName string) *Attr
	CreateEvent(interfaceName string) *Event
	CreateRange() *Range
	CreateNodeIterator(root *Node, whatToShow uint64, filter *NodeFilter) *NodeIterator
	CreateTreeWalker(root *Node, whatToShow uint64, filter *NodeFilter) *TreeWalker
}
type ElementCreationOptions struct {
	is string
}

type DOMImplementation interface {
	CreateDocumentType(qualifiedName string, publicId string, systemId string) *DocumentType
	CreateDocument(namespace string, qualifiedName string, doctype *DocumentType) *XMLDocument
	CreateHTMLDocument(title string) *Document
	HasFeature() *DocumentType // Always True
}

type DocumentType struct {
	name     string
	publicId string
	systemId string
}
type DocumentTypeInterface interface {
	GetName() string
	GetPublicId() string
	GetSystemId() string
}

type DocumentFragment struct {
	Node
}
type DocumentFragmentInterface struct {
	NodeInterface
}

type HTMLCollection interface {
	GetLength() uint64
	GetTtem(index uint64) *Element
	GetNamedItem(name string) *Element
}
