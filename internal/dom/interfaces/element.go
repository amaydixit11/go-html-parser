package dom

type Element struct {
	Node
	namespaceURI string
	prefix       string
	localName    string
	tagName      string
	id           string
	className    string
	classList    DOMTokenList
	slot         string
	attributes   NamedNodeMap
	shadowRoot   ShadowRoot
}
type ElementInterface interface {
	NodeInterface
	GetNamespaceURI() string
	GetPrefix() string
	GetLocalName() string
	GetTagName() string
	GetId() string
	SetId(id string) string
	GetClassName() string
	SetClassName(className string) string
	GetClassList() *DOMTokenList
	GetSlot() string
	SetSlot(slot string) string
	HasAttributes() bool
	GetAttributes() *NamedNodeMap
	GetAttributeNames() []string
	GetAttribute(name string) string
	GetAttributeNS(namespace string, localName string) string
	SetAttribute(name string, value string)
	SetAttributeNS(namespace string, name string, value string)
	RemoveAttribute(name string)
	RemoveAttributeNS(namespace string, localName string)
	ToggleAttribute(name string, force bool) bool
	HasAttribute(name string) bool
	HasAttributeNS(namespace string, localName string) bool
	GetAttributeNode(name string) *Attr
	GetAttributeNodeNS(namespace string, localName string) *Attr
	SetAttributeNode(attr *Attr) *Attr
	SetAttributeNodeNS(attr *Attr) *Attr
	RemoveAttributeNode(attr *Attr) *Attr
	AttachShadow(init *ShadowRoot) *ShadowRoot
	GetShadowRoot() *ShadowRoot
	Closest(selectors string) *Element
	Matches(selectors string) bool
	WebkitMatchesSelector(selectors string) bool
	GetElementsByTagName(name string) *HTMLCollection
	GetElementsByTagNameNS(namespace string, localName string) *HTMLCollection
	GetElementsByClassName(classNames string) *HTMLCollection
	InsertAdjacentElement(where string, element *Element) *Element
	InsertAdjacentText(where string, data string)
}

type NamedNodeMap struct {
	length    uint64
	item      *Attr
	namedItem *Attr
}

type NamedNodeMapInterface interface {
	GetLength() uint64
	item(index uint64) *Attr
	getNamedItem(name string) *Attr
	getNamedItemNS(namespace string, localName string) *Attr
	setNamedItem(attr *Attr) *Attr
	setNamedItemNS(attr *Attr) *Attr
	removeNamedItem(name string) *Attr
	removeNamedItemNS(namespace string, localName string) *Attr
}

type DOMTokenList interface {
	GetLength() uint64
	GetItem(index uint64) string
	Contains(token string) bool
	Add(tokens ...string)
	Remove(tokens ...string)
	Toggle(token string, force bool) bool
	Replace(token string, newToken string) bool
	Supports(token string) bool
	GetValue() string
}
