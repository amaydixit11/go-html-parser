package dom

type Attr struct {
	namespaceURI string
	prefix       string
	localName    string
	name         string
	value        string
	ownerElement *Element
	specified    bool
}

type AttrInterface interface {
	GetValue() string
	SetValue(value string) string
}
