package dom

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
