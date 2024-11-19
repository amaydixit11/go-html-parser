package tokenizer

import "strings"

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
			token, err := t.tokenizeTag()
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
	var token Token
	var tagName string
	var tagData string
	var attrs []Attribute

	ch := t.nextChar()

	if ch == '/' {
		tagName = "EndTag"
	}

	if ch == '-' {
		return t.TokenizeComment()
	}

	for {
		if ch == '>' {
			break
		}

		ch = t.nextChar()
	}
}

// <!DOCTYPE html>
// <html>
// <head>
// <title>Page Title</title>
// </head>
// <body>

// <h1>This is a Heading</h1>
// <p>This is a paragraph.</p>

// </body>
// </html>

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
