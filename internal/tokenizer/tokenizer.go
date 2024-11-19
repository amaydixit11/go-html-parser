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

	for {
		ch := t.nextChar()
		if ch == 0 {
			break
		}

		switch {
		case ch == '<':
			token, err := t.TokenizeTag()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r':
			continue
		default:
			tokens = append(tokens, Token{Type: Text, Data: string(ch)})

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

		attrName := ""
		for ch != '=' && ch != ' ' && ch != '>' && ch != '/' {
			attrName += string(ch)
			ch = t.nextChar()
		}

		if ch == '=' {
			ch = t.nextChar()
		}

		for ch == ' ' {
			ch = t.nextChar()
		}

		attrValue := ""
		quote := ch
		if ch == '"' || ch == '\'' {
			ch = t.nextChar()
		}

		for ch != quote && ch != '>' && ch != '/' {
			attrValue += string(ch)
			ch = t.nextChar()
		}

		if ch == quote {
			ch = t.nextChar()
		}

		attrs = append(attrs, Attribute{
			Name:  attrName,
			Value: attrValue,
			Quote: quote,
		})

		for ch == ' ' {
			ch = t.nextChar()
		}
	}

	if ch == '/' {
		selfClosing = true
		ch = t.nextChar()
	}

	if ch != '>' {
		return Token{}, fmt.Errorf("expected '>', got %v", ch)
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
