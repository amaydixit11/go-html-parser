package dom

// Abstract CharacterData implementation
type CharacterData struct {
	Node
	data   string
	length uint64
}

type CharacterDataInterface interface {
	NodeInterface
	GetData() string
	SetData(data string) string
	GetLength() uint64
	substringData(offset uint64, count uint64) string
	appendData(data string)
	insertData(offset uint64, data string)
	deleteData(offset uint64, count uint64)
	replaceData(offset uint64, count uint64, data string)
}
