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

// Text Implementation
type Text struct {
	CharacterData
	wholeText string
}

type TextInterface interface {
	CharacterDataInterface
	splitText(offset uint64) *Text
}

// CDATASection Implementation
type CDATASection struct {
	Text
}

type CDATASectionInterface interface {
	TextInterface
}

// ProcessingInstruction Implementation
type ProcessingInstruction struct {
	CharacterData
	target string
}

type ProcessingInstructionInterface interface {
	CharacterDataInterface
}

// Comment
type Comment struct {
	CharacterData
}

type CommentInterface interface {
	CharacterDataInterface
}
