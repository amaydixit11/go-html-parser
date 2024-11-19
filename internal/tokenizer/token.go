package tokenizer

type TokenType string

const (
	StartTag TokenType = "StartTag"
	EndTag   TokenType = "EndTag"
	Text     TokenType = "Text"
	Comment  TokenType = "Comment"
	Doctype  TokenType = "Doctype"
	// Whitespace TokenType = "Whitespace"
	// SelfClosing TokenType = "SelfClosing"
	// Error TokenType = "Error"
)

type Attribute struct {
	Name  string
	Value string
	Quote rune // `'` or `"` or 0 for unquoted
}

type Token struct {
	Type        TokenType
	Data        string
	Attributes  []Attribute
	SelfClosing bool
}
