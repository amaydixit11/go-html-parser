package tokenizer

import (
	"strconv"
)

type TokenType uint32

const (
	Error TokenType = iota
	StartTag
	EndTag
	Text
	Comment
	Doctype
	SelfClosingTag
)

type Attribute struct {
	Name  string
	Value string
	Quote rune // `'` or `"` or 0 for unquoted
}

type Token struct {
	Type       TokenType
	Data       string
	Attributes []Attribute
}

func (t TokenType) String() string {
	switch t {
	case StartTag:
		return "StartTag"
	case EndTag:
		return "EndTag"
	case SelfClosingTag:
		return "SelfClosingTag"
	case Text:
		return "Text"
	case Comment:
		return "Comment"
	case Doctype:
		return "Doctype"
	case Error:
		return "Error"
	}
	return "Invalid(" + strconv.Itoa(int(t)) + ")"
}

func (t Token) tagString() string {
	if len(t.Attributes) == 0 {
		return t.Data
	}
	s := t.Data
	for _, a := range t.Attributes {
		s += (" " + a.Name + `="` + a.Value + "\"")
	}
	return s
}

func (t Token) String() string {
	switch t.Type {
	case Error:
		return ""
	case Text:
		return t.Data
	case StartTag:
		return "<" + t.tagString() + ">"
	case EndTag:
		return "</" + t.tagString() + ">"
	case SelfClosingTag:
		return "<" + t.tagString() + "/>"
	case Comment:
		return "<!--" + t.Data + "-->"
	case Doctype:
		return "<!DOCTYPE " + t.Data + ">"
	}
	return "Invalid(" + strconv.Itoa(int(t.Type)) + ")"
}
