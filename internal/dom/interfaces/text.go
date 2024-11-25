package main

// Text Implementation
type Text struct {
	CharacterData
	wholeText string
}

type TextInterface interface {
	CharacterDataInterface
	splitText(offset uint64) *Text
}
