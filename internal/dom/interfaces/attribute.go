package dom

type Attr struct {
	Node
	namespaceURI string
	prefix       string
	localName    string
	name         string
	value        string
	ownerElement *Element
	specified    bool
}

type AttrInterface interface {
	NodeInterface
	GetValue() string
	SetValue(value string) string
}
