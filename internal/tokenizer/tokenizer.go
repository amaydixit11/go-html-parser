package tokenizer

import (
	"fmt"
	"strings"
)

type Tokenizer struct {
	input    string
	position int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input:    input,
		position: 0,
	}
}

func (t *Tokenizer) nextChar() rune {
	if t.position >= len(t.input) {
		return 0
	}
	ch := rune(t.input[t.position])
	t.position++
	return ch
}

func (t *Tokenizer) peekChar() rune {
	if t.position >= len(t.input) {
		return 0
	}
	return rune(t.input[t.position])
}

func (t *Tokenizer) Tokenize() ([]Token, error) {
	var tokens []Token
	var textBuffer strings.Builder

	for {
		ch := t.nextChar()
		if ch == 0 {

			if textBuffer.Len() > 0 {
				tokens = append(tokens, Token{
					Type: Text,
					Data: textBuffer.String(),
				})
			}
			break
		}

		switch {
		case ch == '<':
			if textBuffer.Len() > 0 {
				tokens = append(tokens, Token{
					Type: Text,
					Data: textBuffer.String(),
				})
				textBuffer.Reset()
			}
			token, err := t.TokenizeTag()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r':
			if textBuffer.Len() > 0 {
				textBuffer.WriteRune(ch)
			}
		default:
			textBuffer.WriteRune(ch)
		}
	}
	return tokens, nil
}

func (t *Tokenizer) TokenizeTag() (Token, error) {
	var tagName TokenType
	var tagData string
	var attrs []Attribute
	var selfClosing bool

	selfClosing = false
	ch := t.nextChar()

	if ch == '/' {
		tagName = EndTag
		ch = t.nextChar()
	} else if ch == '-' {
		tagName = Comment
		return t.TokenizeComment()
	} else if strings.HasPrefix(string(t.input[t.position-1:]), "<!DOCTYPE") {
		tagName = Doctype
		return t.TokenizeDoctype()
	} else {
		tagName = StartTag
	}

	for ch != ' ' && ch != '>' && ch != '/' {
		tagData += string(ch)
		ch = t.nextChar()
	}

	for ch == ' ' {
		ch = t.nextChar()
	}

	for ch != '>' && ch != '/' {
		if ch == ' ' {
			ch = t.nextChar()
			continue
		}

		attr, err := t.parseAttribute()
		if err != nil {
			return Token{}, err
		}

		attrs = append(attrs, attr)
		for ch == ' ' {
			ch = t.nextChar()
		}
	}

	if ch == '/' {
		selfClosing = true
		ch = t.nextChar()
	}

	if ch != '>' {
		return Token{}, fmt.Errorf("unexpected character '%v' in tag at position %d", ch, t.position)

	}

	if selfClosing {
		tagName = SelfClosing
	}

	return Token{
		Type:       tagName,
		Data:       tagData,
		Attributes: attrs,
	}, nil
}

func (t *Tokenizer) TokenizeComment() (Token, error) {
	var commentText strings.Builder
	t.position += 4

	for {
		ch := t.nextChar()
		if ch == 0 || (ch == '-' && t.peekChar() == '-' && t.input[t.position] == '>') {
			break
		}
		commentText.WriteRune(ch)
	}

	return Token{
		Type: Comment,
		Data: commentText.String(),
	}, nil
}

func (t *Tokenizer) TokenizeDoctype() (Token, error) {
	var doctypeName strings.Builder
	ch := t.nextChar()

	for ch != '>' && ch != 0 {
		doctypeName.WriteRune(ch)
		ch = t.nextChar()
	}

	if ch == 0 {
		return Token{}, fmt.Errorf("unexpected EOF while parsing DOCTYPE")
	}

	return Token{
		Type: Doctype,
		Data: doctypeName.String(),
	}, nil
}

func (t *Tokenizer) parseAttribute() (Attribute, error) {
	var attrName, attrValue strings.Builder
	var quote rune

	// Parse attribute name
	for {
		ch := t.nextChar()
		if ch == '=' || ch == ' ' || ch == '>' || ch == '/' || ch == 0 {
			break
		}
		attrName.WriteRune(ch)
	}

	if t.peekChar() == '=' {
		t.nextChar() // consume '='
	}

	// Parse attribute value
	quote = t.peekChar()
	if quote == '"' || quote == '\'' {
		t.nextChar() // consume quote
	} else {
		quote = 0 // No quotes
	}

	for {
		ch := t.nextChar()
		if (quote != 0 && ch == quote) || (quote == 0 && (ch == ' ' || ch == '>' || ch == '/' || ch == 0)) {
			break
		}
		attrValue.WriteRune(ch)
	}

	if quote != 0 && t.peekChar() == quote {
		t.nextChar() // consume closing quote
	}

	return Attribute{
		Name:  attrName.String(),
		Value: attrValue.String(),
		Quote: quote,
	}, nil
}
