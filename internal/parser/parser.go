package main

type parser struct {
	// tokenizer provides the tokens for the parser.
	tokenizer *Tokenizer
	// tok is the most recently read token.
	tok Token
	// Self-closing tags like <hr/> are treated as start tags, except that
	// hasSelfClosingToken is set while they are being processed.
	hasSelfClosingToken bool
	// doc is the document root element.
	doc *Node
	// The stack of open elements (section 12.2.4.2) and active formatting
	// elements (section 12.2.4.3).
	oe, afe nodeStack
	// Element pointers (section 12.2.4.4).
	head, form *Node
	// Other parsing state flags (section 12.2.4.5).
	scripting, framesetOK bool
	// The stack of template insertion modes
	templateStack insertionModeStack
	// im is the current insertion mode.
	im insertionMode
	// originalIM is the insertion mode to go back to after completing a text
	// or inTableText insertion mode.
	originalIM insertionMode
	// fosterParenting is whether new elements should be inserted according to
	// the foster parenting rules (section 12.2.6.1).
	fosterParenting bool
	// quirks is whether the parser is operating in "quirks mode."
	quirks bool
	// fragment is whether the parser is parsing an HTML fragment.
	fragment bool
	// context is the context element when parsing an HTML fragment
	// (section 12.4).
	context *Node
}
